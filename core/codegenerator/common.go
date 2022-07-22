package codegenerator

import "strings"

// getModelPackageName 获取 model 包名
func getModelPackageName(packageName string) string {
	return packageName + "model"
}

// getPageSearchModelStructName 获取搜索结构体名称
func getPageSearchModelStructName(requestName string) string {
	return requestName + opPageSearch
}

// getBaseModelStructName 获取数据模型基本信息
func getBaseModelStructName(requestName string) string {
	return requestName + opInfo
}

// getAbbreviationLetter 获取接收者简写,
// 不能等于c或C
func getAbbreviationLetter(requestName string) string {
	if requestName == "" {
		return "cc"
	}
	if strings.HasPrefix(requestName, "c") || strings.HasPrefix(requestName, "C") {
		if len(requestName) < 2 {
			return requestName[0:1] + requestName[0:1]
		}
		return requestName[0:2]
	}
	return requestName[0:1]
}
