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
	SelectMenuByRoleId(roleId uint, auth *req.AuthReq) (mapList []interface{}, e error)
	List(auth *req.AuthReq) (res []interface{}, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemMenuResp, e error)
	Add(addReq req.SystemMenuAddReq, auth *req.AuthReq) (e error)
	Edit(editReq req.SystemMenuEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type menuService struct {
	db            *gorm.DB
	tenantPermSrv TenantPermService
	rolePermSrv   RolePermService
	config        *config.Config
}

func (m menuService) SelectMenuByRoleId(roleId uint, auth *req.AuthReq) (mapList []interface{}, e error) {
	var role system.Role
	if err := m.db.Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, fmt.Errorf("角色不存在")
	}

	var menuIds []uint
	if role.TenantID > 1 {
		tenantMenuIds, err := m.tenantPermSrv.SelectMenuIdsByTenantId(role.TenantID, auth)
		if err != nil {
			return nil, err
		}
		if role.IsAdmin == 1 {
			menuIds = tenantMenuIds
		} else {
			roleMenuIds, err := m.rolePermSrv.SelectMenuIdsByRoleId(roleId, auth)
			if err != nil {
				return nil, err
			}
			menuIds = commonIds(roleMenuIds, tenantMenuIds)
		}
	} else {
		var err error
		if role.IsAdmin == 1 {
			var mIds []uint
			m.db.Where("type in (?)", []int{1, 2}).Order("sort asc, id desc").Pluck("id", &mIds)
			menuIds = mIds
		} else {
			menuIds, err = m.rolePermSrv.SelectMenuIdsByRoleId(roleId, auth)
			if err != nil {
				return nil, err
			}
		}
	}

	var menus []system.Menu
	if err := m.db.Where("id in (?) and is_disable = ? and type in (?)", menuIds, 0, []int{1, 2}).Order("sort asc, id desc").Find(&menus).Error; err != nil || len(menus) == 0 {
		return nil, fmt.Errorf("菜单不存在")
	}

	var respList []resp.SystemMenuResp
	response.Copy(&respList, menus)
	mapList = util.ArrayUtil.ListToTree(util.ConvertUtil.StructsToMaps(respList), "id", "parentId", "children")
	return mapList, nil
}

func (m menuService) List(auth *req.AuthReq) (res []interface{}, e error) {
	var menus []system.Menu
	sql := m.db.Where("is_disable = ?", 0).Order("sort asc, id desc")
	if auth.TenantID > 1 {
		tenantMenuIds, err := m.tenantPermSrv.SelectMenuIdsByTenantId(auth.TenantID, auth)
		if err != nil {
			return nil, err
		}
		sql = sql.Where("id in (?)", tenantMenuIds)
	}
	if err := sql.Find(&menus).Error; err != nil {
		return nil, err
	}
	var respList []resp.SystemMenuResp
	response.Copy(&respList, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(respList), "id", "parentId", "children"), nil
}

func (m menuService) Detail(id uint, auth *req.AuthReq) (res resp.SystemMenuResp, e error) {
	var menu system.Menu
	if err := m.db.Where("id = ? and is_disable = ?", id, 0).First(&menu).Error; err != nil {
		return res, fmt.Errorf("菜单不存在")
	}
	response.Copy(&res, menu)
	if menu.Type == 2 {
		var buttons []system.Menu
		m.db.Model(&system.Menu{}).Where("parent_id = ? and type = ?", id, 3).Find(&buttons)
		var buttonList []resp.SystemMenuButton
		if len(buttons) > 0 {
			response.Copy(&buttonList, buttons)
		}
		res.Button = buttonList
	}
	return res, nil
}

func (m menuService) Add(addReq req.SystemMenuAddReq, auth *req.AuthReq) (e error) {
	var count int64
	m.db.Model(&system.Menu{}).Where("key = ? or auth = ?", addReq.Key, addReq.Auth).Count(&count)
	if count > 0 {
		return fmt.Errorf("菜单标识或权限已存在")
	}
	var menu system.Menu
	response.Copy(&menu, addReq)
	if err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}
		if menu.Type == 2 {
			for _, button := range addReq.Button {
				b := system.Menu{
					ParentID: menu.ID,
					Label:    button.Label,
					Type:     3,
					Auth:     button.Auth,
					Sort:     button.Sort,
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

func (m menuService) Edit(editReq req.SystemMenuEditReq, auth *req.AuthReq) (e error) {
	var count int64
	m.db.Model(&system.Menu{}).Where("(key = ? and id <> ?) or (auth = ? and id <> ?)", editReq.Key, editReq.ID, editReq.Auth, editReq.ID).Count(&count)
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
		if menu.Type == 2 {
			var buttons []system.Menu
			tx.Model(&system.Menu{}).Where("parent_id = ? and type = ?", menu.ID, 3).Find(&buttons)
			for _, button := range editReq.Button {
				var b system.Menu
				if err := tx.Where("parent_id = ? and type = ? and auth = ?", menu.ID, 3, button.Auth).First(&b).Error; err == nil {
					response.Copy(&b, button)
					if err := tx.Save(&b).Error; err != nil {
						return err
					}
				} else {
					b := system.Menu{
						ParentID: menu.ID,
						Label:    button.Label,
						Type:     3,
						Auth:     button.Auth,
						Sort:     button.Sort,
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
		tx.Model(&system.Menu{}).Where("parent_id = ?", id).Count(&count)
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
