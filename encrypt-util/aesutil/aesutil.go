package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"github.com/intrsokx/kitset/encrypt-util/aesutil/ecb"
)

//default iv
var IV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

//default encrypt is cfb encrypted
func Encrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if iv == nil {
		iv = IV
	}

	if aes.BlockSize != len(iv) {
		return nil, errors.New("IV length must equal block size")
	}

	cfbEncrypter := cipher.NewCFBEncrypter(block, iv)

	cipherText := make([]byte, len(origData))
	cfbEncrypter.XORKeyStream(cipherText, origData)

	return cipherText, nil
}

func Decrypt(cipheredData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if aes.BlockSize != len(iv) {
		return nil, errors.New("IV length must equal block size")
	}

	cfbDecrypter := cipher.NewCFBDecrypter(block, iv)

	plainText := make([]byte, len(cipheredData))
	cfbDecrypter.XORKeyStream(plainText, cipheredData)

	return plainText, nil
}


func CBCEncrypt(origData, key, iv []byte, pf PaddingFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if iv == nil {
		iv = IV
	}

	if aes.BlockSize != len(iv) {
		return nil, errors.New("IV length must equal block size")
	}

	//加密（1、先填充数据。 2、再加密）
	origData = pf(origData, block.BlockSize())
	dst := make([]byte, len(origData))

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(dst, origData)

	return dst, nil
}

func CBCDecrypt(cipherData, key, iv []byte, upf UnPaddingFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if iv == nil {
		iv = IV
	}

	if aes.BlockSize != len(iv) {
		return nil, errors.New("IV length must equal block size")
	}

	//解密（1、先解密。 2、再去除填充数据）
	dst := make([]byte, len(cipherData))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(dst, cipherData)

	return upf(dst), nil
}

func ECBEncrypt(origData, key []byte, pf PaddingFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//加密（1、先填充数据。 2、再加密）
	origData = pf(origData, block.BlockSize())
	cipherText := make([]byte, len(origData))

	ecbEncrypter := ecb.NewECBEncrypter(block)
	ecbEncrypter.CryptBlocks(cipherText, origData)

	return cipherText, nil
}

func ECBDecrypt(cipheredData, key []byte, upf UnPaddingFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//解密（1、先解密。 2、再去除填充数据）
	ecbDecrypter := ecb.NewECBDencrypter(block)

	plainText := make([]byte, len(cipheredData))
	ecbDecrypter.CryptBlocks(plainText, cipheredData)

	return upf(plainText), nil
}
