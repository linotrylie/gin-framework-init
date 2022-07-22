package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArrayInt(t *testing.T) {
	var testInt = []struct {
		A    int
		B    []int
		Want bool
	}{
		{1, []int{11, 2, 3}, false},
		{-1, []int{1, 2, 3, 0}, false},
		{32, []int{-2147483648, 2147483647}, false},
		{-2147483648, []int{-2147483648, 2147483647}, true},
	}
	for _, test := range testInt {
		actual := InArray(test.A, test.B)
		assert.Equal(t, test.Want, actual)
	}
}

func TestInArrayStr(t *testing.T) {
	var testStr = []struct {
		A    string
		B    []string
		Want bool
	}{
		{"a", []string{"aaa", "b", "c"}, false},
		{"中文", []string{"是对的", "jsbhbhb", "dddd", "中文", "问"}, true},
		{"にほんご", []string{"にほんごss", "ddd"}, false},
		{"hello", []string{"hello", "ddd"}, true},
	}
	for _, test := range testStr {
		actual := InArray(test.A, test.B)
		assert.Equal(t, test.Want, actual)
	}
}
