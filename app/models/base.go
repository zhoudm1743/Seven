package model

import (
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID         uint  `gorm:"primarykey;comment:'主键'" json:"id"`
	CreateTime int64 `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"createTime"`
	UpdateTime int64 `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"updateTime"`
}

type SoftDelete struct {
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag;default:0;comment:'删除标识'" json:"-"`
}

type FormTenant struct {
	TenantID uint `gorm:"not null;default:0;comment:'租户ID'"`
}
