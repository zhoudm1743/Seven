package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Tenant struct {
	model.Model
	Name      string `gorm:"column:name;not null;default:'';comment:租户名称" json:"name"`
	Code      string `gorm:"column:code;not null;default:'';comment:租户编码" json:"code"`
	Domain    string `gorm:"column:domain;not null;default:'';comment:租户域名" json:"domain"`
	Remark    string `gorm:"column:remark;not null;default:'';comment:租户备注" json:"remark"`
	Contact   string `gorm:"column:contact;not null;default:'';comment:租户联系人" json:"contact"`
	Phone     string `gorm:"column:phone;not null;default:'';comment:租户联系电话" json:"phone"`
	Email     string `gorm:"column:email;not null;default:'';comment:租户联系邮箱" json:"email"`
	ExpireAt  int64  `gorm:"column:expire_at;not null;default:0;comment:租户过期时间" json:"expire_at"`
	IsDisable uint8  `gorm:"column:is_disable;not null;default:0;comment:是否禁用" json:"is_disable"`
}

// TableName 表名
func (Tenant) TableName() string {
	return "sys_tenant"
}

type TenantPerm struct {
	ID       string `gorm:"primaryKey"`
	TenantID uint   `gorm:"column:tenant_id"`
	MenuID   uint   `gorm:"column:menu_id"`
}

func (TenantPerm) TableName() string {
	return "sys_tenant_perm"
}
