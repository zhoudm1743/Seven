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

type postDep struct {
	fx.In
	PostSrv system.PostService
}

func postRoutes(t postDep, r *contracts.AdminRouter) {
	post := r.Group("/system/post")

	post.GET("/all", t.all)
	post.GET("/list", t.list)
	post.GET("/detail", t.detail)
	post.POST("/add", t.add)
	post.POST("/edit", t.edit)
	post.POST("/del", t.del)
}

func (t postDep) all(c *gin.Context) {
	res, err := t.PostSrv.All(query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t postDep) list(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.PostSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t postDep) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.PostSrv.Detail(detailReq.ID, query.GetAuthReq(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t postDep) add(c *gin.Context) {
	var addReq req.SystemPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.PostSrv.Add(addReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t postDep) edit(c *gin.Context) {
	var editReq req.SystemPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.PostSrv.Edit(editReq, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}

func (t postDep) del(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.PostSrv.Del(delReq.ID, query.GetAuthReq(c))
	response.CheckAndResp(c, err)
}
