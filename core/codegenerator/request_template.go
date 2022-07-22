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
	"strconv"
	"strings"
)

// GenerateModelRequest 生成入参及其校验信息
// 如果有文件,不能覆盖
func GenerateModelRequest(tableName string, tplColumns []*StructColumn, packageName string) error {
	if len(tplColumns) < 1 {
		return errors.New("没有数据表字段,请先完善信息")
	}

	requestName := UnderscoreToLowerCamelCase(strings.Replace(tableName, consts.TablePrefix, "", 1))
	requestName = strutil.LeftUpper(requestName)

	// 当前目录
	getwd, _ := os.Getwd()
	getwd = strings.Replace(getwd, "codegenerator", "", -1)

	// 目录
	filePath := getwd + "model" + fileutil.FileSeparator() + packageName + "model" + fileutil.FileSeparator() + "request"

	// 文件名
	fileName := strings.Replace(tableName, consts.TablePrefix, "", 1)
	if !fileutil.DirIsExist(filePath) {
		if err := os.MkdirAll(filePath, 0777); err != nil {
			return err
		}
	}

	// 完整的文件路径
	fullPath := filePath + fileutil.FileSeparator() + fileName + "_req" + goFileExtension
	if fileutil.FileIsExist(fullPath) {
		return errors.New(fmt.Sprintf("request file exist,file=%s", fullPath))
	}
	allContents := strings.Builder{}

	// 新增数据的参数信息
	allContents.Write([]byte(getCreateParamInfo(requestName, tplColumns)))

	// 编辑数据的参数信息
	allContents.Write([]byte(getUpdateParamInfo(requestName, tplColumns)))

	// 分页搜索参数信息
	allContents.Write([]byte(getSearchPaginationParamInfo(requestName)))

	bytes, err := fileutil.WriteContent(fullPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, allContents.String())
	if err != nil {
		return err
	}
	if bytes < 1 {
		return errors.New(fmt.Sprintf("%s写入了0字节", fullPath))
	}
	log.Println(fmt.Sprintf("request 文件生成成功,file=%s", fullPath))
	return nil
}

// getSearchPaginationParamInfo 获取分页搜索参数
func getSearchPaginationParamInfo(requestName string) string {
	pageSearchModelStructName := getPageSearchModelStructName(requestName)
	writeContents := strings.Builder{}
	writeContents.Write([]byte(newLine + newLine))

	writeContents.Write([]byte(fmt.Sprintf("// %s", pageSearchModelStructName+" 分页获取 ") + newLine))
	writeContents.Write([]byte(fmt.Sprintf("type %s struct {", pageSearchModelStructName) + newLine))
	writeContents.Write([]byte("Name string" + newLine))
	writeContents.Write([]byte("PageNo int" + newLine))
	writeContents.Write([]byte("PageSize int" + newLine))
	writeContents.Write([]byte(fmt.Sprintf(newLine + "}" + newLine)))

	// 生成数据验证
	writeContents.Write([]byte(fmt.Sprintf("func (param *%s) Validate() error {", pageSearchModelStructName) + newLine))
	writeContents.Write([]byte(`   if param == nil {
        return errors.New(consts.ParamRequired)
    }` + newLine))

	writeContents.Write([]byte(`return validation.ValidateStruct(param,` + newLine))
	pageNoFormat := fmt.Sprintf(`validation.Field(&param.PageNo,
		validation.Min(1).Error(consts.MinIsOne),
	),`)
	pageSizeFormat := fmt.Sprintf(`validation.Field(&param.PageSize,
		validation.Min(1).Error(consts.MinIsOne),
		validation.Max(20).Error(consts.MaxPageSizeIsTwenty),
	),`)
	writeContents.Write([]byte(pageNoFormat + newLine))
	writeContents.Write([]byte(pageSizeFormat + newLine))
	writeContents.Write([]byte(fmt.Sprintf(")" + newLine)))
	writeContents.Write([]byte(fmt.Sprintf("}" + newLine)))
	return writeContents.String()
}

// getCreateParamInfo 获取新增数据 Request 数据
func getCreateParamInfo(requestName string, tplColumns []*StructColumn) string {
	createModelStructName := requestName + opCreate
	writeContents := strings.Builder{}
	writeContents.Write([]byte(requestPackage))
	writeContents.Write([]byte(newLine + newLine))
	writeContents.Write([]byte(fmt.Sprintf("// %s", createModelStructName+" 创建 ") + newLine))
	writeContents.Write([]byte(fmt.Sprintf("type %s struct {", createModelStructName) + newLine))
	// 新增数据选中的字段
	checkedColumns := []*StructColumn{}
	for _, column := range tplColumns {
		// 忽略自增字段
		if column.Extra == autoIncrement {
			continue
		}

		// 忽略新增数据指定的字段
		columnNameFormat := UnderscoreToUpperCamelCase(column.Name)
		isIgnore := utils.InArray(columnNameFormat, createIgnoreColumns)
		if isIgnore == true {
			continue
		}

		writeContents.Write([]byte(fmt.Sprintf("   // %s", column.Comment+newLine)))
		info := fmt.Sprintf("  %s %s `%s`", columnNameFormat, column.Type, GetColumnName(column.Tag))
		writeContents.Write([]byte(info + newLine))
		checkedColumns = append(checkedColumns, column)
	}
	writeContents.Write([]byte(fmt.Sprintf(newLine + "}" + newLine)))

	// 生成数据验证
	writeContents.Write([]byte(fmt.Sprintf("func (param *%s) Validate() error {", createModelStructName) + newLine))
	writeContents.Write([]byte(`   if param == nil {
        return errors.New(consts.ParamRequired)
    }` + newLine))

	rulesContent := getValidateRules(checkedColumns)
	writeContents.Write([]byte(rulesContent + newLine))
	writeContents.Write([]byte(fmt.Sprintf("}" + newLine)))
	return writeContents.String()
}

