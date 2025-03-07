package util

import (
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"go.uber.org/zap"
	"net/url"
	"path"
	"strings"
)

var (
	UrlUtil = urlUtil{config: config.NewConfig()}
)

// urlUtil 文件路径处理工具
type urlUtil struct {
	config *config.Config
}

// ToAbsoluteUrl 转绝对路径
func (uu urlUtil) ToAbsoluteUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(uu.config.Server.PublicUrl)
	if err != nil {
		zap.S().Errorf("ToAbsoluteUrl Parse err: err=[%+v]", err)
		return u
	}
	if strings.HasPrefix(u, "/api/static/") {
		up.Path = path.Join(up.Path, u)
		return up.String()
	}
	engine := "local"
	if engine == "local" {
		up.Path = path.Join(up.Path, uu.config.Server.PublicPrefix, u)
		return up.String()
	}
	// TODO: 其他engine
	return u
}

func (uu urlUtil) ToRelativeUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(u)
	if err != nil {
		zap.S().Errorf("ToRelativeUrl Parse err: err=[%+v]", err)
		return u
	}
	engine := "local"
	if engine == "local" {
		lu := up.String()
		return strings.Replace(
			strings.Replace(lu, uu.config.Server.PublicUrl, "", 1),
			uu.config.Server.PublicPrefix, "", 1)
	}
	// TODO: 其他engine
	return u
}
