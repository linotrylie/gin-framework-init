package passwordutil

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// ValidatePasswd 密码比对
// pwd : 明文
// passwd: 密文
func ValidatePasswd(pwd, salt, passwd string) bool {
	return Md5Encode(pwd+salt) == passwd
}

// MakePasswd 加密
func MakePasswd(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}
