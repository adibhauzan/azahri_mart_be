package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

type ProductTypeService interface {
	Create(request request.CreateProductTypeRequest) response.CreateProductTypeResponse
}
