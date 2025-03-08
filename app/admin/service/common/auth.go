package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	sysSrv "github.com/zhoudm1743/Seven/app/admin/service/system"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"runtime/debug"
	"time"
)

type AuthService interface {
	Tenant() (res []resp.Select, e error)
	Login(c *gin.Context, req *req.LoginReq) (res resp.LoginResp, e error)
	Logout(req *req.LogoutReq) (e error)
}

type authService struct {
	db       *gorm.DB
	cfg      *config.Config
	adminSrv sysSrv.AdminService
}

func (a authService) Tenant() (res []resp.Select, e error) {
	var list []system.Tenant
	if err := a.db.Find(&list).Error; err != nil {
		zap.S().Errorf("TenantList err: err=[%+v]", err)
		e = response.Failed
		return
	}
	for _, v := range list {
		res = append(res, resp.Select{
			Value: v.ID,
			Label: v.Name,
		})
	}
	return
}

func (a authService) Login(c *gin.Context, req *req.LoginReq) (res resp.LoginResp, e error) {
	sysAdmin, err := a.adminSrv.FindByTenantIdAndUsername(req.TenantID, req.Username)
	if err != nil {
		zap.S().Errorf("Login FindByUsername err: err=[%+v]", err)
		e = response.Failed
		return
	}
	if sysAdmin.IsDisable == 1 {
		e = response.LoginDisableError
		return
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password + "zdm")
	if sysAdmin.Password != md5Pwd {
		e = response.LoginAccountError
		return
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			// 自定义类型
			case response.RespType:
				panic(r)
			// 其他类型
			default:
				zap.S().Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
				panic(response.Failed)
			}
		}
	}()
	token := util.ToolsUtil.MakeToken()
	key := fmt.Sprintf("%d_%d", req.TenantID, sysAdmin.ID)

	//非多处登录
	if sysAdmin.IsMultipoint == 0 {
		sysAdminSetKey := a.cfg.Admin.BackstageTokenSet + fmt.Sprintf("%d", sysAdmin.ID)
		ts := util.RedisUtil.SGet(sysAdminSetKey)
		if len(ts) > 0 {
			var keys []string
			for _, t := range ts {
				keys = append(keys, t)
			}
			util.RedisUtil.Del(keys...)
		}
		util.RedisUtil.Del(sysAdminSetKey)
		util.RedisUtil.SSet(sysAdminSetKey, token)
	}

	// 缓存登录信息
	util.RedisUtil.Set(a.cfg.Admin.BackstageTokenKey+token, key, a.cfg.Admin.BackstageTokenExpireTime)
	a.adminSrv.CacheAdminUserByUid(sysAdmin.ID)

	// 更新登录信息
	err = a.db.Model(&sysAdmin).Updates(
		system.Admin{LastLoginIp: c.ClientIP(), LastLoginTime: time.Now().Unix()}).Error
	if err != nil {
		if e = response.CheckErr(err, "Login Updates err"); e != nil {
			return
		}
	}
	// 返回登录信息
	return resp.LoginResp{Token: token}, nil
}

func (a authService) Logout(req *req.LogoutReq) (e error) {
	util.RedisUtil.Del(a.cfg.Admin.BackstageTokenKey + req.Token)
	return
}

func NewAuthService(db *gorm.DB, cfg *config.Config, adminSrv sysSrv.AdminService) AuthService {
	return &authService{db: db, cfg: cfg, adminSrv: adminSrv}
}
