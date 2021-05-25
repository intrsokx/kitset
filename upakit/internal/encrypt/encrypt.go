package encrypt

import (
	"encoding/base64"
	"github.com/intrsokx/encrypt-util/aesutil"
)

func EncryptAesBase64(data []byte, key []byte) (string, error) {
	cipher, err := aesutil.ECBEncrypt(data, key, aesutil.PKCS5Padding)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}
