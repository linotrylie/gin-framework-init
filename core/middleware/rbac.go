package middleware

import (
	"equity/core/model/authmodel"
	"equity/core/model/commonreq"
	"equity/core/model/commonres"
	"equity/core/service"
	"equity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

var authService service.ServicesGroup

// RBAC 管理员权限
func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户id
		uid, err := utils.GetUidByToken(c)
		if err != nil {
			commonres.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		if uid < 1 {
			commonres.FailWithMessage(fmt.Sprintf("用户id获取失败,uid=%d", uid), c)
			c.Abort()
			return
		}

		// 超级管理员不需要验证
		getUserParam := commonreq.PrimaryIdParam{Id: uid}
		user, err := authService.RBACServiceGroup.UserService.UserDetail(c, &getUserParam)
		if err != nil {
			commonres.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}

		if user == nil || user.Info == nil {
			commonres.FailWithMessage(fmt.Sprintf("用户信息获取失败,id=%d", uid), c)
			c.Abort()
			return
		}

		if user.Info.IsRoot == authmodel.AuthUserIsRootTrue {
			c.Next()
			return
		}

		// 获取用户拥有的节点url
		urls, err := authService.RBACServiceGroup.UserService.GetUserNodeUrls(c, uid)
		if err != nil {
			commonres.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}

		if len(urls) < 1 {
			commonres.FailWithCodeMessage(commonres.ErrorNotAuth, "无权限访问", c)
			c.Abort()
			return
		}

		// 当前请求的路由
		exists := utils.InArray(c.FullPath(), urls)
		if exists == false {
			commonres.FailWithCodeMessage(commonres.ErrorNotAuth, "未授权", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
