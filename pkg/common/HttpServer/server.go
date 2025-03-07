package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/logger"
	"github.com/zhoudm1743/Seven/pkg/common/middleware"
	"net/http"
	"time"
)

type HttpServer struct {
	Gin    *gin.Engine
	Server *http.Server
}

func NewHttpServer(config *config.Config) *HttpServer {
	gin.SetMode(config.Server.Mode)
	engine := gin.Default()
	engine.Use(middleware.Cors())
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 设置静态资源
	engine.StaticFS("/assets", http.Dir("./webroot/assets"))
	engine.StaticFS("/resource", http.Dir("./webroot/resource"))
	engine.StaticFile("/favicon.ico", "./webroot/favicon.ico")
	engine.GET("/", func(c *gin.Context) {
		c.File("./webroot/index.html")
	})
	engine.NoRoute(func(c *gin.Context) {
		c.File("./webroot/index.html")
	})
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	server := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &HttpServer{
		Gin:    engine,
		Server: server,
	}
}
