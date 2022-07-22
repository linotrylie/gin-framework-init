package passwordutil

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"equity/utils/fileutil"
	"errors"
	"fmt"
	"os"
)

// RSAKey 秘钥对
type RSAKey struct {
	PrivateKey []byte // 私钥
	PublicKey  []byte // 公钥
}

// RsaGenKeyToBytes 生成RSA秘钥对
// 参数bits: 指定生成的秘钥的长度, 单位: bit
func RsaGenKeyToBytes(bits int) (*RSAKey, error) {
	// GenerateKey函数使用随机数据生成器 random 生成一对具有指定字位数的RSA密钥
	// 参数1: Reader是一个全局、共享的密码用强随机数生成器
	// 参数2: 秘钥的位数 - bit
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	// MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	// Block代表PEM编码的结构, 对其进行设置
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// 创建buffer
	privateBuffer := new(bytes.Buffer)
	publicBuffer := new(bytes.Buffer)
	err = pem.Encode(privateBuffer, &block)
	if err != nil {
		return nil, err
	}

	// 生成公钥
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, err
	}
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	// 编码公钥
	err = pem.Encode(publicBuffer, &block)
	if err != nil {
		return nil, err
	}
	rsaKey := &RSAKey{
		PrivateKey: privateBuffer.Bytes(),
		PublicKey:  publicBuffer.Bytes(),
	}
	return rsaKey, nil
}

// RsaGenKey 参数bits: 指定生成的秘钥的长度, 单位: bit
// 保存到文件
func RsaGenKey(bits int) error {
	// 文件分隔符
	fileSeparator := fileutil.FileSeparator()
	// 1. 生成私钥文件
	// GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	// 参数1: Reader是一个全局、共享的密码用强随机数生成器
	// 参数2: 秘钥的位数 - bit
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 2. MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	// 3. Block代表PEM编码的结构, 对其进行设置
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// 4. 创建文件
	privateFileName := fmt.Sprintf(".%sprivate.pem", fileSeparator)
	privateFile, err := os.Create(privateFileName)
	if err != nil {
		return err
	}

	// 5. 使用pem编码, 并将数据写入文件中
	err = pem.Encode(privateFile, &block)
	if err != nil {
		return err
	}

	// 6. 最后的时候关闭文件
	defer func() {
		_ = privateFile.Close()
	}()

	// 7. 生成公钥文件
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	publicFileName := fmt.Sprintf(".%spublic.pem", fileSeparator)
	pubFile, err := os.Create(publicFileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = pubFile.Close()
	}()

	// 8. 编码公钥, 写入文件
	return pem.Encode(pubFile, &block)
}

// RSAEncryptByFile rsa 通过文件加密
// wantEncryptData 要加密的数据
// 公钥文件的路径
func RSAEncryptByFile(wantEncryptData []byte, filename string) ([]byte, error) {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// 2. 读文件
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	allText := make([]byte, info.Size())

	_, err = file.Read(allText)
	if err != nil {
		return nil, err
	}

	// 3. 关闭文件
	defer func() {
		_ = file.Close()
	}()

	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	if block == nil {
		return nil, err
	}

	// 5. 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 6. 公钥加密
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, wantEncryptData)
}

// RSADecryptByFile rsa 解密
// wantDecryptData 要解密的数据
// 私钥文件的路径
func RSADecryptByFile(wantDecryptData []byte, filename string) ([]byte, error) {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// 2. 读文件
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	allText := make([]byte, info.Size())
	_, err = file.Read(allText)
	if err != nil {
		return nil, err
	}

	// 3. 关闭文件
	defer func() {
		_ = file.Close()
	}()

	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)

	// 5. 解析一个pem格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	// 6. 私钥解密
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, wantDecryptData)
}

// RSAEncrypt rsa 通过私钥加密
// wantEncryptData 要加密的数据
// privateKey 私钥
func RSAEncrypt(wantEncryptData, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("读取数据失败")
	}

	// 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 公钥加密
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, wantEncryptData)
}

// RSADecrypt 通过私钥解密
// publicKey 私钥
func RSADecrypt(wantDecryptData []byte, publicKey []byte) ([]byte, error) {
	// 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("pem.Decode err")
	}

	// 解析一个pem格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 私钥解密
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, wantDecryptData)
}
