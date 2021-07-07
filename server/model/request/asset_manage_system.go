package request

import "gin-vue-admin/model"

type Asset_manage_systemSearch struct{
    model.Asset_manage_system
    PageInfo
}