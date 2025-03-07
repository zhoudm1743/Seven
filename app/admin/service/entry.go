package service

import (
	"github.com/zhoudm1743/Seven/app/admin/service/test"
	"go.uber.org/fx"
)

var Module = fx.Module("adminServices",
	fx.Provide(test.NewTestService),
)
