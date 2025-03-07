package boot

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zhoudm1743/Seven/app/admin"
	"github.com/zhoudm1743/Seven/app/cms"
	web "github.com/zhoudm1743/Seven/pkg/common/HttpServer"
	"github.com/zhoudm1743/Seven/pkg/common/cache"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/database"
	"github.com/zhoudm1743/Seven/pkg/common/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	cache.Module,
	database.Module,
	web.Module,

	admin.Module,
	cms.Module,

	fx.Invoke(run),
)

func run(
	lifecycle fx.Lifecycle,
	server *web.HttpServer,
	db *gorm.DB,
	redis *redis.Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.S().Info("启动Web服务器...", server.Server.Addr)
				server.Server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Web服务器正常退出！")
			// 关闭数据库
			sqlDB, _ := db.DB()
			_ = sqlDB.Close()
			if redis != nil {
				_ = redis.Close()
			}
			return server.Server.Shutdown(ctx)
		},
	})
}
