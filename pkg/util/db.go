package util

import "gorm.io/gorm"

var DbUtil = dbUtil{}

type dbUtil struct{}

func (dbUtil) DBTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema.Table
}
