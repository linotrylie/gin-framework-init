package utils

import (
	"fmt"
	pkgErrors "github.com/pkg/errors"
)

// WapErrWithTrace 获取增强版的error信息
func WapErrWithTrace(err error, desc string) string {
	err = pkgErrors.Wrap(err, desc)
	return fmt.Sprintf("%+v", err)
}
