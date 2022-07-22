package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// LoadTls https 访问
func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
