package codegenerator

import (
	"equity/core/consts"
	"equity/utils"
	"equity/utils/fileutil"
	"equity/utils/strutil"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// GenerateModelResponse 返参
// 如果有文件,不能覆盖
func GenerateModelResponse(tableName string, tplColumns []*StructColumn, packageName string) error {
	if len(tplColumns) < 1 {
		return errors.New("没有数据表字段,请先完善信息")
	}

	requestName := UnderscoreToLowerCamelCase(strings.Replace(tableName, consts.TablePrefix, "", 1))
	requestName = strutil.LeftUpper(requestName)
	baseModelStructName := getBaseModelStructName(requestName)

	// 当前目录
	getwd, _ := os.Getwd()
	getwd = strings.Replace(getwd, "codegenerator", "", -1)
	filePath := getwd + "model" + fileutil.FileSeparator() + packageName + "model" + fileutil.FileSeparator() + "response"
	fileName := strings.Replace(tableName, consts.TablePrefix, "", 1)
	if !fileutil.DirIsExist(filePath) {
		if err := os.MkdirAll(filePath, 0777); err != nil {
			return err
		}
	}

	// 完整路径
	fullPath := filePath + fileutil.FileSeparator() + fileName + "_res" + goFileExtension
	if fileutil.FileIsExist(fullPath) {
		return errors.New(fmt.Sprintf("response file exist,file=%s", fullPath))
	}

	writeContents := strings.Builder{}
	writeContents.Write([]byte(responsePackage))
	writeContents.Write([]byte(newLine + newLine))

	// 是否需要导入依赖
	imports := []string{}
	for _, column := range tplColumns {
		if column.Type == "time.Time" {
			if val, ok := importMap[column.Type]; ok {
				isExist := utils.InArray(val, imports)
				if isExist == false {
					imports = append(imports, importMap[column.Type])
				}
			}
		}
	}

	if len(imports) > 0 {
		writeContents.Write([]byte("import (" + newLine))
		for _, importName := range imports {
			writeContents.Write([]byte(fmt.Sprintf(`   "%s"`, importName) + newLine))
		}
		writeContents.Write([]byte(")" + newLine + newLine))
	}

	writeContents.Write([]byte(fmt.Sprintf("// %s", baseModelStructName+" 基本信息 ") + newLine))
	writeContents.Write([]byte(fmt.Sprintf("type %s struct {", baseModelStructName) + newLine))
	for _, column := range tplColumns {
		columnNameFormat := UnderscoreToUpperCamelCase(column.Name)
		isIgnore := utils.InArray(columnNameFormat, getDataIgnoreColumns)
		if isIgnore == true {
			continue
		}
		writeContents.Write([]byte(fmt.Sprintf("   // %s", column.Comment+newLine)))
		info := fmt.Sprintf("  %s %s `%s`", columnNameFormat, column.Type, GetColumnName(column.Tag))
		writeContents.Write([]byte(info + newLine))
	}
	writeContents.Write([]byte(fmt.Sprintf(newLine + "}" + newLine)))
	bytes, err := fileutil.WriteContent(fullPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, writeContents.String())
	if err != nil {
		return err
	}
	if bytes < 1 {
		return errors.New(fmt.Sprintf("%s写入了0字节", fullPath))
	}
	log.Println(fmt.Sprintf("response 文件生成成功,file=%s", fullPath))
	return nil
}
