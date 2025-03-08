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

type roleDep struct {
	fx.In
	RoleSrv system.RoleService
}

func roleRoutes(t roleDep, r *contracts.AdminRouter) {
	role := r.Group("/system/role")

	role.GET("/all", t.all)
	role.GET("/list", t.list)
	role.GET("/detail", t.detail)
	role.POST("/add", t.add)
	role.POST("/edit", t.edit)
	role.POST("/del", t.del)
}

func (t roleDep) all(c *gin.Context) {
	res, err := t.RoleSrv.All(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t roleDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.RoleSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t roleDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.RoleSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t roleDep) add(c *gin.Context) {
	var addReq req.SystemRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.RoleSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t roleDep) edit(c *gin.Context) {
	var editReq req.SystemRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.RoleSrv.Edit(editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t roleDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.RoleSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
