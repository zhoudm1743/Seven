package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Dept struct {
	model.Model
	Pid       uint   `gorm:"not null;default:0;comment:'上级主键'"`
	Name      string `gorm:"not null;default:'';comment:'部门名称''"`
	Duty      string `gorm:"not null;default:'';comment:'负责人名'"`
	Mobile    string `gorm:"not null;default:'';comment:'联系电话'"`
	Sort      uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否禁用'"`
	model.FormTenant
	model.SoftDelete
}

func (Dept) TableName() string {
	return "sys_dept"
}
