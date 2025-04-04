package system

import (
	"fmt"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type TenantPermService interface {
	SelectMenuIdsByTenantId(tenantId uint, auth *req.AuthReq) (menuIds []uint, e error)
	CacheTenantMenusByTenantId(tenantId uint) (e error)
	BatchSaveByMenuIds(tenantId uint, db *gorm.DB, menuIds string, auth *req.AuthReq) (e error)
	BatchDeleteByTenantId(tenantId uint, db *gorm.DB, auth *req.AuthReq) (e error)
	BatchDeleteByMenuId(menuId uint, db *gorm.DB, auth *req.AuthReq) (e error)
}

type tenantPermService struct {
	db  *gorm.DB
	cfg *config.Config
}

func (t tenantPermService) SelectMenuIdsByTenantId(tenantId uint, auth *req.AuthReq) (menuIds []uint, e error) {
	if tenantId == 0 {
		return []uint{}, nil
	}
	e = t.db.Where("tenant_id =?", tenantId).Pluck("menu_id", &menuIds).Error
	return
}

func (t tenantPermService) CacheTenantMenusByTenantId(tenantId uint) (e error) {
	if tenantId == 0 {
		return fmt.Errorf("tenantId is empty")
	}
	var menuIds []uint
	e = t.db.Where("tenant_id = ?", tenantId).Pluck("menu_id", &menuIds).Error
	if e != nil || len(menuIds) == 0 {
		return
	}
	var menus []system.Menu
	e = t.db.Where("id in (?) and menuType in (?)", menuIds, []string{"dir", "page"}).Order("id desc").Find(&menus).Error
	if e != nil {
		return
	}
	if len(menus) == 0 {
		return fmt.Errorf("tenantId %d not found", tenantId)
	}
	var menuArray []string
	for _, menu := range menus {
		menuArray = append(menuArray, strings.Trim(menu.Name, ""))
	}
	// 其他权限
	if len(t.cfg.Admin.CommonUri) > 0 {
		menuArray = append(menuArray, t.cfg.Admin.CommonUri...)
	}

	util.RedisUtil.HSet(t.cfg.Admin.BackstageTenantsKey, strconv.FormatUint(uint64(tenantId), 10), strings.Join(menuArray, ","), 0)
	return
}

func (t tenantPermService) BatchSaveByMenuIds(tenantId uint, db *gorm.DB, menuIds string, auth *req.AuthReq) (e error) {
	if tenantId == 0 {
		return fmt.Errorf("tenantId is empty")
	}
	if menuIds == "" {
		return fmt.Errorf("menuIds is empty")
	}
	if db == nil {
		db = t.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system.TenantPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, system.TenantPerm{ID: util.ToolsUtil.MakeUuid(), TenantID: tenantId, MenuID: uint(menuId)})
		}
		err := tx.Create(&perms).Error
		if err != nil {
			return err
		}
		return nil
	})
	e = response.CheckErr(err, "BatchSaveByMenuIds Transaction err")
	return
}

func (t tenantPermService) BatchDeleteByTenantId(tenantId uint, db *gorm.DB, auth *req.AuthReq) (e error) {
	if db == nil {
		db = t.db
	}
	return db.Where("tenant_id =?", tenantId).Delete(&system.TenantPerm{}).Error
}

func (t tenantPermService) BatchDeleteByMenuId(menuId uint, db *gorm.DB, auth *req.AuthReq) (e error) {
	if db == nil {
		db = t.db
	}
	return db.Where("menu_id =?", menuId).Delete(&system.TenantPerm{}).Error
}

func NewTenantPermService(db *gorm.DB, cfg *config.Config) TenantPermService {
	return &tenantPermService{
		db:  db,
		cfg: cfg,
	}
}
