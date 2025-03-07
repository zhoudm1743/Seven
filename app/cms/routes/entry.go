package routes

import (
	"github.com/zhoudm1743/Seven/app/cms/contracts"
	"github.com/zhoudm1743/Seven/app/cms/routes/test"
	web "github.com/zhoudm1743/Seven/pkg/common/HttpServer"
	"go.uber.org/fx"
)

var Module = fx.Module("cmsRoutes",
	fx.Provide(NewCmsRouter),
	// 注册子路由模块
	test.Module,
)

type CmsRouterDeps struct {
	fx.In
	Http *web.HttpServer
}

// ==============================
func NewCmsRouter(deps CmsRouterDeps) *contracts.CmsRouter {
	return &contracts.CmsRouter{deps.Http.Gin.Group("/cms")}
}
