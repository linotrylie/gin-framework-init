package middleware

import (
	"equity/core/consts"
	"equity/core/model/commonreq"
	"equity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

// Recovery 错误恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "\n%v :\n"
				baseInfo := fmt.Sprintf(s, err)
				panicInfo := baseInfo + utils.PanicInfo()
				if strings.HasPrefix(c.FullPath(), "/"+consts.AdminApiPrefix) {
					var req commonreq.FrontRequest
					err := c.ShouldBindBodyWith(&req, binding.JSON)
					if err != nil {
						//commonres.Response(c, consts.OpenApiErrorDataParseFail, err)
						c.Abort()
						return
					}
					//commonres.Response(c, consts.OpenApiSystemErrorStatus, errors.New(panicInfo))
					c.Abort()
					return
				} else {
					fmt.Println(panicInfo)
					// 管理后台或
					//go message.SendMsgToFeiShu(panicInfo)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
