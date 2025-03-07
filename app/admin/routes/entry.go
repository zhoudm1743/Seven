package routes

import (
	"github.com/zhoudm1743/Seven/app/admin/contracts"
	"github.com/zhoudm1743/Seven/app/admin/routes/test"
	web "github.com/zhoudm1743/Seven/pkg/common/HttpServer"
	"go.uber.org/fx"
)

var Module = fx.Module("adminRoutes",
	fx.Provide(NewAdminRouter),
	// 注册子路由模块
	test.Module,
)

type RouterDeps struct {
	fx.In
	Http *web.HttpServer
}

func NewAdminRouter(deps RouterDeps) *contracts.AdminRouter {
	return &contracts.AdminRouter{
		deps.Http.Gin.Group("/admin"),
	}
}
