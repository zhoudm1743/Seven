package routes

import (
	"github.com/zhoudm1743/Seven/app/admin/contracts"
	"github.com/zhoudm1743/Seven/app/admin/routes/common"
	"github.com/zhoudm1743/Seven/app/admin/routes/system"
	"github.com/zhoudm1743/Seven/app/admin/routes/test"
	web "github.com/zhoudm1743/Seven/pkg/common/HttpServer"
	"github.com/zhoudm1743/Seven/pkg/common/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module("adminRoutes",
	fx.Provide(NewAdminRouter),
	// 注册子路由模块
	test.Module,
	// 注册路由
	common.Module,
	system.Module,
)

type RouterDeps struct {
	fx.In
	Http *web.HttpServer
}

func NewAdminRouter(deps RouterDeps) *contracts.AdminRouter {
	return &contracts.AdminRouter{
		RouterGroup: deps.Http.Gin.Group("/admin", middleware.AuthCheck()),
	}
}
