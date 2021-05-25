package model

import (
	"encoding/json"
)

type UpaAuthResp struct {
	AuthInfo     *UpaAuthInfo
	Data         string `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrMessage   string `json:"errMessage"`
	RequestCode  string `json:"requestCode"`
	ResponseCode string `json:"responseCode"`
}

type UpaAuthInfo struct {
	Token         string   `json:"token"`
	Ttl           int      `json:"ttl"`
	ServerTime    int      `json:"serverTime"`
	ServerAddress []string `json:"serverAddress"`
}

type RepoRequestQuery struct {
	ReqParam   string `json:"reqParam"`
	OutputType string `json:"outputType"`
}

func (r *RepoRequestQuery) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RepoRequest struct {
	Header     *ReqHeader `json:"header"`
	ResourceId int        `json:"resourceId"`
	Query      string     `json:"query"`
}
type ReqHeader struct {
	DevelopmentId string `json:"developmentId"`
	RequestCode   string `json:"requestCode"`
	Token         string `json:"token"`
}

func (r *RepoRequest) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RepoResponse struct {
	ErrorCode    string      `json:"errorCode"`
	RequestCode  string      `json:"requestCode"`
	ResponseCode string      `json:"responseCode"`
	Data         string      `json:"data"`
	Plain        interface{} `json:"plain"`
}

func (r *RepoResponse) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RepoResponse) Bytes() []byte {
	b, _ := json.Marshal(r)
	return b
}
