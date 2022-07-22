package passwordutil

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

//  AES: Advanced Encryption Standard
//  优点: 算法公开、计算量小、加密速度快、加密效率高
//  缺点: 发送方和接收方必须商定好密钥,然后使双方都能保存好密钥,密钥管理成为双方的负担。
//  使用场景: 相对大一点的数据量或关键数据的加密
//  秘钥长度: 16, 24, 32,分别对应位数长度为:128,192,256
//  AES数据块长度为128位，所以IV长度需要为16个字符（ECB模式不用IV)

// AesCBCEncrypt 加密
func AesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	err := validateKeyLen(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != 16 {
		return nil, errors.New("iv长度必须为16位")
	}

	// 创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 判断加密快的大小
	blockSize := block.BlockSize()

	// 填充
	encryptBytes := pKCS5Padding(data, blockSize)

	// 初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))

	// 使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)

	return crypted, nil
}

// AesCBCDecrypt 解密
func AesCBCDecrypt(data, key, iv []byte) ([]byte, error) {
	err := validateKeyLen(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != 16 {
		return nil, errors.New("iv长度必须为16位")
	}

	// 创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用cbc
	blockMode := cipher.NewCBCDecrypter(block, iv)

	// 初始化解密数据接收切片
	crypted := make([]byte, len(data))

	// 执行解密
	blockMode.CryptBlocks(crypted, data)
	// 去除填充
	return pKCS5UnPadding(crypted)
}

// validateKeyLen
func validateKeyLen(key []byte) error {
	switch len(key) {
	default:
		return errors.New(fmt.Sprintf("秘钥长度只允许16,24,32,当前秘钥 = %s", key))
	case 16, 24, 32:
		break
	}
	return nil
}
