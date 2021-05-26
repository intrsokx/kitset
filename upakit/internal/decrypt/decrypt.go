package decrypt

import (
	"encoding/base64"
	"fmt"
	"github.com/intrsokx/encrypt-util/aesutil"
	"github.com/intrsokx/encrypt-util/rsautil"
	"net/url"
	"strconv"
)

func RsaDecryptData(cipher, priKey string) (string, error) {
	cfg := &rsautil.Config{
		ParsePrivateKey:    rsautil.ParsePKCS8PrivateKey,
		DecryptWithPrivate: rsautil.DecryptPKCS1v15,
		CryptSub:           true,
		SubStep:            128,
	}

	data := hexStrToByte(cipher)

	plain, err := rsautil.Decrypt(data, []byte(priKey), cfg)
	if err != nil {
		return "", err
	}

	return url.QueryUnescape(string(plain))
}

func hexStrToByte(str string) []byte {
	b := []byte{}
	for i := 0; i < len(str)-1; i = i + 2 {
		base, _ := strconv.ParseInt(str[i:i+2], 16, 10)
		b = append(b, byte(base))
	}

	return b
}

func ByteToHexStr(b []byte) string {
	var s string
	for i := 0; i < len(b); i++ {
		s = s + fmt.Sprintf("%02X", b[i])
	}
	return s
}

func DecryptAesBase64(cipher []byte, key []byte) (string, error) {
	txt, err := base64.StdEncoding.DecodeString(string(cipher))
	if err != nil {
		return "", err
	}

	plain, err := aesutil.ECBDecrypt(txt, key, aesutil.PKCS5UnPadding)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}
