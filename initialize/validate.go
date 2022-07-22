package initialize

import (
	"bufio"
	"equity/utils/fileutil"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ValidateCode() error {
	err := codeValidate()
	if err != nil {
		return err
	}
	return nil
}

// 代码检查
// 验证 gorm 实参形参个数是否一致
// 是否含有中文问号和中文逗号
// 验证 gorm 形参和实参是否一一对应
func codeValidate() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var filePaths []string

	// 获取当前目录下的所有文件或目录信息
	err = filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, path := range filePaths {
		err := GormCodeValidate(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// getCurrentFilePath 获取当前文件路径
func getCurrentFilePath() string {
	workDir, _ := os.Getwd()
	currentFile := http.Dir(workDir + fileutil.FileSeparator() + "initialize" + fileutil.FileSeparator() + "validate.go")
	return string(currentFile)
}

// GormCodeValidate 验证gorm形参实参是否匹配
func GormCodeValidate(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	lineNum := 0
	buf := bufio.NewReader(file)
	for {
		codes, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		lineNum++
		codes = strings.TrimSpace(codes)
		currentFile := getCurrentFilePath()
		// 不能含有中文问号
		if strings.Contains(codes, "？ ") {
			if string(currentFile) != path && strings.Contains(path, "_dao.go") {
				msg := fmt.Sprintf("[GORM]不能含有中文问号`？ `:\nfile = %s,\nline = %d,\ncode = %s", path, lineNum, codes)
				return errors.New(msg)
			}
		}

		// 不能含有中文逗号
		if strings.Contains(codes, "，") && strings.Contains(path, "_dao.go") {
			if string(currentFile) != path {
				msg := fmt.Sprintf("[GORM]不能含有中文逗号`，`:\nfile = %s,\nline = %d,\ncode = %s", path, lineNum, codes)
				return errors.New(msg)
			}
		}

		matchString, err := regexp.Match(`\("*\??"`, []byte(codes)) // ("id = ? ", pid)
		if err != nil {
			return err
		}
		if matchString == true {
			// 形参个数
			formalParameterCount := strings.Count(codes, "?")
			// 实参个数
			actualParameterCount := strings.Count(codes, ",")
			if formalParameterCount > 0 && formalParameterCount != actualParameterCount {
				if strings.Contains(path, "_dao.go") {
					msg := fmt.Sprintf("[GORM]形参实参对应错误:\n\rfile = %s,\n\rline = %d,\n\rcode = %s,\r\n形参个数 = %d,\r\n实参个数 = %d",
						path, lineNum, codes, formalParameterCount, actualParameterCount)
					return errors.New(msg)
				}
			}
		}
	}
	return nil
}
