package app

import (
	"equity/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type ListData struct {
	List  interface{} `json:"list"`
	Pager Pager       `json:"pager"`
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalRows int `json:"totalRows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponseWithData 操作成功且返回数据
func (r *Response) ToResponseWithData(data interface{}) {
	res := gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	}
	r.Ctx.JSON(http.StatusOK, res)
}

// ToResponse 返回普通数据
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{"code": 0, "msg": "ok!"}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToResponseList 返回列表数据
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	data := gin.H{
		"code": 0,
		"msg":  "ok",
		"data": ListData{
			List: list,
			Pager: Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
