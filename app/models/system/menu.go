package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

/*
*

	{
	      label: 'Dashboard',
	      key: 'dashboard',
	      type: 1,
	      subtitle: 'dashboard',
	      openType: 1,
	      auth: 'dashboard',
	      path: '/dashboard',
	      children: [
	        {
	          label: '主控台',
	          key: 'console',
	          type: 1,
	          subtitle: 'console',
	          openType: 1,
	          auth: 'console',
	          path: '/dashboard/console',
	        },
	        {
	          label: '工作台',
	          key: 'workplace',
	          type: 1,
	          subtitle: 'workplace',
	          openType: 1,
	          auth: 'workplace',
	          path: '/dashboard/workplace',
	        },
	      ],
	    },
*/
type Menu struct {
	model.Model
	ParentID uint   `gorm:"column:parent_id"`
	Label    string `gorm:"column:label"`
	Key      string `gorm:"column:key"`
	Type     uint   `gorm:"column:type"`
	Subtitle string `gorm:"column:subtitle"`
	OpenType uint   `gorm:"column:open_type"`
	Auth     string `gorm:"column:auth"`
	Path     string `gorm:"column:path"`
	Icon     string `gorm:"column:icon"`
	Sort     uint   `gorm:"column:sort"`
	Children []Menu `gorm:"-"`
}

// TableName 表名
func (Menu) TableName() string {
	return "sys_menu"
}
