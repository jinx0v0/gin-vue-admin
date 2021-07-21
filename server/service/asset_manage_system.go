package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"reflect"
	_ "time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAsset_manage_system
//@description: 创建Asset_manage_system记录
//@param: asset_manage_system model.Asset_manage_system
//@return: err error

func CreateAsset_manage_system(asset_manage_system model.Asset_manage_system) (err error) {
	err = global.GVA_DB.Create(&asset_manage_system).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAsset_manage_system
//@description: 删除Asset_manage_system记录
//@param: asset_manage_system model.Asset_manage_system
//@return: err error

func DeleteAsset_manage_system(asset_manage_system model.Asset_manage_system) (err error) {
	err = global.GVA_DB.Delete(&asset_manage_system).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAsset_manage_systemByIds
//@description: 批量删除Asset_manage_system记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAsset_manage_systemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Asset_manage_system{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAsset_manage_system
//@description: 更新Asset_manage_system记录
//@param: asset_manage_system *model.Asset_manage_system
//@return: err error

func UpdateAsset_manage_system(asset_manage_system model.Asset_manage_system) (err error) {
	err = global.GVA_DB.Save(&asset_manage_system).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAsset_manage_system
//@description: 根据id获取Asset_manage_system记录
//@param: id uint
//@return: err error, asset_manage_system model.Asset_manage_system

func GetAsset_manage_system(id uint) (err error, asset_manage_system model.Asset_manage_system) {
	err = global.GVA_DB.Where("id = ?", id).First(&asset_manage_system).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAsset_manage_systemInfoList
//@description: 分页获取Asset_manage_system记录
//@param: info request.Asset_manage_systemSearch
//@return: err error, list interface{}, total int64

func GetAsset_manage_systemInfoList(info request.Asset_manage_systemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Asset_manage_system{})
	var asset_manage_systems []model.Asset_manage_system
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Asset_system_name != "" {
		db = db.Where("`asset_system_name` LIKE ?", "%"+info.Asset_system_name+"%")
	}
	if info.Asset_system_manager != "" {
		db = db.Where("`asset_system_manager` = ?", info.Asset_system_manager)
	}
	if info.Asset_system_domain != "" {
		db = db.Where("`asset_system_domain` LIKE ?","%"+ info.Asset_system_domain+"%")
	}
	if info.Extranet_ip != "" {
		db = db.Where("`extranet_ip` LIKE ?", "%"+info.Extranet_ip+"%")
	}
	if info.Extranet_port != "" {
		db = db.Where("`extranet_port` LIKE ?", "%"+info.Extranet_port+"%")
	}
	if info.Intranet_ip != "" {
		db = db.Where("`intranet_ip` LIKE ?", "%"+info.Intranet_ip+"%")
	}
	if info.Intranet_port != "" {
		db = db.Where("`intranet_port` LIKE ?", "%"+info.Intranet_port+"%")
	}
	if info.Is_test_environment != nil {
		db = db.Where("`is_test_environment` = ?", info.Is_test_environment)
	}
	if info.Web_status_code != 0 {
		db = db.Where("`web_status_code` = ?", info.Web_status_code)
	}
	if info.Is_important_asset != nil {
		db = db.Where("`is_important_asset` = ?", info.Is_important_asset)
	}
	if info.More_record != "" {
		db = db.Where("`more_record` LIKE ?", "%"+info.More_record+"%")
	}
	if info.Url != "" {
		db = db.Where("`url` LIKE ?", "%"+info.Url+"%")
	}
	if info.Fingerprint != "" {
		db = db.Where("`fingerprint` LIKE ?", "%"+info.Fingerprint+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&asset_manage_systems).Error
	return err, asset_manage_systems, total
}

func ExportAsset_manage_system_resultsByIds(iterms []model.Asset_manage_system, filePath string) error {

	var Is_important_asset string
	var Is_test_environment string
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"序号", "系统名称", "产品经理（负责人）", "域名", "重点资产", "外网ip", "外网端口", "内网ip", "内网端口", "测试环境", "web状态码", "备注", "时间", "url", "web指纹"}) //
	for i, menu := range iterms {
		//处理tinyint导出问题
		if reflect.ValueOf(menu.Is_important_asset).IsNil() {
			Is_important_asset = "未指定"

		} else {
			if *menu.Is_important_asset {
				Is_important_asset = "是"
			} else {
				Is_important_asset = "否"
			}
		}
		if reflect.ValueOf(menu.Is_test_environment).IsNil() {
			Is_test_environment = "未指定"

		} else {
			if *menu.Is_important_asset {
				Is_test_environment = "是"
			} else {
				Is_test_environment = "否"
			}
		}

		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{

			i + 1,
			menu.Asset_system_name,
			menu.Asset_system_manager,
			menu.Asset_system_domain,
			Is_important_asset, //tinyint得用指针
			menu.Extranet_ip,
			menu.Extranet_port,
			menu.Intranet_ip,
			menu.Intranet_port,
			Is_test_environment, //tinyint得用指针
			menu.Web_status_code,
			menu.More_record,
			menu.UpdatedAt.String(),
			menu.Fingerprint,
			menu.Url,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}

func ExportAsset_manage_system_resultsByConditions(info model.Asset_manage_system, filePath string) (err error) {
	var Is_important_asset string
	var Is_test_environment string
	db := global.GVA_DB.Model(&model.Asset_manage_system{})
	var asset_manage_systems []model.Asset_manage_system
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Asset_system_name != "" {

		db = db.Where("`asset_system_name` LIKE ?", "%"+info.Asset_system_name+"%")
	}
	if info.Asset_system_manager != "" {
		db = db.Where("`asset_system_manager` = ?", info.Asset_system_manager)
	}
	if info.Asset_system_domain != "" {
		print("domain key:"+info.Asset_system_domain)
		db = db.Where("`asset_system_domain` LIKE ?", "%"+info.Asset_system_domain+"%")
	}
	if info.Extranet_ip != "" {
		db = db.Where("`extranet_ip` LIKE ?", "%"+info.Extranet_ip+"%")
	}
	if info.Extranet_port != "" {
		db = db.Where("`extranet_port` LIKE ?", "%"+info.Extranet_port+"%")
	}
	if info.Intranet_ip != "" {
		db = db.Where("`intranet_ip` LIKE ?", "%"+info.Intranet_ip+"%")
	}
	if info.Intranet_port != "" {
		db = db.Where("`intranet_port` LIKE ?", "%"+info.Intranet_port+"%")
	}
	if info.Is_test_environment != nil {
		db = db.Where("`is_test_environment` = ?", info.Is_test_environment)
	}
	if info.Web_status_code != 0 {
		db = db.Where("`web_status_code` = ?", info.Web_status_code)
	}
	if info.Is_important_asset != nil {
		db = db.Where("`is_important_asset` = ?", info.Is_important_asset)
	}
	if info.More_record != "" {
		db = db.Where("`more_record` LIKE ?", "%"+info.More_record+"%")
	}
	if info.Url != "" {
		db = db.Where("`url` LIKE ?", "%"+info.Url+"%")
	}
	if info.Fingerprint != "" {

		db = db.Where("`fingerprint` LIKE ?", "%"+info.Fingerprint+"%")
	}
	//err = db.Count(&total).Error
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"序号", "系统名称", "产品经理（负责人）", "域名", "重点资产", "外网ip", "外网端口", "内网ip", "内网端口", "测试环境", "web状态码", "备注", "时间", "url", "web指纹"}) //"时间"
	err = db.Find(&asset_manage_systems).Error
	for i, menu := range asset_manage_systems {

		//处理tinyint导出问题
		if reflect.ValueOf(menu.Is_important_asset).IsNil() {
			Is_important_asset = "未指定"

		} else {
			if *menu.Is_important_asset {
				Is_important_asset = "是"
			} else {
				Is_important_asset = "否"
			}
		}
		if reflect.ValueOf(menu.Is_test_environment).IsNil() {
			Is_test_environment = "未指定"

		} else {
			if *menu.Is_important_asset {
				Is_test_environment = "是"
			} else {
				Is_test_environment = "否"
			}
		}
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			i + 1,
			menu.Asset_system_name,
			menu.Asset_system_manager,
			menu.Asset_system_domain,
			Is_important_asset, //tinyint得用指针
			menu.Extranet_ip,
			menu.Extranet_port,
			menu.Intranet_ip,
			menu.Intranet_port,
			Is_test_environment, //tinyint得用指针
			menu.Web_status_code,
			menu.More_record,
			menu.UpdatedAt.String(),
			menu.Fingerprint,
			menu.Url,
		})
	}
	err = excel.SaveAs(filePath)
	return err
}

