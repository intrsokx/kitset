package aesutil

import (
	"crypto/md5"
	"fmt"
	"testing"
)

var key = fmt.Sprintf("%x", md5.Sum([]byte("aes key")))
var origin = "origin data"

func TestEncrypt(t *testing.T) {
	cipher, err := Encrypt([]byte(origin), []byte(key), IV)
	if err != nil {
		t.Error("aes encrypt fail:", err)
	}

	plain, err := Decrypt(cipher, []byte(key), IV)
	if err != nil {
		t.Error("aes decrypt fail:", err)
	}

	if string(plain) != origin {
		t.Errorf("plain must be origin; plain is %s", plain)
	}
}

func TestCBCEncrypt(t *testing.T) {
	cipher, err := CBCEncrypt([]byte(origin), []byte(key), IV, PKCS5Padding)
	if err != nil {
		t.Error("aes cbc encrypt fail:", err)
	}

	plain, err := CBCDecrypt(cipher, []byte(key), IV, PKCS5UnPadding)
	if err != nil {
		t.Error("aes cbc decrypt fail:", err)
	}

	if string(plain) != origin {
		t.Errorf("plain must be origin; plain is %s", plain)
	}
}

func TestECBEncrypt(t *testing.T) {
	cipher, err := ECBEncrypt([]byte(origin), []byte(key), PKCS5Padding)
	if err != nil {
		t.Error("aes cbc encrypt fail:", err)
	}

	plain, err := ECBDecrypt(cipher, []byte(key), PKCS5UnPadding)
	if err != nil {
		t.Error("aes cbc decrypt fail:", err)
	}

	if string(plain) != origin {
		t.Errorf("plain must be origin; plain is %s", plain)
	}
}