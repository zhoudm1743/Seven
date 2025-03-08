package system

import "go.uber.org/fx"

var Module = fx.Module("systemRoutes",
	fx.Invoke(
		tenantRoutes,
		deptRoutes,
		postRoutes,
		menuRoutes,
		roleRoutes,
		adminRoutes,
	),
)
