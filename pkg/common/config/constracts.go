package config

const (
	configFile = "config/config.yml"
	configType = "yml"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Redis    Redis    `yaml:"redis"`
	Database Database `yaml:"database"`
	Admin    Admin    `yaml:"admin"`
}

type Server struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Mode         string `yaml:"mode"`
	Secret       string `yaml:"secret"`
	PublicUrl    string `yaml:"public_url"`
	PublicPrefix string `yaml:"public_prefix"`
}

type Log struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxage"`
}

type Redis struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Password      string `yaml:"password"`
	DB            int    `yaml:"db"`
	RedisPoolSize int    `yaml:"redis_pool_size"`
	RedisPrefix   string `yaml:"redis_prefix"`
}

type Database struct {
	Driver      string `yaml:"driver"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	MaxIdle     int    `yaml:"maxidle"`
	MaxOpen     int    `yaml:"maxopen"`
	Prefix      string `yaml:"prefix"`
	AutoMigrate bool   `yaml:"autoMigrate"`
}

type Admin struct {
	BackstageTenantsKey      string `yaml:"backstage_tenants_key"`
	BackstageRolesKey        string `yaml:"backstage_roles_key"`
	BackstageTokenKey        string `yaml:"backstage_token_key"`
	BackstageTokenSet        string `yaml:"backstage_token_set"`
	BackstageTokenExpireTime int    `yaml:"backstage_token_expire_time"`
	BackstageManageKey       string `yaml:"backstage_manage_key"`
	NotLoginUri              []string
	CommonUri                []string
	NotAuthUri               []string
}
