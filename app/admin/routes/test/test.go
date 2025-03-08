package test

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/admin/contracts"
	"github.com/zhoudm1743/Seven/app/admin/service/test"
	"go.uber.org/fx"
)

var Module = fx.Module("testRoutes", fx.Invoke(registerRoutes))

type routeDep struct {
	fx.In
	TestSrv test.TestService
}

func registerRoutes(dep routeDep, r *contracts.AdminRouter) {
	testRoute := r.Group("/test")

	testRoute.GET("/test", dep.test)
}

func (dep routeDep) test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
		"data":    dep.TestSrv.Test(),
	})
}