// getUpdateParamInfo 获取编辑数据 Request 数据
func getUpdateParamInfo(requestName string, tplColumns []*StructColumn) string {
	updateModelStructName := requestName + opUpdate
	writeContents := strings.Builder{}
	writeContents.Write([]byte(newLine + newLine))
	writeContents.Write([]byte(fmt.Sprintf("// %s", updateModelStructName+" 编辑 ") + newLine))
	writeContents.Write([]byte(fmt.Sprintf("type %s struct {", updateModelStructName) + newLine))
	// 编辑数据选中的字段
	checkedColumns := []*StructColumn{}
	for _, column := range tplColumns {
		// 忽略编辑数据指定的字段
		columnNameFormat := UnderscoreToUpperCamelCase(column.Name)
		isIgnore := utils.InArray(columnNameFormat, updateIgnoreColumns)
		if isIgnore == true {
			continue
		}
		writeContents.Write([]byte(fmt.Sprintf("   // %s", column.Comment+newLine)))
		info := fmt.Sprintf("  %s %s `%s`", columnNameFormat, column.Type, GetColumnName(column.Tag))
		writeContents.Write([]byte(info + newLine))
		checkedColumns = append(checkedColumns, column)
	}
	writeContents.Write([]byte(fmt.Sprintf(newLine + "}" + newLine)))

	// 生成数据验证
	writeContents.Write([]byte(fmt.Sprintf("func (param *%s) Validate() error {", updateModelStructName) + newLine))
	writeContents.Write([]byte(`   if param == nil {
        return errors.New(consts.ParamRequired)
    }` + newLine))

	rulesContent := getValidateRules(checkedColumns)
	writeContents.Write([]byte(rulesContent + newLine))
	writeContents.Write([]byte(fmt.Sprintf("}" + newLine)))
	return writeContents.String()
}

// getValidateRules 获取验证规则
func getValidateRules(columns []*StructColumn) string {
	writeContents := strings.Builder{}
	if len(columns) < 1 {
		writeContents.Write([]byte("  return nil " + newLine))
		return writeContents.String()
	}

	writeContents.Write([]byte(`return validation.ValidateStruct(param,` + newLine))
	for _, checkedColumn := range columns {
		columnNameFormat := UnderscoreToUpperCamelCase(checkedColumn.Name)
		columnType := checkedColumn.Type
		writeContents.Write([]byte(fmt.Sprintf(`validation.Field(&param.%s,`, columnNameFormat) + newLine))

		// 整型字段验证规则
		intType := []string{"int8", "int16", "int32", "int", "int64"}
		isInt := utils.InArray(columnType, intType)
		if isInt == true {
			writeContents.Write([]byte(`  validation.Min(0).Error(consts.MinIsZero),` + newLine))
		}

		// 必填字段验证
		if checkedColumn.IsNullAble == IsNullAbleNO {
			writeContents.Write([]byte(`  validation.Required.Error(consts.ValidateColumnRequired),` + newLine))
		}

		switch columnType {
		case "int8":
			writeContents.Write([]byte(`  validation.Max(consts.CannotGreatThanInt8MaxValue).Error(),` + newLine))
		case "int16":
			writeContents.Write([]byte(`  validation.Max(consts.CannotGreatThanInt16MaxValue).Error(),` + newLine))
		case "int32", "int":
			writeContents.Write([]byte(`  validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),` + newLine))
		case "int64":
			writeContents.Write([]byte(`  validation.Max(consts.MaxInt64Value).Error(consts.CannotGreatThanInt64MaxValue),` + newLine))
		case "bool":
			writeContents.Write([]byte(`  validation.In(true,false).Error(consts.DataNotInAllowedRange)`))
		case "string":
			maxLen := getInterfaceToInt(checkedColumn.CharacterMaximumLength)
			if maxLen > 0 {
				minLen := 0
				if checkedColumn.IsNullAble == IsNullAbleNO {
					minLen = 1
				}
				lenContent := fmt.Sprintf(` validation.Length(%d, %d).Error("长度在%d到%d之间"),`, minLen, maxLen, minLen, maxLen)
				writeContents.Write([]byte(lenContent + newLine))
			}

			// email
			if strings.Contains(columnNameFormat, EmailColumn) {
				writeContents.Write([]byte(consts.EmailFormatError + newLine))
			}

		case "time.Time":

		case "float64":

		default:
		}
		writeContents.Write([]byte(")," + newLine))
	}
	writeContents.Write([]byte(`)` + newLine))
	return writeContents.String()
}

// getInterfaceToInt Interface{} 转整型
func getInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}
