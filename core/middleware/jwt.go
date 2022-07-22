package middleware

import (
	"equity/core/consts"
	"equity/core/model/commonres"
	"equity/pkg/app"
	"equity/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT jwt
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			errCode = errcode.Success
		)
		if s, exist := c.GetQuery(consts.Authorization); exist {
			token = s
		} else {
			token = c.GetHeader(consts.Authorization)
		}
		if token == "" {
			errCode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					errCode = errcode.UnauthorizedAuthTokenTimeout
				default:
					errCode = errcode.UnauthorizedAuthTokenError
				}
			}
		}
		if errCode != errcode.Success {
			commonres.FailWithMessage(errCode.Msg(), c)
			c.Abort()
			return
		}

		c.Set(consts.Authorization, token)
		payload, err := app.GetJwtPayload(token)
		if err != nil {
			commonres.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		if payload.Uid < 1 {
			commonres.FailWithMessage("获取用户id失败", c)
			c.Abort()
			return
		}
		c.Set(consts.AdminUid, payload.Uid)
		c.Next()
	}
}
