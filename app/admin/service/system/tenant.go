package system

import (
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"gorm.io/gorm"
)

type TenantService interface {
	All(auth *req.AuthReq) (res []resp.SystemTenantResp, e error)
	List(page req.PageReq, auth *req.AuthReq) (res response.PageResp, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemTenantResp, e error)
	Add(addReq req.SystemTenantAddReq, auth *req.AuthReq) (e error)
	Edit(editReq req.SystemTenantEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type tenantService struct {
	db        *gorm.DB
	tenantSrv TenantPermService
}

func (t tenantService) All(auth *req.AuthReq) (res []resp.SystemTenantResp, e error) {
	var tenants []system.Tenant
	sql := query.AuthQuery(t.db, auth)
	err := sql.Order("id asc").Find(&tenants).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, tenants)
	return
}

func (t tenantService) List(page req.PageReq, auth *req.AuthReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	sql := query.AuthQuery(t.db.Model(&system.Tenant{}), auth)

	var count int64
	err := sql.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var tenants []system.Tenant
	err = sql.Limit(limit).Offset(offset).Order("id asc").Find(&tenants).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var tenantResp []resp.SystemTenantResp
	response.Copy(&tenantResp, tenants)

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    tenantResp,
	}, nil
}

func (t tenantService) Detail(id uint, auth *req.AuthReq) (res resp.SystemTenantResp, e error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantService) Add(addReq req.SystemTenantAddReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantService) Edit(editReq req.SystemTenantEditReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantService) Del(id uint, auth *req.AuthReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func NewTenantService(db *gorm.DB, tenantSrv TenantPermService) TenantService {
	return &tenantService{db: db, tenantSrv: tenantSrv}
}
