package database

import (
	"github.com/zhoudm1743/Seven/app/models/system"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	models := dst()
	err := db.AutoMigrate(models...)
	return err
}

func dst() []interface{} {
	return []interface{}{
		// system
		&system.Tenant{},
		&system.TenantPerm{},
		&system.Role{},
		&system.RolePerm{},
		&system.Menu{},
		&system.Dept{},
		&system.Post{},
		&system.Admin{},
		system.Config{},
	}
}