func ParseExcel2assetList() ([]model.Asset_manage_system, error) {

	is_test_environment := false
	skipHeader := true
	fixedHeader := []string{"系统名称", "产品经理（产品线-负责人）", "域名", "外网ip", "外网ip端口", "内网ip","内网ip端口","url","是否测试环境（是/否）","备注"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "外部资产导入模板.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]model.Asset_manage_system, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var menu model.Asset_manage_system
		row, err := rows.Columns()
		print(len(row)) //test
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		//if len(row) != len(fixedHeader) {
		//	continue
		//}
		if len(row) < len(fixedHeader) {
			for i:=1;i<11;i++{
				row = append(row,"")
			}
		}
		if len(row) < len(fixedHeader) {
			continue
		}
		if row[8] == "是"{

			is_test_environment = true
			menu = model.Asset_manage_system{
				//GVA_MODEL.CreatedAt
				Asset_system_name:      row[0],
				Asset_system_manager:      row[1],
				Asset_system_domain:    row[2],
				Extranet_ip:  row[3],
				Extranet_port:      row[4],
				Intranet_ip: row[5],
				Intranet_port: row[6],
				Url: row[7],
				Is_test_environment: &is_test_environment,
				More_record: row[9],
			}
		}else if row[8] == "否"{
			is_test_environment = false
			menu = model.Asset_manage_system{
				//GVA_MODEL.CreatedAt
				Asset_system_name:      row[0],
				Asset_system_manager:      row[1],
				Asset_system_domain:    row[2],
				Extranet_ip:  row[3],
				Extranet_port:      row[4],
				Intranet_ip: row[5],
				Intranet_port: row[6],
				Url: row[7],
				Is_test_environment: &is_test_environment,
				More_record: row[9],
			}
		}else {
			menu = model.Asset_manage_system{
			//GVA_MODEL.CreatedAt
			Asset_system_name:      row[0],
			Asset_system_manager:      row[1],
			Asset_system_domain:    row[2],
			Extranet_ip:  row[3],
			Extranet_port:      row[4],
			Intranet_ip: row[5],
			Intranet_port: row[6],
			Url: row[7],
			More_record: row[9],
		}}
		menus = append(menus, menu)
	}
	//print("test import..")
	//for i,menu := range menus{
	//	print(i)
	//	print(menu.Asset_system_name)
	//}
	return menus, nil
}


