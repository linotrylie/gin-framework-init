package commonres

import (
	"equity/global"
)

// GetPageSize 获取每页数据条数
func GetPageSize(pageSize int) int {
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

// GetPageOffset 获取分页偏移量
func GetPageOffset(page, pageSize int) int {
	if pageSize == 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

// GetPager 获取分页和offset
func GetPager(page, pageSize int) (limit, offset int) {
	return GetPageSize(pageSize), GetPageOffset(page, pageSize)
}

// LikeSearch 模糊搜索
func LikeSearch(kwd string) string {
	return "%" + kwd + "%"
}
