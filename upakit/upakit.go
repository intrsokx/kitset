package upakit

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/intrsokx/kitset/upakit/internal/decrypt"
	"github.com/intrsokx/kitset/upakit/internal/encrypt"
	"github.com/intrsokx/kitset/upakit/internal/model"
	"github.com/intrsokx/kitset/upakit/pkg/httputil"
)

const (
	//https://auth.boxincredit.com/repoweb/auth
	//https://authsandbox.boxincredit.com/repoweb/auth
	UPA_AUTH_URL         = "https://auth.boxincredit.com/repoweb/auth"
	REPOWEB_URL_PATH_FMT = "%s/repoweb/api/v3/%s.do"
)

var (
	DefaultTimeOut = time.Second * 30

	//upa kit工具包自定义Msg
	MsgNet           = "[UPA KIT] network err"
	MsgEncrypt       = "[UPA KIT] encrypt err"
	MsgDecrypt       = "[UPA KIT] decrypt err"
	MsgBadDataFormat = "[UPA KIT] bad data format"
)

//TODO add log
type UPAUtil struct {
	client *httputil.HttpUtil
	lock   sync.RWMutex

	developmentId string
	authSignature string
	baseKey       string
	rsaPriKey     string
	aesKey        []byte
	Auth          *model.UpaAuthResp
}

func NewUPAUtil(developmentId, authSignature, key string) *UPAUtil {
	upa := &UPAUtil{
		developmentId: developmentId,
		authSignature: authSignature,
		baseKey:       key,
		client:        httputil.NewHttpUtil(DefaultTimeOut),
	}

	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", upa.baseKey))
	upa.rsaPriKey = buf.String()

	h := md5.New()
	h.Write([]byte(upa.baseKey))
	upa.aesKey = h.Sum(nil)

	return upa
}

func (upa *UPAUtil) SetHttpClient(client *httputil.HttpUtil) {
	upa.client = client
}

func (upa *UPAUtil) GetRepoWebHost() string {
	upa.lock.RLock()
	defer upa.lock.RUnlock()

	if len(upa.Auth.AuthInfo.ServerAddress) > 0 {
		return upa.Auth.AuthInfo.ServerAddress[0]
	}
	//default https://sandbox.geality.com
	return "https://sandbox.geality.com"
}

func (upa *UPAUtil) GetRepoToken() string {
	upa.lock.RLock()
	defer upa.lock.RUnlock()

	return upa.Auth.AuthInfo.Token
}

//如果授权码认证失败，则尝试刷新认证
func (upa *UPAUtil) RefreshAuth() error {
	upa.lock.Lock()
	defer upa.lock.Unlock()

	authResp, err := upa.getAuthResp()
	if err != nil {
		fmt.Println("upa auth fail:", err)
		return err
	}

	//更新Auth
	upa.Auth = authResp
	return nil
}

func (upa *UPAUtil) getAuthResp() (*model.UpaAuthResp, error) {
	html, err := upa.authPost()

	authResp := &model.UpaAuthResp{}
	if err := json.Unmarshal(html, authResp); err != nil {
		return authResp, err
	}

	plain, err := decrypt.RsaDecryptData(authResp.Data, upa.rsaPriKey)
	if err != nil {
		return nil, err
	}

	authInfo := &model.UpaAuthInfo{}
	if err := json.Unmarshal([]byte(plain), authInfo); err != nil {
		return nil, err
	}

	authResp.AuthInfo = authInfo
	return authResp, nil
}

