package system

import (
	"fmt"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"
)

type MenuService interface {
	SelectMenuByRoleId(auth *req.AuthReq) (mapList interface{}, e error)
	List(auth *req.AuthReq) (res interface{}, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemMenuResp, e error)
	Add(addReq *req.SystemMenuAddReq, auth *req.AuthReq) (e error)
	Edit(editReq *req.SystemMenuEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type menuService struct {
	db            *gorm.DB
	tenantPermSrv TenantPermService
	rolePermSrv   RolePermService
	config        *config.Config
}

func (m menuService) SelectMenuByRoleId(auth *req.AuthReq) (mapList interface{}, e error) {
	var role system.Role
	if err := m.db.Where("id = ?", auth.RoleId).First(&role).Error; err != nil {
		return nil, fmt.Errorf("角色不存在")
	}

	var menuIds []uint
	if role.TenantId > 1 {
		tenantMenuIds, err := m.tenantPermSrv.SelectMenuIdsByTenantId(role.TenantId, auth)
		if err != nil {
			return nil, err
		}
		if role.IsAdmin == 1 {
			menuIds = tenantMenuIds
		} else {
			roleMenuIds, err := m.rolePermSrv.SelectMenuIdsByRoleId(role.ID, auth)
			if err != nil {
				return nil, err
			}
			menuIds = commonIds(roleMenuIds, tenantMenuIds)
		}
	} else {
		var err error
		if role.IsAdmin == 1 {
			var mIds []uint
			m.db.Model(&system.Menu{}).Where("menuType in (?)", []string{"dir", "page"}).Order("sort asc, id desc").Pluck("id", &mIds)
			menuIds = mIds
		} else {
			menuIds, err = m.rolePermSrv.SelectMenuIdsByRoleId(role.ID, auth)
			if err != nil {
				return nil, err
			}
		}
	}

	var menus []system.Menu
	if err := m.db.Where("id in (?)", menuIds).Order("sort asc, id desc").Find(&menus).Error; err != nil || len(menus) == 0 {
		return nil, fmt.Errorf("菜单不存在")
	}

	var respList []resp.SystemMenuResp
	response.Copy(&respList, menus)

	return util.ArrayUtil.ListToTree(util.ConvertUtil.StructsToMaps(respList), "id", "pid", "children"), nil
	//return respList, nil
}

func (m menuService) List(auth *req.AuthReq) (res interface{}, e error) {
	var menus []system.Menu
	sql := m.db.Order("id desc")
	if auth.TenantID > 1 {
		tenantMenuIds, err := m.tenantPermSrv.SelectMenuIdsByTenantId(auth.TenantID, auth)
		if err != nil {
			return nil, err
		}
		sql = sql.Where("id in (?)", tenantMenuIds)
	}
	if err := sql.Order("sort asc, id desc").Find(&menus).Error; err != nil {
		return nil, err
	}
	var respList []resp.SystemMenuResp
	response.Copy(&respList, menus)
	//return util.ArrayUtil.ListToTree(
	//	util.ConvertUtil.StructsToMaps(respList), "id", "pid", "children"), nil
	return respList, nil
}

func (m menuService) Detail(id uint, auth *req.AuthReq) (res resp.SystemMenuResp, e error) {
	var menu system.Menu
	if err := m.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return res, fmt.Errorf("菜单不存在")
	}
	response.Copy(&res, menu)
	if menu.MenuType == "page" {
		var buttons []system.Menu
		m.db.Model(&system.Menu{}).Where("pid = ? and menuType = ?", id, "action").Order("sort asc, id desc").Find(&buttons)
		var buttonList []resp.SystemMenuButton
		if len(buttons) > 0 {
			response.Copy(&buttonList, buttons)
		}
		res.Button = buttonList
	}
	return res, nil
}

func (m menuService) Add(addReq *req.SystemMenuAddReq, auth *req.AuthReq) (e error) {
	var count int64
	m.db.Model(&system.Menu{}).Where("name = ?", addReq.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("菜单标识或权限已存在")
	}
	var menu system.Menu
	response.Copy(&menu, addReq)
	if err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}
		if menu.MenuType == "page" {
			for _, button := range addReq.Button {
				b := system.Menu{
					Pid:      menu.ID,
					Title:    button.Title,
					MenuType: "action",
					Name:     button.Name,
					Hide:     true,
				}
				if err := tx.Create(&b).Error; err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	util.RedisUtil.Del(m.config.Admin.BackstageRolesKey)
	return nil
}

func (m menuService) Edit(editReq *req.SystemMenuEditReq, auth *req.AuthReq) (e error) {
	var count int64
	m.db.Model(&system.Menu{}).Where("name = ? and id != ?", editReq.Name, editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("菜单标识或权限已存在")
	}
	var menu system.Menu
	if err := m.db.Where("id = ?", editReq.ID).First(&menu).Error; err != nil {
		return fmt.Errorf("菜单不存在")
	}
	response.Copy(&menu, editReq)
	if err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&menu).Error; err != nil {
			return err
		}
		if menu.MenuType == "page" {
			var buttons []system.Menu
			tx.Model(&system.Menu{}).Where("pid = ? and menuType = ?", menu.ID, "action").Find(&buttons)
			for _, button := range editReq.Button {
				var b system.Menu
				if err := tx.Where("pid = ? and menuType = ? and name = ?", menu.ID, "action", button.Name).First(&b).Error; err == nil {
					response.Copy(&b, button)
					if err := tx.Save(&b).Error; err != nil {
						return err
					}
				} else {
					b := system.Menu{
						Pid:      menu.ID,
						Hide:     true,
						Title:    button.Title,
						MenuType: "action",
						Name:     button.Name,
					}
					if err := tx.Create(&b).Error; err != nil {
						return err
					}
				}
			}
			for _, button := range buttons {
				if button.ID == 0 {
					continue
				}
				var b system.Menu
				if err := tx.Where("id = ?", button.ID).First(&b).Error; err != nil {
					return err
				}
				if err := tx.Delete(&b).Error; err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	util.RedisUtil.Del(m.config.Admin.BackstageRolesKey)
	return nil
}

func (m menuService) Del(id uint, auth *req.AuthReq) (e error) {
	var menu system.Menu
	if err := m.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return fmt.Errorf("菜单不存在")
	}
	if err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&menu).Error; err != nil {
			return err
		}
		var count int64
		tx.Model(&system.Menu{}).Where("pid = ?", id).Count(&count)
		if count > 0 {
			return fmt.Errorf("请先删除子菜单")
		}
		err := m.tenantPermSrv.BatchDeleteByMenuId(id, tx, auth)
		if err != nil {
			return err
		}
		err = m.rolePermSrv.BatchDeleteByMenuId(id, tx, auth)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func NewMenuService(db *gorm.DB, tenantPermSrv TenantPermService, rolePermSrv RolePermService, config *config.Config) MenuService {
	return &menuService{
		db:            db,
		tenantPermSrv: tenantPermSrv,
		rolePermSrv:   rolePermSrv,
		config:        config,
	}
}

func commonIds(ids []uint, ids2 []uint) []uint {
	var res []uint
	for _, id := range ids {
		for _, id2 := range ids2 {
			if id == id2 {
				res = append(res, id)
				break
			}
		}
	}
	return res
}
