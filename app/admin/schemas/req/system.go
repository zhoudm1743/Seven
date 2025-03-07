package req

type SystemTenantAddReq struct {
	Name      string `json:"name" binding:"required" form:"name"`
	Code      string `json:"code" binding:"required" form:"code"`
	Domain    string `json:"domain"  form:"domain"`
	Remark    string `json:"remark"  form:"remark"`
	Contact   string `json:"contact"  form:"contact"`
	Phone     string `json:"phone"  form:"phone"`
	Email     string `json:"email"  form:"email"`
	ExpireAt  int64  `json:"expireAt"  form:"expireAt"`
	IsDisable uint8  `json:"isDisable"  form:"isDisable"`
}

type SystemTenantEditReq struct {
	ID        uint   `json:"id" binding:"required" form:"id"`
	Name      string `json:"name" binding:"required" form:"name"`
	Code      string `json:"code" binding:"required" form:"code"`
	Domain    string `json:"domain"  form:"domain"`
	Remark    string `json:"remark"  form:"remark"`
	Contact   string `json:"contact"  form:"contact"`
	Phone     string `json:"phone"  form:"phone"`
	Email     string `json:"email"  form:"email"`
	ExpireAt  int64  `json:"expireAt"  form:"expireAt"`
	IsDisable uint8  `json:"isDisable"  form:"isDisable"`
}

type SystemTenantQueryReq struct {
	Name      string `json:"name" form:"name"`
	Code      string `json:"code" form:"code"`
	Domain    string `json:"domain"  form:"domain"`
	Contact   string `json:"contact"  form:"contact"`
	IsDisable int8   `json:"isDisable"  form:"isDisable" default:"-1"`
}

type SystemMenuAddReq struct {
	ParentId uint               `json:"parentId" binding:"required" form:"parentId"`
	Label    string             `json:"label" binding:"required" form:"label"`
	Key      string             `json:"key" binding:"required" form:"key"`
	Type     uint               `json:"type" binding:"required" form:"type"`
	Subtitle string             `json:"subtitle" form:"subtitle"`
	OpenType uint               `json:"openType" form:"openType"`
	Auth     string             `json:"auth" form:"auth"`
	Path     string             `json:"path" form:"path"`
	Icon     string             `json:"icon" form:"icon"`
	Sort     uint               `json:"sort" form:"sort"`
	Button   []SystemMenuButton `json:"button" form:"button"`
}
type SystemMenuEditReq struct {
	ID       uint               `json:"id" binding:"required" form:"id"`
	ParentId uint               `json:"parentId" binding:"required" form:"parentId"`
	Label    string             `json:"label" binding:"required" form:"label"`
	Key      string             `json:"key" binding:"required" form:"key"`
	Type     uint               `json:"type" binding:"required" form:"type"`
	Subtitle string             `json:"subtitle" form:"subtitle"`
	OpenType uint               `json:"openType" form:"openType"`
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

// SystemRoleAddReq 新增角色参数
type SystemRoleAddReq struct {
	Name      string `form:"name" json:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" json:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`   // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" json:"remark" binding:"max=200"`           // 角色备注
	MenuIds   string `form:"menuIds" json:"menuIds"`                           // 关联菜单
}

// SystemRoleEditReq 编辑角色参数
type SystemRoleEditReq struct {
	ID        uint   `form:"id" json:"id" binding:"required,gt=0"`             // 主键
	Name      string `form:"name" json:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" json:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`   // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" json:"remark" binding:"max=200"`           // 角色备注
	MenuIds   string `form:"menuIds" json:"menuIds"`                           // 关联菜单
}
