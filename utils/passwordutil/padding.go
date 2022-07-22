package passwordutil

import (
	"bytes"
	"github.com/pkg/errors"
)

// pKCS5Padding 明文补码算法
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// pKCS5UnPadding 明文减码算法
func pKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unPadding := int(origData[length-1])
	if length < unPadding {
		return nil, errors.New("解密去填充错误,请核对加密模式,填充,秘钥长度是否正确")
	}
	return origData[:(length - unPadding)], nil
}
