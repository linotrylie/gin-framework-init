package fileutil

import (
	"equity/utils"
	"os"
)

// DirIsExist 目录是否存在
// path 目录路径
func DirIsExist(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// FileIsExist 判断文件是否存在
// filePath 文件路径
func FileIsExist(filePath string) bool {
	if !IsFile(filePath) {
		return false
	}
	fileInfo, err := os.Stat(filePath)
	if fileInfo != nil && err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
// path 文件夹路径
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
// path 文件路径
func IsFile(path string) bool {
	return !IsDir(path)
}

// WriteContent 写入文件
// filePath 文件路径
// contents 要写入的内容
func WriteContent(filePath string, flag int, contents string) (int, error) {
	f, err := os.OpenFile(filePath, flag, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = f.Close()
	}()
	return f.Write([]byte(contents))
}

// FileSeparator 获取文件分隔符
func FileSeparator() string {
	if utils.OsIsWindows() {
		return "\\"
	}
	return "/"
}
