package req

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	TenantID uint   `json:"tenantId" form:"tenantId" binding:"required"`
}

type LogoutReq struct {
	Token string `json:"token" binding:"required"`
}
