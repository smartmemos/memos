package db

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
)

func BuildQuery(f any) (string, []any) {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	var conditions []string
	var args []any
	ns := schema.NamingStrategy{}

	for i := range t.NumField() {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if strings.HasSuffix(field.Type.PkgPath(), "/db") && strings.HasPrefix(field.Type.Name(), "F[") {
			op := fieldVal.FieldByName("Op").String()
			val := fieldVal.FieldByName("Value").Interface()
			if op != "" {
				var columnName string
				if tag := field.Tag.Get("gorm"); tag != "" {
					for _, part := range strings.Split(tag, ";") {
						if strings.HasPrefix(part, "column:") {
							columnName = strings.TrimPrefix(part, "column:")
						}
					}
				}
				if columnName == "" {
					columnName = ns.ColumnName("", field.Name)
				}
				conditions = append(conditions, fmt.Sprintf("%s %s", columnName, op))
				args = append(args, val)
			}
		}
	}
	return strings.Join(conditions, " AND "), args
}

// Filter 接口
type Filter interface {
	GetFields() string
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
	for v := range strings.SplitSeq(f.OrderBy, ",") {
		if v = strings.TrimSpace(v); v != "" {
			order := "ASC"
			switch v[0] {
			case '-':
				v = v[1:]
				order = "DESC"
			case '+':
				v = v[1:]
			}
			orders = append(orders, fmt.Sprintf("%s %s", EscapeField(v), order))
		}
	}
	return strings.Join(orders, ",")
}

// NewBaseFilter
func NewBaseFilter(page, size int, orderBy string) BaseFilter {
	return BaseFilter{
		Page:     int64(page),
		PageSize: int64(size),
		OrderBy:  orderBy,
	}
}

type F[T any] struct {
	Op    string
	Value T
}

func NewF[T any](op string, value T) F[T] {
	return F[T]{Op: op, Value: value}
}

func NewEmptyF[T any](v T) F[T] {
	return F[T]{Op: "", Value: v}
}

func Eq[T any](v T) F[T] {
	return NewF("= ?", v)
}

func NotEq[T any](v T) F[T] {
	return NewF("!= ?", v)
}

func Gt[T any](v T) F[T] {
	return NewF("> ?", v)
}

func Gte[T any](v T) F[T] {
	return NewF(">= ?", v)
}

func Lt[T any](v T) F[T] {
	return NewF("< ?", v)
}

func Lte[T any](v T) F[T] {
	return NewF("<= ?", v)
}

func Like[T any](v T) F[T] {
	return NewF("LIKE ?", v)
}

func NotLike[T any](v T) F[T] {
	return NewF("NOT LIKE ?", v)
}

func In[T any](v []T) F[[]T] {
	return NewF("IN (?)", v)
}

func NotIn[T any](v []T) F[[]T] {
	return NewF("NOT IN (?)", v)
}

func OmitEq[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Eq(v)
}

func OmitNotEq[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return NotEq(v)
}

func OmitGt[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Gt(v)
}

func OmitGte[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Gte(v)
}

func OmitLt[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Lt(v)
}

func OmitLte[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Lte(v)
}

func OmitLike[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Like(v)
}

func OmitNotLike[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return NotLike(v)
}

func OmitIn[T any](v []T) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return In(v)
}

func OmitNotIn[T any](v []T) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return NotIn(v)
}

// Omit 如果值为零值，则不添加条件
func Omit[T any](v T, fn func(T) F[T]) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return fn(v)
}

// Omits 如果值为空，则不添加条件
func Omits[T any](v []T, fn func([]T) F[[]T]) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return fn(v)
}
