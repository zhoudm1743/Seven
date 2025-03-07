package query

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"seven-admin/app/admin/schemas/req"
)

func AuthQuery(db *gorm.DB, auth *req.AuthReq) *gorm.DB {
	if auth == nil {
		return db
	}
	if auth.IsAdmin || auth.IsSuperTenant {
		return db
	} else {
		return db.Where("tenant_id = ?", auth.TenantID)
	}
}

func GetAuthReq(ctx *gin.Context) *req.AuthReq {
	auth := ctx.MustGet("auth").(*req.AuthReq)
	if auth == nil {
		color.Redln("GetAuthReq error", auth)
		return nil
	}
	return auth
}
