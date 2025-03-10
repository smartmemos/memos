package db

import (
	"fmt"
	"strings"
)

// Filter 接口
type Filter interface {
	GetFields() string
	GetQuery() (string, []any)
	GetPage() int64
	GetPageSize() int64
	GetOrder() string
	HasNextPage(total int64) bool
}

// BaseFilter 实现Filter接口的结构体
type BaseFilter struct {
	Fields   string
	OrderBy  string
	Page     int64
	PageSize int64
}

// GetFields 查询字段
func (f BaseFilter) GetFields() string {
	return ""
}

// GetQuery 查询条件
func (f BaseFilter) GetQuery() (string, []any) {
	panic("method GetQuery not implemented")
}

// GetPage 页码
func (f BaseFilter) GetPage() int64 {
	return max(f.Page, 1)
}

// GetPageSize 每页大小
func (f BaseFilter) GetPageSize() int64 {
	if f.PageSize == 0 {
		f.PageSize = 15
	}
	return min(f.PageSize, 10000)
}

// HasNextPage 是否有下一页
func (f BaseFilter) HasNextPage(total int64) bool {
	return total > (f.GetPage()-1)*f.GetPageSize()
}

// GetOrder 排序
func (f BaseFilter) GetOrder() string {
	var orders []string
	for _, v := range strings.Split(f.OrderBy, ",") {
		if v = strings.TrimSpace(v); v != "" {
			order := "ASC"
			if v[0] == '-' {
				v = v[1:]
				order = "DESC"
			} else if v[0] == '+' {
				v = v[1:]
			}
			orders = append(orders, fmt.Sprintf("%s %s", EscapeField(v), order))
		}
	}
	return strings.Join(orders, ",")
}
