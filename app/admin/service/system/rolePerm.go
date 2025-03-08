package system

import (
	"fmt"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type RolePermService interface {
	SelectMenuIdsByRoleId(roleId uint, auth *req.AuthReq) (menuIds []uint, e error)
	CacheRoleMenusByRoleId(roleId uint) (e error)
	BatchSaveByMenuIds(roleId uint, db *gorm.DB, menuIds string, auth *req.AuthReq) (e error)
	BatchDeleteByRoleId(roleId uint, db *gorm.DB, auth *req.AuthReq) (e error)
	BatchDeleteByMenuId(menuId uint, db *gorm.DB, auth *req.AuthReq) (e error)
}

type rolePermService struct {
	db  *gorm.DB
	cfg *config.Config
}

func (r rolePermService) SelectMenuIdsByRoleId(roleId uint, auth *req.AuthReq) (menuIds []uint, e error) {
	if roleId == 0 {
		return []uint{}, nil
	}
	e = r.db.Where("role_id =?", roleId).Pluck("menu_id", &menuIds).Error
	if e == nil {
		return []uint{}, nil
	}
	return menuIds, nil
}

func (r rolePermService) CacheRoleMenusByRoleId(roleId uint) (e error) {
	if roleId == 0 {
		return fmt.Errorf("roleId is empty")
	}
	var role system.Role
	r.db.Where("id =?", roleId).First(&role)
	if role.ID == 0 {
		return fmt.Errorf("roleId %d not found", roleId)
	}
	var menuIds []uint
	r.db.Where("role_id = ?", roleId).Pluck("menu_id", &menuIds)
	var menus []system.Menu
	r.db.Where("id in (?) and type in (?) and is_disable = ?", menuIds, []int{1, 2}, 0).Order("sort asc, id desc").Find(&menus)
	if len(menus) == 0 {
		return fmt.Errorf("roleId %d not found", roleId)
	}
	var menuArray []string
	for _, menu := range menus {
		menuArray = append(menuArray, strings.Trim(menu.Auth, ""))
	}
	// 其他权限
	if len(r.cfg.Admin.CommonUri) > 0 {
		menuArray = append(menuArray, r.cfg.Admin.CommonUri...)
	}
	key := fmt.Sprintf("%d:%d", role.TenantId, roleId)
	util.RedisUtil.HSet(r.cfg.Admin.BackstageRolesKey, key, strings.Join(menuArray, ","), 0)
	return
}

func (r rolePermService) BatchSaveByMenuIds(roleId uint, db *gorm.DB, menuIds string, auth *req.AuthReq) (e error) {
	if roleId == 0 {
		return fmt.Errorf("roleId is empty")
	}
	if menuIds == "" {
		return fmt.Errorf("menuIds is empty")
	}
	if db == nil {
		db = r.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system.RolePerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, system.RolePerm{ID: util.ToolsUtil.MakeUuid(), RoleID: roleId, MenuID: uint(menuId)})
		}
		err := tx.Create(&perms).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r rolePermService) BatchDeleteByRoleId(roleId uint, db *gorm.DB, auth *req.AuthReq) (e error) {
	if db == nil {
		db = r.db
	}
	return db.Where("role_id = ?", roleId).Delete(system.RolePerm{}).Error
}

func (r rolePermService) BatchDeleteByMenuId(menuId uint, db *gorm.DB, auth *req.AuthReq) (e error) {
	if db == nil {
		db = r.db
	}
	return db.Where("menu_id = ?", menuId).Delete(system.RolePerm{}).Error
}

func NewRolePermService(db *gorm.DB, cfg *config.Config) RolePermService {
	return &rolePermService{
		db:  db,
		cfg: cfg,
	}
}
