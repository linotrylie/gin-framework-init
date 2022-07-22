package codegenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnum(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"是", "true"},
		{"否", "false"},
		{"按钮", "button"},
		{"文件", "dir"},
	}

	for _, test := range tests {
		enum, err := GetEnum(test.input)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, enum.EnglishName, test.want)
		}
	}
}
