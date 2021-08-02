package aesutil

import (
	"bytes"
	"crypto/aes"
)

//一些跟aes加密相关的数据填充方法
type PaddingFunc func(origin []byte, blockSize int) []byte
type UnPaddingFunc func(cipher []byte) []byte

func PKCS5Padding(origData []byte, blockSize int) []byte {
	padLength := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padLength)}, padLength)

	return append(origData, padText...)
}

func PKCS5UnPadding(cipherData []byte) []byte {
	length := len(cipherData)
	padLength := int(cipherData[length-1])

	return cipherData[:(length - padLength)]
}

/**
zeroPadding的话，有个特殊情况的bug，不推荐使用。

假设 originData = []byte{byte(0), byte(0), byte(0)}
padding后 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
在zeroUnPadding后就会变成一个空的字节数组。而不是originData
*/
func ZeroPadding(origData []byte) []byte {
	padLength := aes.BlockSize - len(origData)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(0)}, padLength)

	return append(origData, padtext...)
}

func ZeroUnPadding(cipherData []byte) []byte {
	return bytes.TrimFunc(cipherData,
		func(r rune) bool {
			return r == rune(0)
		})
}