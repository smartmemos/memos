package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var dbConn *gorm.DB

// Init 初始化
func Init(cfg Config) {
	var err error
	switch cfg.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?collation=utf8mb4_general_ci&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
}

// EscapeField 字段名称加上双引号，防止字段名与SQL关键字相同导致查询错误。
func EscapeField(field string) string {
	if strings.Contains(field, "(") { // 字段是一个函数，不需要转义
		return field
	} else if strings.Contains(field, "->") { // json_field->>'$.name' 或 json_field->'$.name'
		return field
	} else if arr := strings.Split(field, "."); len(arr) == 2 { // 字段名是 a.id 类型
		return fmt.Sprintf("%s.`%s`", arr[0], arr[1])
	} else {
		return fmt.Sprintf("`%s`", field) // 默认直接加上双引号``
	}
}

// GetDB 返回实例
func GetDB(ctx context.Context) *gorm.DB {
	return dbConn.WithContext(ctx).Debug()
}

// Create 新增记录
func Create(ctx context.Context, v any) error {
	return GetDB(ctx).Create(v).Error
}

// Save 保存记录
func Save(ctx context.Context, v any) error {
	return GetDB(ctx).Save(v).Error
}

// HasNextPage 是否有下一页
func HasNextPage(total int64, page int64, pageSize int64) bool {
	return total > (page-1)*pageSize
}

// HasRecrods 当前页码是否有记录
func HasRecrods(total int64, page int64, pageSize int64) bool {
	if total == 0 {
		return false
	} else if pageSize > 0 && page > 0 {
		return total > (page-1)*pageSize
	} else {
		return true
	}
}

// Updates 更新数据
func Updates(ctx context.Context, m schema.Tabler, f Filter, v any) (int64, error) {
	query, args := BuildQuery(f)
	if query == "" {
		return 0, errors.New("更新条件不能为空")
	}
	res := GetDB(ctx).Where(query, args...).Model(m).Updates(v)
	return res.RowsAffected, res.Error
}

// Update 更新数据
func Update(ctx context.Context, m schema.Tabler, v any) error {
	return GetDB(ctx).Model(m).Updates(v).Error
}

// Delete 删除记录
func Delete(ctx context.Context, m schema.Tabler, f Filter) (int64, error) {
	query, args := BuildQuery(f)
	if query == "" {
		return 0, errors.New("删除条件不能为空")
	}
	res := GetDB(ctx).Where(query, args...).Delete(m)
	return res.RowsAffected, res.Error
}

// Count 计数
func Count(ctx context.Context, v schema.Tabler, f Filter) (total int64, err error) {
	query, args := BuildQuery(f)
	err = GetDB(ctx).Where(query, args...).Model(v).Count(&total).Error
	return
}

// FindOne 只查一条记录
func FindOne(ctx context.Context, f Filter, v any) error {
	query, args := BuildQuery(f)
	err := GetDB(ctx).
		Where(query, args...).
		Order(f.GetOrder()).
		First(v).Error
	return err
}

// Find 查多条记录
func Find(ctx context.Context, f Filter, v any) error {
	query, args := BuildQuery(f)
	err := GetDB(ctx).
		Where(query, args...).
		Offset(int((f.GetPage() - 1) * f.GetPageSize())).
		Limit(int(f.GetPageSize())).
		Order(f.GetOrder()).
		Find(v).Error
	return err
}

func IsRecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

func IsDbError(err error) bool {
	return err != nil && err != gorm.ErrRecordNotFound
}
