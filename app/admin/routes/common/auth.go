package common

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/admin/contracts"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/service/common"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"go.uber.org/fx"
)

type authDep struct {
	fx.In
	AuthSrv common.AuthService
}

func authRoutes(dep authDep, r *contracts.AdminRouter) {
	auth := r.Group("/common")

	auth.POST("/auth/login", dep.login)
	auth.POST("/auth/logout", dep.logout)
	auth.GET("/auth/tenant", dep.tenant)
}

func (dep authDep) login(c *gin.Context) {
	var loginReq req.LoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := dep.AuthSrv.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

func (dep authDep) logout(c *gin.Context) {
	var logoutReq req.LogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &logoutReq)) {
		return
	}
	err := dep.AuthSrv.Logout(&logoutReq)
	response.CheckAndResp(c, err)
}

func (dep authDep) tenant(c *gin.Context) {
	res, err := dep.AuthSrv.Tenant()
	response.CheckAndRespWithData(c, res, err)
}
