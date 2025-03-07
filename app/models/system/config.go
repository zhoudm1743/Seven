package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Config struct {
	model.Model
	Type  string `gorm:"default:'';comment:'类型''"`
	Name  string `gorm:"not null;default:'';comment:'键'"`
	Value string `gorm:"type:text;comment:'值'"`
	model.FormTenant
}
