package db

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
)

// BuildQuery 构建查询条件
func BuildQuery(f any) (string, []any) {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Pointer {
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
			if op := fieldVal.FieldByName("Op").String(); op != "" {
				var columnName string
				if tag := field.Tag.Get("gorm"); tag != "" {
					for part := range strings.SplitSeq(tag, ";") {
						if strings.HasPrefix(part, "column:") {
							columnName, _ = strings.CutPrefix(part, "column:")
							break
						}
					}
				}
				if columnName == "" {
					columnName = ns.ColumnName("", field.Name)
				}
				conditions = append(conditions, fmt.Sprintf("`%s` %s", columnName, op))
				val := fieldVal.FieldByName("Value").Interface()
				args = append(args, val)
			}
		}
	}
	return strings.Join(conditions, " AND "), args
}

// QueryFilter 接口
type QueryFilter interface {
	GetFields() string
	GetPage() int64
	GetPageSize() int64
	GetOrder() string
	HasNextPage(total int64) bool
}

// Query 实现Filter接口的结构体
type Query struct {
	fields   string
	orderBy  string
	page     int64
	pageSize int64
}

// GetFields 查询字段
func (q Query) GetFields() string {
	return ""
}

// GetPage 页码
func (q Query) GetPage() int64 {
	if q.page < 1 {
		q.page = DefaultPage
	}
	return q.page
}

// GetPageSize 每页大小
func (q Query) GetPageSize() int64 {
	if q.pageSize < 1 {
		q.pageSize = MaxPageSize
	} else if q.pageSize > MaxPageSize {
		q.pageSize = MaxPageSize
	}
	return q.pageSize
}

// HasNextPage 是否有下一页
func (q Query) HasNextPage(total int64) bool {
	return total > (q.GetPage()-1)*q.GetPageSize()
}

// GetOrder 排序
func (q Query) GetOrder() string {
	var orders []string
	for v := range strings.SplitSeq(q.orderBy, ",") {
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

const (
	// DefaultPage 默认页码
	DefaultPage = 1

	// DefaultPageSize 默认每页大小
	DefaultPageSize = 15

	// MaxPageSize 单次查询最大数量
	MaxPageSize = 10000
)

// NewQuery 创建一个查询Query
func NewQuery(opts ...QueryOption) Query {
	q := Query{page: DefaultPage, pageSize: DefaultPageSize}
	for _, opt := range opts {
		opt(&q)
	}
	return q
}

// NewQueryAll 创建一个查询所有数据的Query
func NewQueryAll(opts ...QueryOption) Query {
	return NewQuery(append(opts, WithPageSize(MaxPageSize))...)
}

// QueryOption 查询选项
type QueryOption func(q *Query)

// WithFields 设置查询字段
func WithFields(fields string) QueryOption {
	return func(q *Query) {
		q.fields = fields
	}
}

// WithPage 设置页码
func WithPage(page int) QueryOption {
	return func(q *Query) {
		q.page = int64(page)
	}
}

// WithPageSize 设置每页大小
func WithPageSize(size int) QueryOption {
	return func(q *Query) {
		q.pageSize = int64(size)
	}
}

// WithOrderBy 设置排序
func WithOrderBy(orderBy string) QueryOption {
	return func(q *Query) {
		q.orderBy = orderBy
	}
}

// FilterOption 过滤选项
type FilterOption func(f *Query)

// F 过滤条件
type F[T any] struct {
	Op    string
	Value T
}

// NewF 创建一个过滤条件
func NewF[T any](op string, value T) F[T] {
	return F[T]{Op: op, Value: value}
}

// NewEmptyF 创建一个空过滤条件
func NewEmptyF[T any](v T) F[T] {
	return F[T]{Op: "", Value: v}
}

// Eq 等于
func Eq[T any](v T) F[T] {
	return NewF("= ?", v)
}

// NotEq 不等于
func NotEq[T any](v T) F[T] {
	return NewF("!= ?", v)
}

// Gt 大于
func Gt[T any](v T) F[T] {
	return NewF("> ?", v)
}

// Gte 大于等于
func Gte[T any](v T) F[T] {
	return NewF(">= ?", v)
}

// Lt 小于
func Lt[T any](v T) F[T] {
	return NewF("< ?", v)
}

// Lte 小于等于
func Lte[T any](v T) F[T] {
	return NewF("<= ?", v)
}

// Like 模糊匹配
func Like[T any](v T) F[T] {
	return NewF("LIKE ?", v)
}

// NotLike 不模糊匹配
func NotLike[T any](v T) F[T] {
	return NewF("NOT LIKE ?", v)
}

// In 包含
func In[T any](v []T) F[[]T] {
	return NewF("IN (?)", v)
}

// NotIn 不包含
func NotIn[T any](v []T) F[[]T] {
	return NewF("NOT IN (?)", v)
}

// OmitEq 如果不为零值，则等于
func OmitEq[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Eq(v)
}

// OmitNotEq 如果不为零值，则不等于
func OmitNotEq[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return NotEq(v)
}

// OmitGt 如果不为零值，则大于
func OmitGt[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Gt(v)
}

// OmitGte 如果不为零值，则大于等于
func OmitGte[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Gte(v)
}

// OmitLt 如果不为零值，则小于
func OmitLt[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Lt(v)
}

// OmitLte 如果不为零值，则小于等于
func OmitLte[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Lte(v)
}

// OmitLike 如果不为零值，则模糊匹配
func OmitLike[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return Like(v)
}

// OmitNotLike 如果不为零值，则不模糊匹配
func OmitNotLike[T any](v T) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return NotLike(v)
}

// OmitIn 如果不为零值，则包含
func OmitIn[T any](v []T) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return In(v)
}

// OmitNotIn 如果不为零值，则不包含
func OmitNotIn[T any](v []T) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return NotIn(v)
}

// Omit 如果不为零值，则添加条件
func Omit[T any](v T, fn func(T) F[T]) F[T] {
	if reflect.ValueOf(v).IsZero() {
		return NewEmptyF(v)
	}
	return fn(v)
}

// Omits 如果不为空，则添加条件
func Omits[T any](v []T, fn func([]T) F[[]T]) F[[]T] {
	if len(v) == 0 {
		return NewEmptyF(v)
	}
	return fn(v)
}
