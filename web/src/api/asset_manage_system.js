import service from '@/utils/request'

// @Tags Asset_manage_system
// @Summary 创建Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "创建Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset_manage_system/createAsset_manage_system [post]
export const createAsset_manage_system = (data) => {
  return service({
    url: '/asset_manage_system/createAsset_manage_system',
    method: 'post',
    data
  })
}

// @Tags Asset_manage_system
// @Summary 删除Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "删除Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset_manage_system/deleteAsset_manage_system [delete]
export const deleteAsset_manage_system = (data) => {
  return service({
    url: '/asset_manage_system/deleteAsset_manage_system',
    method: 'delete',
    data
  })
}

// @Tags Asset_manage_system
// @Summary 删除Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset_manage_system/deleteAsset_manage_system [delete]
export const deleteAsset_manage_systemByIds = (data) => {
  return service({
    url: '/asset_manage_system/deleteAsset_manage_systemByIds',
    method: 'delete',
    data
  })
}

// @Tags Asset_manage_system
// @Summary 更新Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "更新Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asset_manage_system/updateAsset_manage_system [put]
export const updateAsset_manage_system = (data) => {
  return service({
    url: '/asset_manage_system/updateAsset_manage_system',
    method: 'put',
    data
  })
}

// @Tags Asset_manage_system
// @Summary 用id查询Asset_manage_system
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset_manage_system true "用id查询Asset_manage_system"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset_manage_system/findAsset_manage_system [get]
export const findAsset_manage_system = (params) => {
  return service({
    url: '/asset_manage_system/findAsset_manage_system',
    method: 'get',
    params
  })
}

// @Tags Asset_manage_system
// @Summary 分页获取Asset_manage_system列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Asset_manage_system列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset_manage_system/getAsset_manage_systemList [get]
export const getAsset_manage_systemList = (params) => {
  return service({
    url: '/asset_manage_system/getAsset_manage_systemList',
    method: 'get',
    params
  })
}