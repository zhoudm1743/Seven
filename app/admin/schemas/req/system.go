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

// SystemAdminListReq 管理员列表参数
type SystemAdminListReq struct {
	Username string `form:"username"`        // 账号
	Nickname string `form:"nickname"`        // 昵称
	Role     *int   `form:"role,default=-1"` // 角色ID
	TenantId *uint  `form:"tenantId"`        // 企业ID
}

// SystemAdminAddReq 管理员新增参数
type SystemAdminAddReq struct {
	DeptId       uint   `form:"deptId" json:"dept_id" binding:"gte=0"`                    // 部门ID
	PostId       uint   `form:"postId" json:"post_id" binding:"gte=0"`                    // 岗位ID
	Username     string `form:"username" json:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" json:"password" binding:"required"`              // 密码
	Avatar       string `form:"avatar" json:"avatar" binding:"required"`                  // 头像
	RoleId       uint   `form:"role" json:"role" binding:"gte=0"`                         // 角色
	Sort         int    `form:"sort" json:"sort" binding:"gte=0"`                         // 排序
	IsDisable    uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`           // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" json:"isMultipoint" binding:"oneof=0 1"`     // 多端登录: [0=否, 1=是]
}

// SystemAdminEditReq 管理员编辑参数
type SystemAdminEditReq struct {
	ID           uint   `form:"id" json:"id" binding:"required,gt=0"`                     // 主键
	DeptId       uint   `form:"deptId" json:"dept_id" binding:"gte=0"`                    // 部门ID
	PostId       uint   `form:"postId" json:"post_id" binding:"gte=0"`                    // 岗位ID
	Username     string `form:"username" json:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" json:"password"`                                 // 密码
	Avatar       string `form:"avatar" json:"avatar"`                                     // 头像
	RoleId       uint   `form:"role" json:"role" binding:"gte=0"`                         // 角色
	Sort         int    `form:"sort" json:"sort" binding:"gte=0"`                         // 排序
	IsDisable    uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`           // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" json:"isMultipoint" binding:"oneof=0 1"`     // 多端登录: [0=否, 1=是]
}

// SystemAdminUpdateReq 管理员更新参数
type SystemAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"`     // 昵称
	Avatar       string `form:"avatar"`                                       // 头像
	Password     string `form:"password" binding:"required"`                  // 密码
	CurrPassword string `form:"currPassword" binding:"required,min=6,max=32"` // 密码
}

// SystemPostListReq 岗位列表参数
type SystemPostListReq struct {
	Code      string `form:"code"`                                        // 岗位编码
	Name      string `form:"name"`                                        // 岗位名称
	IsDisable int8   `form:"isDisable,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

// SystemPostAddReq 岗位新增参数
type SystemPostAddReq struct {
	Code      string `form:"code" json:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name      string `form:"name" json:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks   string `form:"remarks" json:"remarks" binding:"max=250"`          // 岗位备注
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`    // 是否停用: [0=否, 1=是]
	Sort      int    `form:"sort" json:"sort" binding:"gte=0"`                  // 排序编号
}

// SystemPostEditReq 岗位编辑参数
type SystemPostEditReq struct {
	ID        string `form:"id" json:"id" binding:"required,gt=0"`              // 主键
	Code      string `form:"code" json:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name      string `form:"name" json:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks   string `form:"remarks" json:"remarks" binding:"max=250"`          // 岗位备注
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`    // 是否停用: [0=否, 1=是]
	Sort      int    `form:"sort" json:"sort" binding:"gte=0"`                  // 排序编号
}

// SystemDeptListReq 部门列表参数
type SystemDeptListReq struct {
	Name      string `form:"name"`                                        // 部门名称
	IsDisable int8   `form:"isDisable,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

// SystemDeptAddReq 部门新增参数
type SystemDeptAddReq struct {
	Pid       uint   `form:"pid" json:"pid" binding:"gte=0"`                    // 部门父级
	Name      string `form:"name" json:"name" binding:"required,min=1,max=100"` // 部门名称
	Duty      string `form:"duty" json:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile    string `form:"mobile" json:"mobile" binding:"omitempty,len=11"`   // 联系电话
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`    // 是否停用: [0=否, 1=是]
	Sort      int    `form:"sort" json:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

// SystemDeptEditReq 部门编辑参数
type SystemDeptEditReq struct {
	ID        uint   `form:"id" json:"id" binding:"required,gt=0"`              // 主键
	Pid       uint   `form:"pid" json:"pid" binding:"gte=0"`                    // 部门父级
	Name      string `form:"name" json:"name" binding:"required,min=1,max=100"` // 部门名称
	Duty      string `form:"duty" json:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile    string `form:"mobile" json:"mobile" binding:"omitempty,len=11"`   // 联系电话
	IsDisable uint8  `form:"isDisable" json:"isDisable" binding:"oneof=0 1"`    // 是否停用: [0=否, 1=是]
	Sort      int    `form:"sort" json:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}
