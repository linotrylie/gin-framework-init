package codegenerator

import (
	"errors"
	"fmt"
	"strings"
)

// DBEnum 数据库数据库注释中文英文对照
type DBEnum struct {
	ColumnNote  string // 数据库注释,eg: 1=是,2=否
	EnglishName string // 是 ==> True,否 ==> False
}

// DBEnumList 数据表注释
var DBEnumList = []DBEnum{
	{"是", "true"},
	{"否", "false"},
	{"目录", "menu"},
	{"文件", "dir"},
	{"菜单", "dir"},
	{"按钮", "button"},
	{"大", "big"},
	{"中", "medium"},
	{"小", "small"},
	{"显示", "show"},
	{"隐藏", "hide"},
	{"启用", "enable"},
	{"禁用", "disable"},
	{"成功", "success"},
	{"失败", "failed"},
	{"待运行", "waitRun"},
}

func GetEnum(name string) (*DBEnum, error) {
	name = strings.Trim(name, "")
	if name == "" {
		return nil, nil
	}
	for _, enum := range DBEnumList {
		if enum.ColumnNote == name {
			return &enum, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("数据表注释对应关系获取失败,input = %s", name))
}
