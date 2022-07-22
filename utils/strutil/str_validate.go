package strutil

import (
	"equity/core/consts"
	"regexp"
)

// CheckIdCard 检验身份证
func CheckIdCard(idNumber string) bool {
	reg := regexp.MustCompile(consts.IdCardRegRuler)
	return reg.MatchString(idNumber)
}
