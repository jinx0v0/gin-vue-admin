// 自动生成模板Asset_manage_system
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Asset_manage_system struct {
      global.GVA_MODEL
      Asset_system_name  string `json:"asset_system_name" form:"asset_system_name" gorm:"column:asset_system_name;comment:"`
      Asset_system_manager  string `json:"asset_system_manager" form:"asset_system_manager" gorm:"column:asset_system_manager;comment:"`
      Asset_system_domain  string `json:"asset_system_domain" form:"asset_system_domain" gorm:"column:asset_system_domain;comment:"`
      Extranet_ip  string `json:"extranet_ip" form:"extranet_ip" gorm:"column:extranet_ip;comment:"`
      Extranet_port  string `json:"extranet_port" form:"extranet_port" gorm:"column:extranet_port;comment:"`
      Intranet_ip  string `json:"intranet_ip" form:"intranet_ip" gorm:"column:intranet_ip;comment:"`
      Intranet_port  string `json:"intranet_port" form:"intranet_port" gorm:"column:intranet_port;comment:"`
      Is_test_environment  *bool `json:"is_test_environment" form:"is_test_environment" gorm:"column:is_test_environment;comment:"`
      Web_status_code  int `json:"web_status_code" form:"web_status_code" gorm:"column:web_status_code;comment:"`
      Web_screenshot  string `json:"web_screenshot" form:"web_screenshot" gorm:"column:web_screenshot;comment:"`
      Is_important_asset  *bool `json:"is_important_asset" form:"is_important_asset" gorm:"column:is_important_asset;comment:"`
      More_record  string `json:"more_record" form:"more_record" gorm:"column:more_record;comment:"`
}


func (Asset_manage_system) TableName() string {
  return "Asset_manage_system"
}

