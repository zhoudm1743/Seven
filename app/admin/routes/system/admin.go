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

type adminDep struct {
	fx.In
	AdminSrv system.AdminService
}

func adminRoutes(t adminDep, r *contracts.AdminRouter) {
	admin := r.Group("/system/admin")

	admin.GET("/self", t.self)
	admin.GET("/list", t.list)
	admin.GET("/detail", t.detail)
	admin.POST("/add", t.add)
	admin.POST("/edit", t.edit)
	admin.POST("/upInfo", t.upInfo)
	admin.POST("/del", t.del)
	admin.POST("/disable", t.disable)
}

func (t adminDep) self(c *gin.Context) {
	res, err := t.AdminSrv.Self(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t adminDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.AdminSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t adminDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.AdminSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t adminDep) add(c *gin.Context) {
	var addReq req.SystemAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.AdminSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t adminDep) edit(c *gin.Context) {
	var editReq req.SystemAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.AdminSrv.Edit(c, editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t adminDep) upInfo(c *gin.Context) {
	var upInfoReq req.SystemAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &upInfoReq)) {
		return
	}
	err := t.AdminSrv.Update(c, upInfoReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t adminDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.AdminSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t adminDep) disable(c *gin.Context) {
	var disableReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &disableReq)) {
		return
	}
	err := t.AdminSrv.Disable(disableReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
