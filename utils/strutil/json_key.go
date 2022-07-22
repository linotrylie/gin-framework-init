package strutil

import (
	"equity/core/consts"
	"regexp"
	"strings"
)

// GetJsonKeys 获取json所有key
func GetJsonKeys(jsonData []byte) []string {
	var result []string
	jsonKV := make(map[string]string)
	reg := regexp.MustCompile(consts.JsonKeyRegRuler)
	matches := reg.FindAll(jsonData, -1)
	for _, v := range matches {
		s := strings.Replace(string(v), ":", "", -1)
		s = strings.Replace(s, `"`, "", -1)
		jsonKV[s] = s
	}

	for k, _ := range jsonKV {
		result = append(result, k)
	}
	return result
}

// JsonErrorLowCase 判断json key 是否小驼峰命名法
// 返回不合法的字段
func JsonErrorLowCase(jsonData string) []string {
	keys := GetJsonKeys([]byte(jsonData))
	var errColumns []string
	for _, key := range keys {
		if strings.Contains(key, "_") ||
			strings.Contains(key, "-") ||
			LeftUpper(key) == key {
			errColumns = append(errColumns, key)
		}
	}
	return errColumns
}
