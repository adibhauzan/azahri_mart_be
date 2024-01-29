package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

type ProductCategoryService interface {
	Create(request request.CreateProductCategoryRequest) response.CreateProductCategoryResponse
}
