package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/google/uuid"
)

type ProductTypeService interface {
	Create(request request.CreateProductTypeRequest) (response.ProductTypeResponse, error)
	Update(request request.UpdateProductTypeRequest) (response.ProductTypeResponse, error)
	Delete(productTypeId uuid.UUID) error
	FindById(productTypeId uuid.UUID) (response.ProductTypeResponse, error)
	FindAll() ([]response.ProductTypeResponse, error)
}
