package strutil

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	// UpperCaseLetters 大写字母
	UpperCaseLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// LowerCaseLetters 小写字母
	LowerCaseLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// Numbers 数字
	Numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// LetterNumbers 大小写字母混合数字
	LetterNumbers = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3",
		"4", "5", "6", "7", "8", "9"}

	// LettersNumbersCaptcha  大小写字母、数字混合(去除易混淆)
	LettersNumbersCaptcha = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "2", "3",
		"4", "5", "6", "7", "8", "9"}
)

// GetRandom 生成随机数
func GetRandom(resLen int, seeds []string) string {
	// 将str拼接n次
	var builder strings.Builder
	for i := 0; i < resLen; i++ {
		// 生成随机数
		str := seeds[rand.Intn(len(seeds))]
		builder.WriteString(str)
	}
	// 返回拼接后的字符串
	return builder.String()
}
