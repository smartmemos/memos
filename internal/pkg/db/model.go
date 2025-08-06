package db

import (
	"time"
)

// Model 基础结构
type Model struct {
	ID        int64     `gorm:"primary_key" json:"id,string"`                       // 主键ID
	CreatedAt time.Time `gorm:"column:created_ts;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_ts;autoUpdateTime" json:"updated_at"` // 更新时间
}
