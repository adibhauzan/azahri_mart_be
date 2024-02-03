package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/google/uuid"
)

type ProductService interface {
	Create(request request.CreateProductRequest) (response.ProductResponse, error)
	Update(request request.UpdateProductRequest) (response.ProductResponse, error)
	Delete(productCategoryId uuid.UUID) error
	FindById(productCategoryId uuid.UUID) (response.ProductResponse, error)
	FindAll() ([]response.ProductResponse, error)
	FindByCategoryId(categoryID uuid.UUID) ([]response.ProductResponse, error)
	FindByTypeId(typeID uuid.UUID) ([]response.ProductResponse, error)
}
