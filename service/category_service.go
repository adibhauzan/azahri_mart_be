package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/google/uuid"
)

type ProductCategoryService interface {
	Create(request request.CreateProductCategoryRequest) (response.ProductCategoryResponse, error)
	Update(request request.UpdateProductCategoryRequest) (response.ProductCategoryResponse, error)
	Delete(productCategoryId uuid.UUID) error
	FindById(productCategoryId uuid.UUID) (response.ProductCategoryResponse, error)
	FindAll() ([]response.ProductCategoryResponse, error)
}
