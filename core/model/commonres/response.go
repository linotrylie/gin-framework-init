package commonres

import (
	"encoding/json"
	_ "equity/docs"
	"equity/utils"
	"equity/utils/strutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 维护codes的map,避免状态码重复
var codes = map[int]string{}

const (
	ERROR   = 1
	SUCCESS = 0
)

// response 返回信息
type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Pager struct {
	Total    int64 `json:"total"`
	PageNo   int   `json:"pageNo"`
	PageSize int   `json:"pageSize"`
}

// PageResult 分页返回
type PageResult struct {
	List  interface{} `json:"list"`
	Pager Pager       `json:"pager"`
}

// ErrorResponse 返回错误信息
type ErrorResponse struct {
	Code   int         `json:"code"`
	Reason string      `json:"reason"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type ErrorResponseWithDetails struct {
	RunMode         string
	Host            string
	Path            string
	RequestDateTime string
	Method          string
	ReqBodyParam    interface{}
	ReqFormParam    interface{}
	ResData         interface{}
}

// NewErrorResp 返回错误信息
func NewErrorResp(code int, msg string) *ErrorResponse {
	if _, ok := codes[code]; ok {
		msg := fmt.Sprintf("错误码 %d 已经存在,请更换一个", code)
		panic(any(msg))
	}
	codes[code] = msg
	return &ErrorResponse{Code: code, Msg: msg}
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	dataJson, _ := json.Marshal(data)
	if len(dataJson) > 0 {
		errColumns := strutil.JsonErrorLowCase(string(dataJson))
		if len(errColumns) > 0 {
			msg := "字段不符合小驼峰命名法:"
			msg += strings.Join(errColumns, "; ")
		}
	}
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}

// errorResult 错误返回
func errorResult(code int, msg string, detail string, c *gin.Context) {
	reason := ""
	if utils.RunModeIsDebug() {
		reason = detail
	}
	c.JSON(http.StatusOK, ErrorResponse{
		code,
		reason,
		msg,
		nil,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithCodeMessage(res *ErrorResponse, detail string, c *gin.Context) {
	errorResult(res.Code, res.Msg, detail, c)
}
