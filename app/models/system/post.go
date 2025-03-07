package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Post struct {
	model.Model
	Code      string `gorm:"not null;default:'';comment:'岗位编码''"`
	Name      string `gorm:"not null;default:'';comment:'岗位名称''"`
	Remarks   string `gorm:"not null;default:'';comment:'岗位备注''"`
	Sort      uint16 `gorm:"not null;default:0;comment:'岗位排序'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'"`
	model.FormTenant
	model.SoftDelete
}

// TableName 表名
func (Post) TableName() string {
	return "sys_post"
}
