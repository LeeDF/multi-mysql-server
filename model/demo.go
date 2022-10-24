package model

import (
	"gorm.io/gorm"
)

// Demo [...]
type Demo struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

// DemoColumns get sql column name.获取数据库列名
var DemoColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
}
