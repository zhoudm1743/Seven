package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path"
)

func NewConfig() *Config {
	defaultcfg := &Config{
		Server: Server{
			Host:         "0.0.0.0",
			Port:         8080,
			Mode:         "debug",
			Secret:       "zdm",
			PublicUrl:    "public",
			PublicPrefix: "public/",
		},
		Log: Log{
			Level:      "debug",
			Filename:   "logs/app.log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
		},
		Redis: Redis{
			Host:          "localhost",
			Port:          6379,
			Password:      "",
			DB:            0,
			RedisPoolSize: 100,
			RedisPrefix:   "zdm_",
		},
		Database: Database{
			Driver:      "mysql",
			Host:        "localhost",
			Port:        3306,
			Username:    "root",
			Password:    "root",
			Database:    "test",
			MaxIdle:     10,
			MaxOpen:     100,
			Prefix:      "test_",
			AutoMigrate: true,
		},
		Admin: Admin{
			BackstageTenantsKey: "backstage:tenants",
			CommonUri:           []string{},
			BackstageRolesKey:   "backstage:roles",
		},
	}
	conf := &Config{}
	viper.SetConfigType(configType)
	src, _ := os.Getwd()
	filePath := path.Join(src, configFile)
	color.Greenln("配置文件路径", filePath)
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		color.Redln("配置信息初始化失败", err)
		return defaultcfg
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		viper.Unmarshal(conf)
		zap.S().Info("配置信息发生变动", conf)
	})
	viper.WatchConfig()

	if err := viper.Unmarshal(conf); err != nil {
		color.Redln("配置信息初始化失败", err)
		return defaultcfg
	}

	return conf
}
