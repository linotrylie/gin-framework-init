package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRunMode(t *testing.T) {
	fmt.Println(gin.Mode())
	fmt.Println(RunModeIsDebug())
	fmt.Println(RunModeIsRelease())
}
