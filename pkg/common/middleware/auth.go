package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/spf13/cast"

	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"strings"
)

var (
	cfg = config.NewConfig()
	//permSrv       = system.NewRolePermService(core.GetDB())
	//roleSrv       = system.NewRoleService(core.GetDB(), permSrv)
	//adminSrv      = system.NewAdminService(core.GetDB(), permSrv, roleSrv)
	//tenantPermSrv = system.NewTenantPermService(core.GetDB())
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/admin/", "", 1), "/", ":")
		// 免登录接口
		if util.ToolsUtil.Contains(cfg.Admin.NotLoginUri, auths) {
			c.Next()
			return
		}
		// 登录检查
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			response.Fail(c, response.TokenEmpty)
			c.Abort()
			return
		}
		// 验证token
		tokenKey := cfg.Admin.BackstageTokenKey + token
		existCnt := util.RedisUtil.Exists(tokenKey)

		if existCnt < 0 {
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		} else if existCnt == 0 {
			response.Fail(c, response.TokenInvalid)
			c.Abort()
			return
		}
		uidStr := util.RedisUtil.Get(tokenKey)
		var uid, cid uint
		if uidStr != "" {
			spilt := strings.Split(uidStr, "_")
			uid = cast.ToUint(spilt[1])
			cid = cast.ToUint(spilt[0])
		}
		color.Redln("uid:", uid, "cid:", cid)
	}
}
