package codegenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetColumnName(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{`json:"modify_time"`, `json:"modifyTime"`},
		{`json:"modify_time_1"`, `json:"modifyTime1"`},
	}
	for _, test := range tests {
		res := GetColumnName(test.input)
		assert.Equal(t, res, test.want)
	}
}
