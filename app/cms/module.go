package cms

import (
	"github.com/zhoudm1743/Seven/app/cms/routes"
	"github.com/zhoudm1743/Seven/app/cms/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Module("cms"),
	routes.Module,
	service.Module,
)
