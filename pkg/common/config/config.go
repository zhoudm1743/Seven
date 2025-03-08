package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path"
	"reflect"
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
			BackstageTenantsKey:      "backstage:tenants",
			BackstageRolesKey:        "backstage:roles",
			BackstageTokenKey:        "backstage:token:",
			BackstageTokenSet:        "backstage:token:set:",
			BackstageManageKey:       "backstage:manage",
			BackstageTokenExpireTime: 86400,
			CommonUri:                []string{},
			NotLoginUri: []string{
				"common:auth:login",
				"common:auth:tenant",
			},
			NotAuthUri: []string{
				"common:auth:logout",
				"system:menu:menus",   // 系统菜单
				"system:menu:route",   // 菜单路由
				"system:admin:upInfo", // 管理员更新
				"system:admin:self",   // 管理员信息
				"system:role:all",     // 所有角色
			},
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

	applyDefaults(conf, defaultcfg)

	return conf
}

func applyDefaults(conf, defaultCfg interface{}) {
	confElem := reflect.ValueOf(conf).Elem()
	defaultElem := reflect.ValueOf(defaultCfg).Elem()

	for i := 0; i < confElem.NumField(); i++ {
		confField := confElem.Field(i)
		defaultField := defaultElem.Field(i)

		// 递归处理嵌套结构体
		if confField.Kind() == reflect.Struct {
			applyDefaults(confField.Addr().Interface(), defaultField.Addr().Interface())
			continue
		}

		// 处理切片类型（保持原有逻辑只处理nil情况）
		if confField.Kind() == reflect.Slice && confField.IsNil() {
			confField.Set(defaultField)
			continue
		}

		// 处理零值替换
		if isZeroValue(confField) {
			confField.Set(defaultField)
		}
	}
}

// 判断是否零值（增强版）
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan:
		return v.IsNil()
	default:
		return false
	}
}
