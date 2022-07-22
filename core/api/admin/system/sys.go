package system

import (
	"equity/pkg/app"
	"github.com/gin-gonic/gin"
)

type SysApi struct{}

// Health 系统健康检查
func (s *SysApi) Health(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ServerInfo 服务器负载信息
func (s *SysApi) ServerInfo(c *gin.Context) {
	// todo
	app.NewResponse(c).ToResponse(nil)
}
