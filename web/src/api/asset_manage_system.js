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
  // console.log(params)
  return service({
    url: '/asset_manage_system/getAsset_manage_systemList',
    method: 'get',
    params
  })
}

export const exportAsset_manage_system_resultsByIds = (iterms, fileName) => {
  service({
    url: '/asset_manage_system/exportAsset_manage_system_resultsByIds',
    method: 'post',
    data: {
      iterms,
      fileName: fileName
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, fileName)
  })
}


const handleFileError = (res, fileName) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function () {
        const message = JSON.parse(reader.result).msg
        Message({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    var downloadUrl = window.URL.createObjectURL(new Blob([res]))
    var a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = fileName
    var event = new MouseEvent('click')
    a.dispatchEvent(event)
  }
}


export const exportAsset_manage_system_resultsByConditions = (params, fileName) => {
  // console.log(params)
  service({
    url: '/asset_manage_system/exportAsset_manage_system_resultsByConditions',
    method: 'post',
    data: {
      params,
      fileName: fileName
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, fileName)
  })
}

export const loadExcelData = () => {
  return service({
    url: '/asset_manage_system/load_excel',
    method: 'get'
  })
}
