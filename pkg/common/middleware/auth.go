package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/service/system"
	sysModel "github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/database"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

var (
	cfg           = config.NewConfig()
	permSrv       = system.NewRolePermService(database.GetDB(), cfg)
	roleSrv       = system.NewRoleService(database.GetDB(), permSrv, cfg)
	adminSrv      = system.NewAdminService(database.GetDB(), permSrv, roleSrv, cfg)
	tenantPermSrv = system.NewTenantPermService(database.GetDB(), cfg)
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/admin/", "", 1), "/", ":")
		color.Blueln("auths: ", auths)
		// 免登录接口
		color.Blueln("cfg.Admin.NotLoginUri: ", cfg.Admin.NotLoginUri)
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

		if !util.RedisUtil.HExists(cfg.Admin.BackstageManageKey, uidStr) {
			err := adminSrv.CacheAdminUserByUid(uid)
			if err != nil {
				zap.S().Errorf("TokenAuth CacheAdminUserByUid err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 校验用户被删除
		var mapping sysModel.Admin
		err := util.ToolsUtil.JsonToObj(util.RedisUtil.HGet(cfg.Admin.BackstageManageKey, uidStr), &mapping)
		if err != nil {
			zap.S().Errorf("TokenAuth Unmarshal err: err=[%+v]", err)
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		}

		// 校验用户被禁用
		if mapping.IsDisable == 1 {
			response.Fail(c, response.LoginDisableError)
			c.Abort()
			return
		}

		// 令牌剩余30分钟自动续签
		if util.RedisUtil.TTL(tokenKey) < 1800 {
			util.RedisUtil.Expire(tokenKey, 7200)
		}

		isAdmin := false
		if mapping.Role != nil && mapping.Role.IsAdmin == 1 {
			isAdmin = true
		}

		auth := req.AuthReq{
			Id:            uid,
			TenantID:      cid,
			IsAdmin:       isAdmin,
			IsSuperTenant: mapping.TenantId == 1,
			DeptId:        mapping.DeptId,
			PostId:        mapping.PostId,
			RoleId:        mapping.RoleId,
			Username:      mapping.Username,
			Nickname:      mapping.Nickname,
		}

		// 检查租户是否停用或过期
		//now := time.Now().Unix()
		//if !auth.IsSuperTenant && mapping.Tenant.ExpireAt > 0 && (mapping.Tenant.IsDisable == 1 || mapping.Tenant.ExpireAt < now) {
		//	response.Fail(c, response.TenantDisableOrExpired)
		//	c.Abort()
		//	return
		//}
		// session 存储用户信息
		c.Set("auth", auth)

		// 免权限验证接口
		if util.ToolsUtil.Contains(cfg.Admin.NotAuthUri, auths) || (auth.IsAdmin && auth.IsSuperTenant) {
			c.Next()
			return
		}
		// 权限检查
		// 校验租户权限或角色权限
		var checkRolePermissions func(roleId uint) bool
		checkRolePermissions = func(roleId uint) bool {
			if util.RedisUtil.HExists(cfg.Admin.BackstageRolesKey, fmt.Sprintf("%d", roleId)) {
				i := cast.ToUint(roleId)
				err = permSrv.CacheRoleMenusByRoleId(i)
				if err != nil {
					zap.S().Errorf("TokenAuth CacheRoleMenusByRoleId err: err=[%+v]", err)
					response.Fail(c, response.SystemError)
					c.Abort()
					return false
				}

				menus := util.RedisUtil.HGet(cfg.Admin.BackstageRolesKey, fmt.Sprintf("%d_%d", cid, roleId))
				return menus != "" && util.ToolsUtil.Contains(strings.Split(menus, ","), auths)
			}
			return false
		}

		roleId := mapping.RoleId
		var hasPermission bool
		if auth.IsSuperTenant {
			hasPermission = checkRolePermissions(roleId)
		} else {
			tenantId := strconv.Itoa(int(mapping.TenantId))
			if util.RedisUtil.HExists(cfg.Admin.BackstageTenantsKey, tenantId) {
				err := tenantPermSrv.CacheTenantMenusByTenantId(mapping.TenantId)
				if err != nil {
					zap.S().Errorf("TokenAuth CacheTenantMenusByTenantId err: err=[%+v]", err)
					response.Fail(c, response.SystemError)
					c.Abort()
					return
				}

				if auth.IsAdmin {
					menus := util.RedisUtil.HGet(cfg.Admin.BackstageTenantsKey, tenantId)
					hasPermission = menus != "" && util.ToolsUtil.Contains(strings.Split(menus, ","), auths)
				} else {
					hasPermission = checkRolePermissions(roleId)
				}
			}
		}

		if !hasPermission {
			response.Fail(c, response.NoPermission)
			c.Abort()
			return
		}

		c.Next()
	}
}
