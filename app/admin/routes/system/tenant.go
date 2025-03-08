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

type tenantDep struct {
	fx.In
	TenantSrv system.TenantService
}

func tenantRoutes(t tenantDep, r *contracts.AdminRouter) {
	tenant := r.Group("/system/tenant")

	tenant.GET("/all", t.all)
	tenant.GET("/list", t.list)
	tenant.GET("/detail", t.detail)
	tenant.POST("/add", t.add)
	tenant.POST("/edit", t.edit)
	tenant.POST("/del", t.del)
}

func (t tenantDep) all(c *gin.Context) {
	res, err := t.TenantSrv.All(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t tenantDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.TenantSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t tenantDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.TenantSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t tenantDep) add(c *gin.Context) {
	var addReq req.SystemTenantAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.TenantSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t tenantDep) edit(c *gin.Context) {
	var editReq req.SystemTenantEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.TenantSrv.Edit(editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t tenantDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.TenantSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