func (upa *UPAUtil) authPost() ([]byte, error) {
	requestCode := strings.Replace(uuid.New().String(), "-", "", -1)
	requestTime := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	sign := fmt.Sprintf("%x", md5.Sum([]byte(upa.authSignature+requestTime)))

	param := map[string]string{}
	param["developmentId"] = upa.developmentId
	param["requestCode"] = requestCode
	param["requestTime"] = requestTime
	param["sign"] = sign
	postData, _ := json.Marshal(param)

	resp, err := upa.client.Post(UPA_AUTH_URL, postData)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (upa *UPAUtil) GetUPACommonPersonAuthServer(resourceId int, cardNo, name, idCard,
	mobile, authCode string, authFlag bool) ([]byte, error) {
	//用户调用参数预处理
	user := &model.UserInfo{
		CardNo:   cardNo,
		Name:     name,
		IdCard:   idCard,
		Mobile:   mobile,
		AuthCode: authCode,
		AuthFlag: authFlag,
	}
	//repo web url
	repoUrl := fmt.Sprintf(REPOWEB_URL_PATH_FMT, upa.GetRepoWebHost(), "UPACommonPersonAuthServer")

	var i int
Retry:
	i++
	resp, err := upa.queryUpaServer(user, resourceId, repoUrl)
	//如果有err，说明访问数据源出错，直接返回错误
	if err != nil {
		return nil, err
	}

	//授权码过期 || 授权码验证失败
	if resp.ErrorCode == "10033" || resp.ErrorCode == "10030" {
		if i > 3 {
			return resp.Bytes(), errors.New("refresh count exceed")
		}
		if err := upa.RefreshAuth(); err != nil {
			return resp.Bytes(), errors.Wrap(err, "refresh auth err")
		}
		goto Retry
	}

	return resp.Bytes(), nil
}

/**
访问银联智策服务
user 用户信息
resourceId 访问服务id
url 访问服务地址
*/
func (upa *UPAUtil) queryUpaServer(user model.UpaUser, resourceId int, url string) (*model.RepoResponse, error) {
	//user -> repoRquestQuery -> repoRequest
	repoQuery := &model.RepoRequestQuery{
		ReqParam:   user.Format(),
		OutputType: "json",
	}
	//encrypt
	queryCipher, err := encrypt.EncryptAesBase64([]byte(repoQuery.String()), upa.aesKey)
	if err != nil {
		return nil, errors.Wrap(err, MsgEncrypt)
	}

	//构建RepoRequest参数
	token := upa.GetRepoToken()
	requestCode := strings.Replace(uuid.New().String(), "-", "", -1)
	header := &model.ReqHeader{
		DevelopmentId: upa.developmentId,
		RequestCode:   requestCode,
		Token:         token,
	}
	repoRequest := &model.RepoRequest{
		Header:     header,
		ResourceId: resourceId,
		Query:      queryCipher,
	}

	resp, err := upa.doRequest(url, repoRequest)
	if err != nil {
		return nil, err
	}

	//若errCode == 0, 则对data解密
	if resp.ErrorCode == "0" {
		//decrypt data
		dataPlain, err := decrypt.DecryptAesBase64([]byte(resp.Data), upa.aesKey)
		if err != nil {
			return resp, errors.Wrap(err, MsgDecrypt)
		}
		resp.Plain = dataPlain
	}

	return resp, nil
}

//发送http请求 && 数据处理（结构化）
func (upa *UPAUtil) doRequest(url string, repoReq *model.RepoRequest) (*model.RepoResponse, error) {
	resp, err := upa.client.Post(url, []byte(repoReq.String()))
	if err != nil {
		return nil, errors.Wrap(err, MsgNet)
	}
	if resp.StatusCode != http.StatusOK {
		err := errors.New(fmt.Sprintf("http status code is %d", resp.StatusCode))
		return nil, errors.Wrap(err, MsgNet)
	}

	repoResp := &model.RepoResponse{}
	if err := json.Unmarshal(resp.Body, repoResp); err != nil {
		return nil, errors.Wrap(err, MsgBadDataFormat)
	}

	return repoResp, nil
}

func (upa *UPAUtil) GetUPAAuthRecognizeServer(cardNo, name, idCard, mobile, mode, merName, authCode string,
	authFlag bool) ([]byte, error) {
	//请求参数 -> model.UpaUser
	user := &model.RecognizeUserInfo{
		CardNo:   cardNo,
		Name:     name,
		IdCard:   idCard,
		Mobile:   mobile,
		Mode:     mode,
		MerName:  merName,
		AuthCode: authCode,
		AuthFlag: authFlag,
	}
	//repo web url
	repoUrl := fmt.Sprintf(REPOWEB_URL_PATH_FMT, upa.GetRepoWebHost(), "UPAAuthRecognizeServer")

	var i int
Retry:
	i++
	resp, err := upa.queryUpaServer(user, 0, repoUrl)
	//如果有err，说明访问数据源出错，直接返回错误
	if err != nil {
		return nil, err
	}

	//授权码过期 || 授权码验证失败
	if resp.ErrorCode == "10033" || resp.ErrorCode == "10030" {
		if i > 3 {
			return resp.Bytes(), errors.New("refresh count exceed")
		}
		if err := upa.RefreshAuth(); err != nil {
			return resp.Bytes(), errors.Wrap(err, "refresh auth err")
		}
		goto Retry
	}

	return resp.Bytes(), nil
}

//https://github.com/intrsokx/kitset
