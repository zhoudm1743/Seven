package test

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/cms/contracts"
	"github.com/zhoudm1743/Seven/app/cms/service/test"
	web "github.com/zhoudm1743/Seven/pkg/common/HttpServer"
	"go.uber.org/fx"
)

var Module = fx.Module("testRoutes", fx.Invoke(registerRoutes))

type routeDep struct {
	fx.In
	Http    *web.HttpServer
	TestSrv test.TestService
}

func registerRoutes(dep routeDep, r *contracts.CmsRouter) {
	testRoute := r.Group("/test")

	testRoute.GET("/test", dep.test)
}

func (dep routeDep) test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
		"data":    dep.TestSrv.Test(),
	})
}
