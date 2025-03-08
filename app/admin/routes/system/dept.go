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

type deptDep struct {
	fx.In
	DeptSrv system.DeptService
}

func deptRoutes(t deptDep, r *contracts.AdminRouter) {
	dept := r.Group("/system/dept")

	dept.GET("/all", t.all)
	dept.GET("/list", t.list)
	dept.GET("/detail", t.detail)
	dept.POST("/add", t.add)
	dept.POST("/edit", t.edit)
	dept.POST("/del", t.del)
}

func (t deptDep) all(c *gin.Context) {
	res, err := t.DeptSrv.All(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t deptDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.DeptSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t deptDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.DeptSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t deptDep) add(c *gin.Context) {
	var addReq req.SystemDeptAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.DeptSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t deptDep) edit(c *gin.Context) {
	var editReq req.SystemDeptEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.DeptSrv.Edit(editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t deptDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.DeptSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
