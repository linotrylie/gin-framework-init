package passwordutil

import (
	"crypto/des"
	"encoding/base64"
	"errors"
)

// DesECBEncrypt ECB 加密
func DesECBEncrypt(data, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	data = pKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("PKCS5Padding error")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

// DesECBDecrypt ECB 解密
func DesECBDecrypt(d string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return []byte(""), err
	}
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return []byte(""), errors.New("len(data)%bs != 0")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return pKCS5UnPadding(out)
}
