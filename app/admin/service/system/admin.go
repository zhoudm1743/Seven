package system

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AdminService interface {
	FindByUsername(username string) (admin system.Admin, err error)
	FindByTenantIdAndUsername(tenantId uint, username string) (admin system.Admin, err error)
	FindByTenantIdAndId(tenantId uint, id uint) (admin system.Admin, err error)
	Self(adminId uint, auth *req.AuthReq) (res resp.SystemAdminSelfResp, e error)
	List(page req.PageReq, listReq req.SystemAdminListReq, auth *req.AuthReq) (res response.PageResp, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemAdminResp, e error)
	Add(addReq req.SystemAdminAddReq, auth *req.AuthReq) (e error)
	Edit(c *gin.Context, editReq req.SystemAdminEditReq, auth *req.AuthReq) (e error)
	Update(c *gin.Context, updateReq req.SystemAdminUpdateReq, adminId uint, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
	Disable(id uint, auth *req.AuthReq) (e error)
	CacheAdminUserByUid(id uint) (err error)
}

type adminService struct {
	db          *gorm.DB
	rolePermSrv RolePermService
	roleSrv     RoleService
	cfg         *config.Config
}

func (a adminService) FindByUsername(username string) (admin system.Admin, err error) {
	err = a.db.Where("username = ?", username).First(&admin).Error
	return
}

func (a adminService) FindByTenantIdAndUsername(tenantId uint, username string) (admin system.Admin, err error) {
	err = a.db.Where("tenant_id = ? AND username = ?", tenantId, username).First(&admin).Error
	return
}

func (a adminService) FindByTenantIdAndId(tenantId uint, id uint) (admin system.Admin, err error) {
	err = a.db.Where("tenant_id = ? AND id = ?", tenantId, id).First(&admin).Error
	return
}

func (a adminService) Self(adminId uint, auth *req.AuthReq) (res resp.SystemAdminSelfResp, e error) {
	var sysAdmin system.Admin
	err := a.db.Where("id = ?", adminId).Limit(1).First(&sysAdmin).Error
	if e = response.CheckErr(err, "Self First err"); e != nil {
		return
	}
	// 角色权限
	var auths []string
	if !auth.IsAdmin {
		roleId := sysAdmin.RoleId
		var menuIds []uint
		if menuIds, e = a.rolePermSrv.SelectMenuIdsByRoleId(roleId, auth); e != nil {
			return
		}

		if len(menuIds) > 0 {
			var menus []system.Menu
			err := a.db.Where(
				"id in ? AND type in ?", menuIds, 0, []int{1, 2}).Order(
				"menu_sort asc, id desc").Find(&menus).Error
			if e = response.CheckErr(err, "Self SystemAuthMenu Find err"); e != nil {
				return
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Auth, " "))
				}
			}
		}
		if auth.IsAdmin {
			var tentMenus []uint
			var menus []system.Menu
			a.db.Model(&system.TenantPerm{}).Where("tenant_id = ?", auth.TenantID).Pluck("menu_id", &tentMenus)
			err := a.db.Where(
				"id in ? AND type in ?", tentMenus, 0, []int{1, 2}).Order(
				"menu_sort asc, id desc").Find(&menus).Error
			if e = response.CheckErr(err, "Self SystemAuthMenu Find err"); e != nil {
				return
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Auth, " "))
				}
			}
		}
		if len(auths) > 0 {
			auths = append(auths, "")
		}
	} else {
		auths = append(auths, "*")
	}
	var admin resp.SystemAdminSelfOneResp
	response.Copy(&admin, sysAdmin)
	admin.Dept = strconv.FormatUint(uint64(sysAdmin.DeptId), 10)
	admin.Avatar = util.UrlUtil.ToAbsoluteUrl(sysAdmin.Avatar)
	admin.SoftSuper = auth.IsAdmin
	admin.SuperTenant = auth.IsSuperTenant
	return resp.SystemAdminSelfResp{User: admin, Permissions: auths}, nil
}

func (a adminService) List(page req.PageReq, listReq req.SystemAdminListReq, auth *req.AuthReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	adminTbName := util.DbUtil.DBTableName(a.db, &system.Admin{})
	roleTbName := util.DbUtil.DBTableName(a.db, &system.Role{})
	deptTbName := util.DbUtil.DBTableName(a.db, &system.Dept{})
	tenantTbName := util.DbUtil.DBTableName(a.db, &system.Tenant{})
	sql := a.db
	adminModel := sql.Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.role::bigint = %s.id", roleTbName, roleTbName)).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).
		Joins(
			fmt.Sprintf("LEFT JOIN %s ON admin.tenant_id = %s.id", tenantTbName, tenantTbName)).
		Select(fmt.Sprintf("admin.*, %s.name as dept, %s.name as role, %s.name as tenant", deptTbName, roleTbName, tenantTbName))
	// 条件
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.Role != nil && *listReq.Role > 0 {
		adminModel = adminModel.Where("role = ?", cast.ToString(*listReq.Role))
	}
	if listReq.TenantId != nil && *listReq.TenantId > 0 {
		adminModel = adminModel.Where("admin.tenant_id = ?", cast.ToString(*listReq.TenantId))
	}
	if !auth.IsSuperTenant {
		adminModel = adminModel.Where("admin.tenant_id = ?", auth.TenantID)
	}
	// 总数
	var count int64
	err := adminModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var adminResp []resp.SystemAdminResp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    adminResp,
	}, nil
}

