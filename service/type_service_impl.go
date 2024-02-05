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

func (service *ProductTypeServiceImpl) Create(request request.CreateProductTypeRequest) (response.ProductTypeResponse, error) {
	productType := domain.ProductType{
		ID:   uuid.New(),
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		productType = service.ProductTypeRepository.Save(productType)
		return nil
	})

	if err != nil {
		return response.ProductTypeResponse{}, err
	}

	return helper.ToProductTypeResponse(productType), nil
}

func (service *ProductTypeServiceImpl) Update(request request.UpdateProductTypeRequest) (response.ProductTypeResponse, error) {
	productTypeData := domain.ProductType{
		ID:   request.ID,
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id := productTypeData.ID
		existingId, err := service.ProductTypeRepository.FindById(id)
		if err != nil {
			return err
		}
		productTypeData, err = service.ProductTypeRepository.Update(existingId.ID, productTypeData)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response.ProductTypeResponse{}, err
	}
	return helper.ToProductTypeResponse(productTypeData), nil

}

func (service *ProductTypeServiceImpl) Delete(productTypeId uuid.UUID) error {
	productTypeData := domain.ProductType{
		ID: productTypeId,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id, err := service.ProductTypeRepository.FindById(productTypeData.ID)
		if err != nil {
			return err
		}
		err = service.ProductTypeRepository.Delete(id)
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

func (service *ProductTypeServiceImpl) FindById(productTypeId uuid.UUID) (response.ProductTypeResponse, error) {
	productTypeData := domain.ProductType{
		ID: productTypeId,
	}
	productTypeData, err := service.ProductTypeRepository.FindById(productTypeId)
	if err != nil {
		return response.ProductTypeResponse{}, err
	}

	if err != nil {
		return helper.ToProductTypeResponse(productTypeData), err
	}

	return helper.ToProductTypeResponse(productTypeData), nil

}

func (service *ProductTypeServiceImpl) FindAll() ([]response.ProductTypeResponse, error) {
	var productTypes []domain.ProductType
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		productTypes, err = service.ProductTypeRepository.FindALl()
		return err
	})

	if err != nil {
		return []response.ProductTypeResponse{}, err
	}

	return helper.ToProductTypeResponses(productTypes), nil
}
