package service

import (
	"github.com/zhoudm1743/Seven/app/cms/service/test"
	"go.uber.org/fx"
)

var Module = fx.Module("cmsServices",
	fx.Provide(test.NewTestService),
)
