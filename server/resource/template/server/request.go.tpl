package request

import "slotsevo-admin/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}