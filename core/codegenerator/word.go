package codegenerator

import (
	"fmt"
	"strings"
	"unicode"
)

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 转小驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// GetColumnName 数据库字段转小驼峰
// `json:"modify_time"` ===> `json:"modifyTime"`
func GetColumnName(colomnName string) string {
	colomnName = strings.Replace(colomnName, "`", "", -1)
	colomnName = strings.Replace(colomnName, "json:\"", "", -1)
	colomnName = strings.Replace(colomnName, "\"", "", -1)
	return fmt.Sprintf(`json:"%s"`, UnderscoreToLowerCamelCase(colomnName))
}
