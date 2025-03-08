package service

import (
	"github.com/zhoudm1743/Seven/app/admin/service/common"
	"github.com/zhoudm1743/Seven/app/admin/service/system"
	"github.com/zhoudm1743/Seven/app/admin/service/test"
	"go.uber.org/fx"
)

var Module = fx.Module("adminServices",
	fx.Provide(test.NewTestService),
	// system services here
	fx.Provide(
		system.NewTenantPermService,
		system.NewRolePermService,
		system.NewMenuService,
		system.NewTenantService,
		system.NewRoleService,
		system.NewAdminService,
		system.NewDeptService,
		system.NewPostService,
	),
	// common services here
	fx.Provide(
		common.NewAuthService,
	),
)
