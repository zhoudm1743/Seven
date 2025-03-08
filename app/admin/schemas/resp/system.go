package resp

import (
	"github.com/zhoudm1743/Seven/pkg/util"
	"github.com/zhoudm1743/Seven/pkg/util/times"
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

// SystemAdminResp 管理员返回信息
type SystemAdminResp struct {
	ID            uint        `json:"id" structs:"id"`             // 主键
	Username      string      `json:"username" structs:"username"` // 账号
	Nickname      string      `json:"nickname" structs:"nickname"` // 昵称
	Avatar        string      `json:"avatar" structs:"avatar"`     // 头像
	RoleId        uint        `json:"roleId" structs:"roleId"`     // 角色ID
	DeptId        uint        `json:"deptId" structs:"deptId"`     // 部门ID
	PostId        uint        `json:"postId" structs:"postId"`     // 岗位ID
	Dept          string      `json:"dept" structs:"dept"`         // 部门
	Tenant        string      `json:"tenant" structs:"tenant"`
	IsMultipoint  uint8       `json:"isMultipoint" structs:"isMultipoint"` // 多端登录: [0=否, 1=是]
	IsDisable     uint8       `json:"isDisable" structs:"isDisable"`       // 是否禁用: [0=否, 1=是]
	TenantId      uint        `json:"tenantId" structs:"tenantId"`
	LastLoginIp   string      `json:"lastLoginIp" structs:"lastLoginIp"`     // 最后登录IP
	LastLoginTime util.TsTime `json:"lastLoginTime" structs:"lastLoginTime"` // 最后登录时间
	CreateTime    util.TsTime `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime    util.TsTime `json:"updateTime" structs:"updateTime"`       // 更新时间
}

// SystemAdminSelfOneResp 当前管理员返回部分信息
type SystemAdminSelfOneResp struct {
	ID            uint        `json:"id" structs:"id"`             // 主键
	Username      string      `json:"username" structs:"username"` // 账号
	Nickname      string      `json:"nickname" structs:"nickname"` // 昵称
	Avatar        string      `json:"avatar" structs:"avatar"`     // 头像
	RoleId        uint        `json:"roleId" structs:"roleId"`     // 角色ID
	DeptId        uint        `json:"deptId" structs:"deptId"`     // 部门ID
	PostId        uint        `json:"postId" structs:"postId"`     // 岗位ID
	Dept          string      `json:"dept" structs:"dept"`         // 部门
	Tenant        string      `json:"tenant" structs:"tenant"`
	TenantId      uint        `json:"tenantId" structs:"tenantId"`
	IsMultipoint  uint8       `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     uint8       `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	LastLoginIp   string      `json:"lastLoginIp" structs:"lastLoginIp"`     // 最后登录IP
	LastLoginTime util.TsTime `json:"lastLoginTime" structs:"lastLoginTime"` // 最后登录时间
	SoftSuper     bool        `json:"softSuper" structs:"softSuper"`
	SuperTenant   bool        `json:"superTenant" structs:"superTenant"`
	CreateTime    util.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime    util.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// SystemAdminSelfResp 当前系统管理员返回信息
type SystemAdminSelfResp struct {
	User        SystemAdminSelfOneResp `json:"user" structs:"user"`               // 用户信息
	Permissions []string               `json:"permissions" structs:"permissions"` // 权限集合: [[*]=>所有权限, ['article:add']=>部分权限]
}

// SystemDeptResp 系统部门返回信息
type SystemDeptResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Pid        uint        `json:"pid" structs:"pid"`               // 部门父级
	Name       string      `json:"name" structs:"name"`             // 部门名称
	Duty       string      `json:"duty" structs:"duty"`             // 负责人
	Mobile     string      `json:"mobile" structs:"mobile"`         // 联系电话
	Sort       uint16      `json:"sort" structs:"sort"`             // 排序编号
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否停用: [0=否, 1=是]
	CreateTime util.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime util.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// SystemPostResp 系统岗位返回信息
type SystemPostResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Code       string      `json:"code" structs:"code"`             // 岗位编号
	Name       string      `json:"name" structs:"name"`             // 岗位名称
	Remarks    string      `json:"remarks" structs:"remarks"`       // 岗位备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 岗位排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否停用: [0=否, 1=是]
	CreateTime util.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime util.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
