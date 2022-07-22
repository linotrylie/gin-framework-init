package httputil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试post请求
func TestPost(t *testing.T) {
	urls := []string{
		"https://mall.xinglico.com/index.php?r=api/test/test/t",
	}
	for _, url := range urls {
		result, err := Post(url, "")
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
		fmt.Println(string(result))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
