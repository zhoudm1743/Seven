package system

import (
	model "github.com/zhoudm1743/Seven/app/models"
)

type Menu struct {
	model.Model
	Pid           uint   `gorm:"column:pid"`
	Name          string `gorm:"column:name"`
	Path          string `gorm:"column:path"`
	Redirect      string `gorm:"column:redirect"`
	ComponentPath string `gorm:"column:componentPath"`
	IsDisabled    bool   `gorm:"column:isDisabled;default:false"`
	Icon          string `gorm:"column:icon"`
	MenuType      string `json:"menuType" gorm:"column:menuType"`                       // dir or page
	Title         string `json:"title" gorm:"column:title"`                             // 页面标题
	RequiresAuth  bool   `json:"requiresAuth" gorm:"column:requiresAuth;default:false"` // 是否需要登录权限
	KeepAlive     bool   `json:"keepAlive" gorm:"column:keepAlive;default:false"`       // 是否开启页面缓存
	Hide          bool   `json:"hide" gorm:"column:hide;default:false"`                 // 不显示在菜单中
	Sort          uint   `json:"sort" gorm:"column:sort;default:0"`                     // 排序
	Href          string `json:"href" gorm:"column:href"`                               // 嵌套外链
	ActiveMenu    string `json:"activeMenu" gorm:"column:activeMenu"`                   // 当前路由高亮菜单
	WithoutTab    bool   `json:"withoutTab" gorm:"column:withoutTab;default:false"`     // 不添加到Tab
	PinTab        bool   `json:"pinTab" gorm:"column:pinTab;default:false"`             // 固定Tab
}

// TableName 表名
func (Menu) TableName() string {
	return "sys_menu"
}
