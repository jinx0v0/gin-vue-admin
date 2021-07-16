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

type Asset_manage_systemExport_Conditions struct {
    Params model.Asset_manage_system `json:"params"`
    FileName string `json:"fileName"`
}

type Asset_manage_systemImport_Excel struct {
    Params model.Asset_manage_system `json:"params"`

}