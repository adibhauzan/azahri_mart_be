package service

import (
	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"gorm.io/gorm"
)

type ProductTypeServiceImpl struct {
	ProductTypeRepository repositories.ProductTypeRepository
	DB                    *gorm.DB
}

func NewProductTypeService(productTypeRepository repositories.ProductTypeRepository, db *gorm.DB) ProductTypeService {
	return &ProductTypeServiceImpl{
		ProductTypeRepository: productTypeRepository,
		DB:                    db,
	}
}

func (service *ProductTypeServiceImpl) Create(request request.CreateProductTypeRequest) response.CreateProductTypeResponse {
	productType := domain.ProductType{
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		productType = service.ProductTypeRepository.Save(productType)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return helper.ToCreateProductTypeResponse(productType)
}
