package resp

type LoginResp struct {
	Token    string                 `json:"token"`
	UserInfo SystemAdminSelfOneResp `json:"userInfo"`
}
