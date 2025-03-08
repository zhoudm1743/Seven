package system

import (
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"
)

type DeptService interface {
	All(auth *req.AuthReq) (res []resp.SystemDeptResp, e error)
	List(listReq req.SystemDeptListReq, auth *req.AuthReq) (mapList []interface{}, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemDeptResp, e error)
	Add(addReq req.SystemDeptAddReq, auth *req.AuthReq) (e error)
	Edit(editReq req.SystemDeptEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type deptService struct {
	db *gorm.DB
}

// All 部门所有
func (a deptService) All(auth *req.AuthReq) (res []resp.SystemDeptResp, e error) {
	var depts []system.Dept
	sql := query.AuthQuery(a.db, auth)
	err := sql.Where("pid > ?", 0).Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []resp.SystemDeptResp{}
	response.Copy(&res, depts)
	return
}

// List 部门列表
func (a deptService) List(listReq req.SystemDeptListReq, auth *req.AuthReq) (mapList []interface{}, e error) {
	deptModel := query.AuthQuery(a.db, auth)
	if listReq.Name != "" {
		deptModel = deptModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsDisable > -1 {
		deptModel = deptModel.Where("is_stop = ?", listReq.IsDisable)
	}
	var depts []system.Dept
	err := deptModel.Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	deptResps := []resp.SystemDeptResp{}
	response.Copy(&deptResps, depts)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(deptResps), "id", "pid", "children")
	return
}

// Detail 部门详情
func (a deptService) Detail(id uint, auth *req.AuthReq) (res resp.SystemDeptResp, e error) {
	var dept system.Dept
	sql := query.AuthQuery(a.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&dept).Error
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dept)
	return
}

// Add 部门新增
func (a deptService) Add(addReq req.SystemDeptAddReq, auth *req.AuthReq) (e error) {
	if addReq.Pid == 0 {
		r := a.db.Where("pid = ?", 0).Limit(1).Find(&system.Dept{})
		if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
			return
		}
		if r.RowsAffected > 0 {
			return response.AssertArgumentError.Make("顶级部门只允许有一个!")
		}
	}
	var dept system.Dept
	response.Copy(&dept, addReq)
	dept.TenantId = auth.TenantID
	err := a.db.Create(&dept).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 部门编辑
func (a deptService) Edit(editReq req.SystemDeptEditReq, auth *req.AuthReq) (e error) {
	var dept system.Dept
	err := a.db.Where("id = ?", editReq.ID).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	if dept.Pid == 0 && editReq.Pid > 0 {
		return response.AssertArgumentError.Make("顶级部门不能修改上级!")
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.Make("上级部门不能是自己!")
	}
	// 更新
	response.Copy(&dept, editReq)
	dept.TenantId = auth.TenantID
	err = a.db.Model(&dept).Updates(dept).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 部门删除
func (a deptService) Del(id uint, auth *req.AuthReq) (e error) {
	var dept system.Dept
	sql := query.AuthQuery(a.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if dept.Pid == 0 {
		return response.AssertArgumentError.Make("顶级部门不能删除!")
	}
	r := sql.Where("pid = ?", id).Limit(1).Find(&system.Dept{})
	if e = response.CheckErr(r.Error, "Del Find dept err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("请先删除子级部门!")
	}
	r = sql.Where("dept_id = ?", id).Limit(1).Find(&system.Admin{})
	if e = response.CheckErr(r.Error, "Del Find admin err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该部门已被管理员使用,请先移除!")
	}
	err = sql.Delete(&dept).Error
	e = response.CheckErr(err, "Del Save err")
	return
}

func NewDeptService(db *gorm.DB) DeptService {
	return &deptService{db: db}
}
