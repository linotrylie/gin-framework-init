package passwordutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试 des ECB 加密解密
func TestDesECBEncrypt(t *testing.T) {
	var tests = []struct {
		waitEncryptData string // 待加密字符串
		key             []byte // 秘钥
		res             string // 加密结果
	}{
		{"ihaveadream", []byte("12345678"), "/+SnK1JDruvZjR6vZV7Ozw=="},
		{"测试中文", []byte("12345678"), "0OlHlpwkA874M0Big4MnKw=="},
		{"测试中文ihaveadream", []byte("12345678"), "0OlHlpwkA848SSW00BZrBFUtJsgc4yNE"},
		{"测试日文:アグア氏の度に叫ぶのはいいことだと思います", []byte("12345678"), "HljgeZ/e2aKcOsu9Vqf6Uc7jbwLTX58d1x5cQ9pAlW1ARdSNcb8JecOinI0sRr3ZhsSVikJjKnJa+1P7RkZ4DNoltdUlEY+7Gc9SzRJE6oM="},
	}
	for _, test := range tests {
		// 加密
		encrypt, err := DesECBEncrypt([]byte(test.waitEncryptData), test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.res, encrypt)

		// 解密
		decrypt, err := DesECBDecrypt(test.res, test.key)
		assert.Nil(t, err)
		assert.Equal(t, []byte(test.waitEncryptData), decrypt)
	}
}
