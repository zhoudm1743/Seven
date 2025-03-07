package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Role struct {
	model.Model
	Name      string `gorm:"not null;default:'';comment:'角色名称''"`
	Remark    string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort      uint16 `gorm:"not null;default:0;comment:'角色排序'"`
	IsAdmin   uint8  `gorm:"not null;default:0;comment:'是否管理员'"`
	model.FormTenant
}

// TableName 表名
func (Role) TableName() string {
	return "sys_role"
}

type RolePerm struct {
	ID     string `gorm:"primaryKey"`
	RoleID uint   `gorm:"column:role_id"`
	MenuID uint   `gorm:"column:menu_id"`
}

func (RolePerm) TableName() string {
	return "sys_role_perm"
}
