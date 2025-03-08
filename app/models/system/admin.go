package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Admin struct {
	model.Model
	Username      string  `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname      string  `gorm:"not null;default:'';comment:'用户昵称'"`
	Password      string  `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar        string  `gorm:"not null;default:'';comment:'用户头像'"`
	RoleId        uint    `gorm:"not null;default:0;comment:'角色ID'"`
	PostId        uint    `gorm:"not null;default:0;comment:'岗位ID'"`
	DeptId        uint    `gorm:"not null;default:0;comment:'部门ID'"`
	Sort          uint16  `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint  uint8   `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable     uint8   `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	LastLoginIp   string  `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime int64   `gorm:"not null;default:0;comment:'最后登录时间'"`
	Post          *Post   `gorm:"foreignKey:PostId;references:ID"`
	Dept          *Dept   `gorm:"foreignKey:DeptId;references:ID"`
	Role          *Role   `gorm:"foreignKey:RoleId;references:ID"`
	Tenant        *Tenant `gorm:"-"`
	model.FormTenant
	model.SoftDelete
}

// TableName 表名
func (Admin) TableName() string {
	return "sys_admin"
}
