package utils

import (
	"reflect"
	"sort"
)

// InArray  判断元素是否在切片或数组中
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

// InArrayWithIndex  判断元素是否在切片或数组中,并返回所在位置
func InArrayWithIndex(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

// ContainsDuplicate 整型切片是否有重复元素
func ContainsDuplicate(numbs []int) bool {
	count := map[int]int{}
	sort.Ints(numbs)
	for i := 0; i < len(numbs); i++ {
		count[numbs[i]]++
		if count[numbs[i]] > 1 {
			return true
		}
	}
	return false
}
