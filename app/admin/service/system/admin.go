package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"gorm.io/gorm"
)

type AdminService interface {
	FindByUsername(username string) (admin system.Admin, err error)
	FindByTenantIdAndUsername(tenantId uint, username string) (admin system.Admin, err error)
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
}

func (a adminService) FindByUsername(username string) (admin system.Admin, err error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) FindByTenantIdAndUsername(tenantId uint, username string) (admin system.Admin, err error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Self(adminId uint, auth *req.AuthReq) (res resp.SystemAdminSelfResp, e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) List(page req.PageReq, listReq req.SystemAdminListReq, auth *req.AuthReq) (res response.PageResp, e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Detail(id uint, auth *req.AuthReq) (res resp.SystemAdminResp, e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Add(addReq req.SystemAdminAddReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Edit(c *gin.Context, editReq req.SystemAdminEditReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Update(c *gin.Context, updateReq req.SystemAdminUpdateReq, adminId uint, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Del(id uint, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) Disable(id uint, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (a adminService) CacheAdminUserByUid(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewAdminService(db *gorm.DB, rolePermSrv RolePermService, roleSrv RoleService) AdminService {
	return &adminService{
		db:          db,
		rolePermSrv: rolePermSrv,
		roleSrv:     roleSrv,
	}
}
