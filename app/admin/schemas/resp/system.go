package resp

import (
	"seven-admin/util"
	"seven-admin/util/times"
)

type SystemTenantResp struct {
	ID         uint         `json:"id" structs:"id"`
	Name       string       `json:"name" structs:"name"`
	Code       string       `json:"code" structs:"code"`
	Domain     string       `json:"domain" structs:"domain"`
	Remark     string       `json:"remark" structs:"remark"`
	Contact    string       `json:"contact" structs:"contact"`
	Phone      string       `json:"phone" structs:"phone"`
	Email      string       `json:"email" structs:"email"`
	ExpireAt   int64        `json:"expireAt" structs:"expireAt"`
	IsDisable  uint8        `json:"isDisable" structs:"isDisable"`
	CreateTime times.TsTime `json:"createTime" structs:"createTime"`
	UpdateTime times.TsTime `json:"updateTime" structs:"updateTime"`
}

type SystemMenuResp struct {
	ID       uint               `json:"id" binding:"required" form:"id"`
	ParentId uint               `json:"parent_id" binding:"required" form:"parent_id"`
	Label    string             `json:"label" binding:"required" form:"label"`
	Key      string             `json:"key" binding:"required" form:"key"`
	Type     uint               `json:"type" binding:"required" form:"type"`
	Subtitle string             `json:"subtitle" form:"subtitle"`
	OpenType uint               `json:"open_type" form:"open_type"`
	Auth     string             `json:"auth" form:"auth"`
	Path     string             `json:"path" form:"path"`
	Icon     string             `json:"icon" form:"icon"`
	Sort     uint               `json:"sort" form:"sort"`
	Button   []SystemMenuButton `json:"button" form:"button"`
}

type SystemMenuButton struct {
	Label string `json:"label" binding:"required" form:"label"`
	Auth  string `json:"auth" binding:"required" form:"auth"`
	Sort  uint   `json:"sort" binding:"required" form:"sort"`
}

// SystemRoleSimpleResp 系统角色返回简单信息
type SystemRoleSimpleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	CreateTime util.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime util.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// SystemRoleResp 系统角色返回信息
type SystemRoleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	Remark     string      `json:"remark" structs:"remark"`         // 角色备注
	Menus      []uint      `json:"menus" structs:"menus"`           // 关联菜单
	Member     int64       `json:"member" structs:"member"`         // 成员数量
	Sort       uint16      `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime util.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime util.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
