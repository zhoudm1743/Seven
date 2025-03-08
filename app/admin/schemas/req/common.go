package req

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	TenantID uint   `json:"tenant_id" binding:"required"`
	Captcha  string `json:"captcha" binding:"required"`
}

type LogoutReq struct {
	Token string `json:"token" binding:"required"`
}
