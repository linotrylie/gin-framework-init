package fileutil

import (
	"equity/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFileSeparator 测试获取文分隔符
func TestFileSeparator(t *testing.T) {
	sep := FileSeparator()
	if utils.OsIsWindows() {
		assert.Equal(t, "\\", sep)
	} else {
		assert.Equal(t, "/", sep)
	}
}
