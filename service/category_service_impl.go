package service

import (
	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/google/uuid"
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

func (service *ProductCategoryServiceImpl) Create(request request.CreateProductCategoryRequest) (response.ProductCategoryResponse, error) {

	productCategory := domain.ProductCategory{
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		productCategory = service.ProductCategoryRepository.Save(productCategory)
		return nil
	})

	if err != nil {
		return response.ProductCategoryResponse{}, err
	}
	return helper.ToProductCategoryResponse(productCategory), nil
}

func (service *ProductCategoryServiceImpl) Update(request request.UpdateProductCategoryRequest) (response.ProductCategoryResponse, error) {
	productCategoryData := domain.ProductCategory{
		ID:   request.ID,
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id := productCategoryData.ID
		existingId, err := service.ProductCategoryRepository.FindById(id)
		if err != nil {
			return err
		}
		productCategoryData, err = service.ProductCategoryRepository.Update(existingId.ID, productCategoryData)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response.ProductCategoryResponse{}, err
	}

	return helper.ToProductCategoryResponse(productCategoryData), nil

}

func (service *ProductCategoryServiceImpl) Delete(productCategoryId uuid.UUID) error {
	productCategoryData := domain.ProductCategory{
		ID: productCategoryId,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id, err := service.ProductCategoryRepository.FindById(productCategoryData.ID)
		if err != nil {
			return err
		}
		err = service.ProductCategoryRepository.Delete(id)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (service *ProductCategoryServiceImpl) FindById(productCategoryId uuid.UUID) (response.ProductCategoryResponse, error) {
	productCategoryData := domain.ProductCategory{
		ID: productCategoryId,
	}
	productCategoryData, err := service.ProductCategoryRepository.FindById(productCategoryId)
	if err != nil {
		return response.ProductCategoryResponse{}, err
	}

	if err != nil {
		return helper.ToProductCategoryResponse(productCategoryData), err
	}

	return helper.ToProductCategoryResponse(productCategoryData), nil
}

func (service *ProductCategoryServiceImpl) FindAll() ([]response.ProductCategoryResponse, error) {

	var productCategories []domain.ProductCategory
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		productCategories, err = service.ProductCategoryRepository.FindAll()
		return err
	})

	if err != nil {
		return []response.ProductCategoryResponse{}, err
	}

	return helper.ToProductCategoryResponses(productCategories), nil
}
