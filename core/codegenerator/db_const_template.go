package codegenerator

import (
	"equity/utils/fileutil"
	"equity/utils/strutil"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// TplColumnEnums 判断数据表注释是否有需要转换为常量
func TplColumnEnums(tplColumns []*StructColumn) []*StructColumn {
	var enumInfo []*StructColumn
	for _, column := range tplColumns {
		if strings.Contains(column.Comment, "=") {
			enumInfo = append(enumInfo, column)
		}
	}
	return enumInfo
}

// GenerateTableConst 生成数据表常量文件
func GenerateTableConst(tableName string, tplColumns []*StructColumn, packageName string, fileName string) error {
	enums := TplColumnEnums(tplColumns)
	if len(enums) < 1 {
		return nil
	}

	// 文件名
	constFile := fileName
	// 如果有文件,覆盖旧数据,没有文件则新增
	writeContents := strings.Builder{}
	writeContents.Write([]byte(newLine + newLine))
	for _, enum := range enums {
		writeContents.Write([]byte("// " + enum.Comment + newLine))
		// 是否激活:1=是,2=否
		if !strings.Contains(enum.Comment, ":") {
			return errors.New("数据库注释不合法,缺少`:`符号")
		}
		items := strings.Split(enum.Comment, ":")
		if len(items) != 2 {
			msg := fmt.Sprintf(`strings.Split(enum.Comment, ":") error,expectLen=2,actLen=%d`, len(items))
			return errors.New(msg)
		}
		columnCommentsStr := items[1] // 1=是,2=否
		if !strings.Contains(columnCommentsStr, ",") {
			return errors.New("数据库注释不合法,缺少`,`符号")
		}

		columnCommentSlice := strings.Split(columnCommentsStr, ",") // [1=是,2=否]
		for _, item := range columnCommentSlice {
			if !strings.Contains(item, "=") {
				return errors.New("数据库注释不合法,缺少`=`符号")
			}
			kv := strings.Split(item, "=")
			if len(kv) != 2 {
				return errors.New(fmt.Sprintf("数据库注释不合法,err=%s", item))
			}

			enumKey, err := GetEnum(kv[1])
			if err != nil {
				return err
			}
			fullName := fmt.Sprintf("%s_%s_%s", tableName, enum.Name, enumKey.EnglishName)
			fullName = UnderscoreToLowerCamelCase(fullName)
			fullName = strutil.LeftUpper(fullName)
			constLine := fmt.Sprintf("const %s %s = %s", fullName, enum.Type, kv[0])
			writeContents.Write([]byte(constLine + newLine))
		}
		writeContents.Write([]byte(newLine))
	}

	// number of bytes written
	bytes, err := fileutil.WriteContent(constFile, os.O_RDWR|os.O_APPEND, writeContents.String())
	if err != nil {
		return err
	}
	if bytes < 1 {
		return errors.New(fmt.Sprintf("%s写入了0字节", constFile))
	}
	log.Println(fmt.Sprintf("const 文件生成成功,file=%s", constFile))
	return nil
}
