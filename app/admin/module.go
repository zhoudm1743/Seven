package admin

import (
	"github.com/zhoudm1743/Seven/app/admin/routes"
	"github.com/zhoudm1743/Seven/app/admin/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Module("admin"),
	routes.Module,
	service.Module,
)
