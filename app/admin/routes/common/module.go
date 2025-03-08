package common

import (
	"go.uber.org/fx"
)

var Module = fx.Module("commonRoutes",
	fx.Invoke(authRoutes),
)
