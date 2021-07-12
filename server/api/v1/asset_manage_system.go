package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Asset_manage_system
// @Summary 创建Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "创建Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset_manage_system/createAsset_manage_system [post]
func CreateAsset_manage_system(c *gin.Context) {
	var asset_manage_system model.Asset_manage_system
	_ = c.ShouldBindJSON(&asset_manage_system)
	if err := service.CreateAsset_manage_system(asset_manage_system); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Asset_manage_system
// @Summary 删除Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "删除Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset_manage_system/deleteAsset_manage_system [delete]
func DeleteAsset_manage_system(c *gin.Context) {
	var asset_manage_system model.Asset_manage_system
	_ = c.ShouldBindJSON(&asset_manage_system)
	if err := service.DeleteAsset_manage_system(asset_manage_system); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Asset_manage_system
// @Summary 批量删除Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /asset_manage_system/deleteAsset_manage_systemByIds [delete]
func DeleteAsset_manage_systemByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAsset_manage_systemByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Asset_manage_system
// @Summary 更新Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "更新Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asset_manage_system/updateAsset_manage_system [put]
func UpdateAsset_manage_system(c *gin.Context) {
	var asset_manage_system model.Asset_manage_system
	_ = c.ShouldBindJSON(&asset_manage_system)
	if err := service.UpdateAsset_manage_system(asset_manage_system); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Asset_manage_system
// @Summary 用id查询Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "用id查询Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset_manage_system/findAsset_manage_system [get]
func FindAsset_manage_system(c *gin.Context) {
	var asset_manage_system model.Asset_manage_system
	_ = c.ShouldBindQuery(&asset_manage_system)
	if err, reasset_manage_system := service.GetAsset_manage_system(asset_manage_system.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reasset_manage_system": reasset_manage_system}, c)
	}
}

// @Tags Asset_manage_system
// @Summary 分页获取Asset_manage_system列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Asset_manage_systemSearch true "分页获取Asset_manage_system列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset_manage_system/getAsset_manage_systemList [get]
func GetAsset_manage_systemList(c *gin.Context) {
	var pageInfo request.Asset_manage_systemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAsset_manage_systemInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
//导出指定id值
func ExportAsset_manage_system_resultsByIds(c *gin.Context) {
	//print("exporting..")
	var asset_manage_systemExport request.Asset_manage_systemExport
	_ = c.ShouldBindJSON(&asset_manage_systemExport)
	print(*asset_manage_systemExport.Iterms[0].Is_important_asset)
	filePath := global.GVA_CONFIG.Excel.Dir + asset_manage_systemExport.FileName
	err := service.ExportAsset_manage_system_resultsByIds(asset_manage_systemExport.Iterms, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Any("err", err))
		response.FailWithMessage("转换Excel失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}