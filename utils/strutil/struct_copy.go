package strutil

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
)

// StructAssign 结构体复制
// binding 需要复制的参数
// value 被复制的参数
// 字段全部相同才可
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() // 获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   // 获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段,有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// Copy 不需要字段全部相同
// 原理就是两个JSON互相转换
func Copy(to, from interface{}) error {
	// 先将被复制结构体转为json
	target, err := json.Marshal(from)
	if err != nil {
		return errors.Wrap(err, "待复制结构体转为json出错")
	}
	err = json.Unmarshal(target, to)
	if err != nil {
		return errors.Wrap(err, "结构体复制出错")
	}
	return nil
}
