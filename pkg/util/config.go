package util

import (
	"errors"
	"github.com/zhoudm1743/Seven/app/models/system"
	"gorm.io/gorm"
)

var ConfigUtil = configUtil{}

// convertUtil 数据库配置操作工具
type configUtil struct{}

// Get 根据类型和名称获取配置字典
func (cu configUtil) Get(db *gorm.DB, cnfType string, tenantId uint, names ...string) (data map[string]string, err error) {
	chain := db.Where("type = ?", cnfType)
	if len(names) > 0 {
		chain.Where("name = ?", names[0])
	}
	if tenantId > 0 {
		chain.Where("tenant_id = ?", tenantId)
	}
	var configs []system.Config
	err = chain.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	data = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		data[configs[i].Name] = configs[i].Value
	}
	return data, nil
}

// GetVal 根据类型和名称获取配置值
func (cu configUtil) GetVal(db *gorm.DB, cnfType string, name string, defaultVal string, tenantId uint) (data string, err error) {
	config, err := cu.Get(db, cnfType, tenantId, name)
	if err != nil {
		return data, err
	}
	data, ok := config[name]
	if !ok {
		data = defaultVal
	}
	return data, nil
}

// GetMap 根据类型和名称获取配置值(Json字符串转dict)
func (cu configUtil) GetMap(db *gorm.DB, cnfType string, name string, tenantId uint) (data map[string]string, err error) {
	val, err := cu.GetVal(db, cnfType, name, "", tenantId)
	if err != nil {
		return data, err
	}
	if val == "" {
		return map[string]string{}, nil
	}
	err = ToolsUtil.JsonToObj(val, &data)
	return data, err
}

// Set 设置配置的值
func (cu configUtil) Set(db *gorm.DB, cnfType string, name string, val string, tenantId uint) (err error) {
	var config system.Config
	err = db.Where("type = ? AND name = ? and tenant_id = ?", cnfType, name, tenantId).First(&config).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		config.TenantId = tenantId
		config.Type = cnfType
		config.Name = name
		config.Value = val
		if err = db.Create(&config).Error; err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	if err = db.Model(&system.Config{}).Where("type = ? AND name = ? and tenant_id = ?", cnfType, name, tenantId).Update("value", val).Error; err != nil {
		return err
	}
	return nil
}
