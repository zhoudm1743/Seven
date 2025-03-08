package system

import (
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"gorm.io/gorm"
)

type PostService interface {
	All(auth *req.AuthReq) (res []resp.SystemPostResp, e error)
	List(page req.PageReq, listReq req.SystemPostListReq, auth *req.AuthReq) (res response.PageResp, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemPostResp, e error)
	Add(addReq req.SystemPostAddReq, auth *req.AuthReq) (e error)
	Edit(editReq req.SystemPostEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type postService struct {
	db *gorm.DB
}

func (a postService) All(auth *req.AuthReq) (res []resp.SystemPostResp, e error) {
	//TODO implement me
	var posts []system.Post
	err := a.db.Order("sort desc, id desc").Find(&posts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []resp.SystemPostResp{}
	response.Copy(&res, posts)
	return
}

func (a postService) List(page req.PageReq, listReq req.SystemPostListReq, auth *req.AuthReq) (res response.PageResp, e error) {
	//TODO implement me
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	postModel := query.AuthQuery(a.db.Model(&system.Post{}), auth)
	if listReq.Code != "" {
		postModel = postModel.Where("code like ?", "%"+listReq.Code+"%")
	}
	if listReq.Name != "" {
		postModel = postModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsDisable > -1 {
		postModel = postModel.Where("is_stop = ?", listReq.IsDisable)
	}
	// 总数
	var count int64
	err := postModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var posts []system.Post
	err = postModel.Limit(limit).Offset(offset).Order("id desc").Find(&posts).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	postResps := []resp.SystemPostResp{}
	response.Copy(&postResps, posts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    postResps,
	}, nil
}

func (a postService) Detail(id uint, auth *req.AuthReq) (res resp.SystemPostResp, e error) {
	//TODO implement me
	var post system.Post
	sql := query.AuthQuery(a.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&post).Error
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, post)
	return
}

func (a postService) Add(addReq req.SystemPostAddReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	r := a.db.Where("code = ? OR name = ?", addReq.Code, addReq.Name).Limit(1).Find(&system.Post{})
	if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位已存在!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("无权限!")
	}
	var post system.Post
	response.Copy(&post, addReq)
	post.TenantId = auth.TenantID
	err := a.db.Create(&post).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

func (a postService) Edit(editReq req.SystemPostEditReq, auth *req.AuthReq) (e error) {
	//TODO implement me
	var post system.Post
	err := a.db.Where("id = ?", editReq.ID).Limit(1).First(&post).Error
	// 校验
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	r := a.db.Where("(code = ? OR name = ?) AND id != ?", editReq.Code, editReq.Name, editReq.ID).Limit(1).Find(&system.Post{})
	if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位已存在!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("无权限!")
	}
	// 更新
	response.Copy(&post, editReq)
	err = a.db.Model(&system.Post{}).Updates(&post).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

func (a postService) Del(id uint, auth *req.AuthReq) (e error) {
	//TODO implement me
	var post system.Post
	err := a.db.Where("id = ?", id).Limit(1).First(&post).Error
	// 校验
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	r := a.db.Where("post_id = ?", id).Limit(1).Find(&system.Admin{})
	if e = response.CheckErr(r.Error, "Del Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位已被管理员使用,请先移除!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("无权限!")
	}
	err = a.db.Delete(&post).Error
	e = response.CheckErr(err, "Del Save err")
	return
}

func NewPostService(db *gorm.DB) PostService {
	return &postService{
		db: db,
	}
}
