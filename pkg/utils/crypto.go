/*
@Time : 2020-02-01 11:33 
@Author : peanut
@File : crypto
@Software: GoLand
*/

package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(orig string, key string) (string, error) {
	origData := []byte(orig)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	encryptedByte := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptedByte, origData)
	return base64.StdEncoding.EncodeToString(encryptedByte), nil
}
func AesDecrypt(encrypted string, key string) (string, error) {
	encryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", nil
	}
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	orig := make([]byte, len(encryptedByte))
	blockMode.CryptBlocks(orig, encryptedByte)
	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

func PKCS7Padding(cipher []byte, size int) []byte {
	padding := size - len(cipher)%size
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipher, text...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	padding := int(origData[length-1])
	return origData[:(length - padding)]
}
