package aesutil

import (
	"bytes"
	"crypto/aes"
)

/*
	一些跟aes加密相关的数据填充方法

	pkcs7填充算法: 根据blockSize计算填充所需要的长度size，然用size所表示的ASCII码(ch)对数据进行填充。
	pkcs5填充算法: 根据固定blockSize=8计算需要填充的长度size，然后使用size 所表示的ASCII码(ch)对数据进行填充。
	tips：PKCS7与PKCS5的区别在于PKCS5只填充到8字节，而PKCS7可以在1-255之间任意填充。
*/
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

func PKCS7Padding(origin []byte, blockSize int) []byte {
	padding := blockSize - len(origin)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origin, padText...)
}
func PKCS7UnPadding(cipher []byte) []byte {
	l := len(cipher)
	unpadding := int(cipher[l-1])
	return cipher[:l-unpadding]
}