func (a adminService) Detail(id uint, auth *req.AuthReq) (res resp.SystemAdminResp, e error) {
	var sysAdmin system.Admin
	sql := query.AuthQuery(a.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&sysAdmin).Error
	if err != nil {
		return res, response.CheckErr(err, "Detail First err")
	}
	response.Copy(&res, sysAdmin)
	res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
	if res.Dept == "" {
		res.Dept = strconv.FormatUint(uint64(res.DeptId), 10)
	}
	return
}

func (a adminService) Add(addReq req.SystemAdminAddReq, auth *req.AuthReq) (e error) {
	var sysAdmin system.Admin
	// 检查username
	sql := query.AuthQuery(a.db, auth)
	r := sql.Where("username = ?", addReq.Username).Limit(1).Find(&sysAdmin)
	err := r.Error
	if e = response.CheckErr(err, "Add Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("账号已存在换一个吧！")
	}
	// 检查nickname
	r = sql.Where("nickname = ?", addReq.Nickname).Limit(1).Find(&sysAdmin)
	err = r.Error
	if e = response.CheckErr(err, "Add Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("昵称已存在换一个吧！")
	}
	var roleResp resp.SystemRoleResp
	if roleResp, e = a.roleSrv.Detail(addReq.RoleId, auth); e != nil {
		return
	}
	if roleResp.IsDisable > 0 {
		return response.AssertArgumentError.Make("当前角色已被禁用!")
	}
	passwdLen := len(addReq.Password)
	if !(passwdLen >= 6 && passwdLen <= 20) {
		return response.Failed.Make("密码必须在6~20位")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("你没有权限新增管理员账号!")
	}
	response.Copy(&sysAdmin, addReq)
	sysAdmin.RoleId = addReq.RoleId
	sysAdmin.Password = util.ToolsUtil.MakeMd5(strings.Trim(addReq.Password, "zdm"))
	sysAdmin.TenantId = auth.TenantID
	if addReq.Avatar == "" {
		addReq.Avatar = "/api/static/backend_avatar.png"
	}
	sysAdmin.Avatar = util.UrlUtil.ToRelativeUrl(addReq.Avatar)
	err = a.db.Create(&sysAdmin).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

func (a adminService) Edit(c *gin.Context, editReq req.SystemAdminEditReq, auth *req.AuthReq) (e error) {
	// 检查id
	err := a.db.Where("id = ?", editReq.ID).Limit(1).First(&system.Admin{}).Error
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 检查username
	var admin system.Admin
	r := a.db.Where("username = ? AND id != ?", editReq.Username, editReq.ID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("账号已存在换一个吧！")
	}
	// 检查nickname
	r = a.db.Where("nickname = ? AND id != ? and tenant_id = ?", editReq.Nickname, editReq.ID, auth.TenantID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("昵称已存在换一个吧！")
	}
	// 检查role
	if editReq.RoleId > 0 && editReq.ID != 1 {
		if _, e = a.roleSrv.Detail(editReq.RoleId, auth); e != nil {
			return
		}
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("你没有权限修改管理员账号!")
	}
	// 更新管理员信息
	adminMap := structs.Map(editReq)
	delete(adminMap, "ID")
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(editReq.Avatar)
	role := editReq.RoleId
	if editReq.ID == 1 {
		role = 0
	}
	adminMap["Role"] = strconv.FormatUint(uint64(role), 10)
	if editReq.ID == 1 {
		delete(adminMap, "Username")
	}
	if editReq.Password != "" {
		passwdLen := len(editReq.Password)
		if !(passwdLen >= 6 && passwdLen <= 20) {
			return response.Failed.Make("密码必须在6~20位")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = a.db.Model(&admin).Where("id = ?", editReq.ID).Updates(adminMap).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
	a.CacheAdminUserByUid(editReq.ID)
	// 如果更改自己的密码,则删除旧缓存
	adminId := auth.Id
	if editReq.Password != "" && editReq.ID == adminId {
		token := c.Request.Header.Get("token")
		util.RedisUtil.Del(a.cfg.Admin.BackstageTokenKey + token)
		key := fmt.Sprintf("%d_%d", admin.TenantId, admin.ID)
		adminSetKey := a.cfg.Admin.BackstageTokenSet + key
		ts := util.RedisUtil.SGet(adminSetKey)
		if len(ts) > 0 {
			var tokenKeys []string
			for _, t := range ts {
				tokenKeys = append(tokenKeys, a.cfg.Admin.BackstageTokenKey+t)
			}
			util.RedisUtil.Del(tokenKeys...)
		}
		util.RedisUtil.Del(adminSetKey)
		util.RedisUtil.SSet(adminSetKey, token)
	}
	return
}

func (a adminService) Update(c *gin.Context, updateReq req.SystemAdminUpdateReq, adminId uint, auth *req.AuthReq) (e error) {
	var admin system.Admin
	err := a.db.Where("id = ?", adminId).Limit(1).First(&admin).Error
	if e = response.CheckErr(err, "Update First err"); e != nil {
		return
	}
	if !auth.IsAdmin {
		return response.Failed.Make("你没有权限操作该账号!")
	}
	// 更新管理员信息
	adminMap := structs.Map(updateReq)
	delete(adminMap, "CurrPassword")
	avatar := "/api/static/backend_avatar.png"
	if updateReq.Avatar != "" {
		avatar = updateReq.Avatar
	}
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(avatar)
	delete(adminMap, "aaa")
	if updateReq.Password != "" {
		currPass := util.ToolsUtil.MakeMd5(updateReq.CurrPassword + "zdm")
		if currPass != admin.Password {
			return response.Failed.Make("当前密码不正确!")
		}
		passwdLen := len(updateReq.Password)
		if !(passwdLen >= 6 && passwdLen <= 20) {
			return response.Failed.Make("密码必须在6~20位")
		}
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(updateReq.Password, "zdm"))
	} else {
		delete(adminMap, "Password")
	}
	err = a.db.Model(&admin).Updates(adminMap).Error
	if e = response.CheckErr(err, "Update Updates err"); e != nil {
		return
	}
	a.CacheAdminUserByUid(adminId)
	// 如果更改自己的密码,则删除旧缓存
	if updateReq.Password != "" {
		token := c.Request.Header.Get("token")
		util.RedisUtil.Del(a.cfg.Admin.BackstageTokenKey + token)
		key := fmt.Sprintf("%d_%d", admin.TenantId, admin.ID)
		adminSetKey := a.cfg.Admin.BackstageTokenSet + key
		ts := util.RedisUtil.SGet(adminSetKey)
		if len(ts) > 0 {
			var tokenKeys []string
			for _, t := range ts {
				tokenKeys = append(tokenKeys, a.cfg.Admin.BackstageTokenKey+t)
			}
			util.RedisUtil.Del(tokenKeys...)
		}
		util.RedisUtil.Del(adminSetKey)
		util.RedisUtil.SSet(adminSetKey, token)
	}
	return
}

func (a adminService) Del(id uint, auth *req.AuthReq) (e error) {
	var admin system.Admin
	err := a.db.Where("id = ?", id).Limit(1).First(&admin).Error
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if id == 1 {
		return response.AssertArgumentError.Make("系统管理员不允许删除!")
	}
	if id == auth.Id {
		return response.AssertArgumentError.Make("不能删除自己!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("没有权限!")
	}
	key := fmt.Sprintf("%d_%d", admin.TenantId, admin.ID)
	adminSetKey := a.cfg.Admin.BackstageTokenSet + key
	util.RedisUtil.Del(adminSetKey)
	err = a.db.Model(&admin).Delete(&admin).Error
	e = response.CheckErr(err, "Del Updates err")

	return
}

func (a adminService) Disable(id uint, auth *req.AuthReq) (e error) {
	var admin system.Admin
	err := a.db.Where("id = ?", id).Limit(1).Find(&admin).Error
	if e = response.CheckErr(err, "Disable Find err"); e != nil {
		return
	}
	if admin.ID == 0 {
		return response.AssertArgumentError.Make("账号已不存在!")
	}
	if id == auth.Id {
		return response.AssertArgumentError.Make("不能禁用自己!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("没有权限!")
	}
	admin.IsDisable = ^admin.IsDisable
	err = a.db.Save(&admin).Error
	e = response.CheckErr(err, "Disable Updates err")
	return
}

func (a adminService) CacheAdminUserByUid(id uint) (err error) {
	var admin system.Admin
	err = a.db.Where("id = ?", id).Limit(1).Preload("Role").First(&admin).Error
	if err != nil {
		return
	}
	if admin.TenantId > 0 {
		var tenant system.Tenant
		err = a.db.Where("id = ?", admin.TenantId).First(&tenant).Error
		if err != nil {
			return
		}
		admin.Tenant = &tenant
	}

	str, err := util.ToolsUtil.ObjToJson(&admin)
	if err != nil {
		return
	}
	key := fmt.Sprintf("%d_%d", admin.TenantId, admin.ID)
	util.RedisUtil.HSet(a.cfg.Admin.BackstageManageKey, key, str, 0)
	return nil
}

func NewAdminService(db *gorm.DB, rolePermSrv RolePermService, roleSrv RoleService, cfg *config.Config) AdminService {
	return &adminService{
		db:          db,
		rolePermSrv: rolePermSrv,
		roleSrv:     roleSrv,
		cfg:         cfg,
	}
}
