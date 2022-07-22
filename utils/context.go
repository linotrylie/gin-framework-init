package utils

import (
	"equity/core/consts"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUidByToken 获取用户id
func GetUidByToken(c *gin.Context) (int32, error) {
	uid, exists := c.Get(consts.AdminUid)
	if !exists {
		return 0, errors.New("获取uid失败")
	}
	uidString := fmt.Sprintf("%v", uid)
	uidInt, err := strconv.Atoi(uidString)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("uid转整型失败,原始数据=%s", uid))
	}
	return int32(uidInt), nil
}
