package strutil

import (
	"strconv"
	"strings"
)

// LeftUpper 首字母转大写
func LeftUpper(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

// LeftLower 首字母转小写
func LeftLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
}

// ToUpper ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Atoi 字符串转整型
func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

// Itoa 整型转字符串
func Itoa(s int) string {
	return strconv.Itoa(s)
}

// StrToFloat64 字符串转Float64
func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
