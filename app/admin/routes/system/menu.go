package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/admin/contracts"
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/service/system"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"go.uber.org/fx"
)

type menuDep struct {
	fx.In
	MenuSrv system.MenuService
}

func menuRoutes(t menuDep, r *contracts.AdminRouter) {
	menu := r.Group("/system/menu")

	menu.GET("/route", t.route)
	menu.GET("/list", t.list)
	menu.GET("/detail", t.detail)
	menu.POST("/add", t.add)
	menu.POST("/edit", t.edit)
	menu.POST("/del", t.del)
}

func (t menuDep) route(c *gin.Context) {
	res, err := t.MenuSrv.SelectMenuByRoleId(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menuDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.MenuSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menuDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.MenuSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menuDep) add(c *gin.Context) {
	var addReq req.SystemMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.MenuSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t menuDep) edit(c *gin.Context) {
	var editReq req.SystemMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.MenuSrv.Edit(editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t menuDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.MenuSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
