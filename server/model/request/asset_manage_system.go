package request

import "gin-vue-admin/model"

type Asset_manage_systemSearch struct{
    model.Asset_manage_system
    PageInfo
}
type Asset_manage_systemExport struct {
    Iterms []model.Asset_manage_system `json:"iterms"`
    FileName string `json:"fileName"`
}