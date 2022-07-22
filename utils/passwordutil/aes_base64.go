package passwordutil

import "encoding/base64"

// AesEncodeToBase64 aes 加密数据转base64
// aesKey
// aesIv
func AesEncodeToBase64(aesKey, aesIv, data string) (string, error) {
	if data == "" {
		return "", nil
	}
	encrypt, err := AesCBCEncrypt([]byte(data), []byte(aesKey), []byte(aesIv))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypt), nil

}

// AesDecodeBase64 aes解密base64编码的数据
func AesDecodeBase64(aesKey, aesIv, data string) (string, error) {
	if data == "" {
		return "", nil
	}

	// 将base64解码 得到 aes 加密串
	base64decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	encrypt, err := AesCBCDecrypt(base64decoded, []byte(aesKey), []byte(aesIv))
	if err != nil {
		return "", err
	}
	return string(encrypt), nil
}
