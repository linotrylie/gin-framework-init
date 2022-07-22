package middleware

import (
	"bytes"
	"equity/global"
	"equity/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// AccessLog 访问日志
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		// todo 进一步完善所需的字段,尽量完善一些,日志要具有可读性,如：时间戳转换为可以阅读的时间格式
		// todo 日志信息同步到专业的日志分析系统
		// todo 对日志进行分析
		// todo 重要的错误日志要报警
		// todo 重要的日志永久保存

		fields := logger.Fields{
			"url":            c.Request.URL,
			"request_query":  c.Request.Form.Encode(),
			"request_body":   c.Request.PostForm.Encode(),
			"request_header": c.Request.Header,
			"response":       bodyWriter.body.String(),
		}
		s := "access log: method: %s,status_code:%d, " + "begin_time: %d,end_time: %d"
		// todo begin_time end_time 换成时间精确到毫秒
		global.Logger.WithFields(fields).Info(s, c.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}
