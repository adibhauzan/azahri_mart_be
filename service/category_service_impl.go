package service

import (
	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"gorm.io/gorm"
)

type ProductCategoryServiceImpl struct {
	ProductCategoryRepository repositories.ProductCategoryRepository
	DB                        *gorm.DB
}

func NewCategoryService(categoryRepo repositories.ProductCategoryRepository, db *gorm.DB) ProductCategoryService {
	return &ProductCategoryServiceImpl{
		ProductCategoryRepository: categoryRepo,
		DB:                        db,
	}
}

func (service *ProductCategoryServiceImpl) Create(request request.CreateProductCategoryRequest) response.CreateProductCategoryResponse {

	productCategory := domain.ProductCategory{
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		productCategory = service.ProductCategoryRepository.Save(productCategory)
		return nil
	})

	if err != nil {
		panic(err)
	}
	return helper.ToCreateProductCategoryResponse(productCategory)
}
