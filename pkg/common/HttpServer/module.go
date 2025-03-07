package web

import "go.uber.org/fx"

var Module = fx.Provide(
	NewHttpServer,
)
