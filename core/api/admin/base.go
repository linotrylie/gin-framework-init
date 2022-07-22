package admin

import "github.com/gin-gonic/gin"

type BaseInter interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}
