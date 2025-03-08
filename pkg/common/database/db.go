package database

import (
	"fmt"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func NewDb(config *config.Config) (dbInstance *gorm.DB, err error) {
	logg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logg,
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.Prefix,
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.Database.MaxIdle)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpen)
	if config.Database.AutoMigrate {
		err := migrate(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

// GetDb 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
