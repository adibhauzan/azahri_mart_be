package service

import (
	"fmt"

	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB                 *gorm.DB
	ProductRepository  repositories.ProductRepository
	CategoryRepository repositories.ProductCategoryRepository
	TypeRepository     repositories.ProductTypeRepository
}

func NewProductService(productRepo repositories.ProductRepository, categoryRepo repositories.ProductCategoryRepository, typeRepo repositories.ProductTypeRepository, db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepository:  productRepo,
		DB:                 db,
		CategoryRepository: categoryRepo,
		TypeRepository:     typeRepo,
	}
}
func (service *ProductServiceImpl) Create(request request.CreateProductRequest) (response.ProductResponse, error) {

	_, err := service.CategoryRepository.FindById(request.CategoryID)
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("not found product category: %v", err)
	}

	// Check if the type_id exists
	_, err = service.TypeRepository.FindById(request.TypeID)
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("not found product type: %v", err)
	}

	product := domain.Product{
		ID:         uuid.New(),
		CategoryID: request.CategoryID,
		TypeID:     request.TypeID,
		Name:       request.Name,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		product = service.ProductRepository.Save(product)
		return nil
	})
	if err != nil {
		return response.ProductResponse{}, err
	}
	return helper.ToProductResponse(product), nil
}

func (service *ProductServiceImpl) Update(request request.UpdateProductRequest) (response.ProductResponse, error) {
	productData := domain.Product{
		ID:         request.ID,
		Name:       request.Name,
		CategoryID: request.CategoryID,
		TypeID:     request.TypeID,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id := productData.ID
		existingId, err := service.ProductRepository.FindById(id)
		if err != nil {
			return fmt.Errorf("not found product: %v", err)
		}
		productData, err = service.ProductRepository.Update(existingId.ID, productData)
		if err != nil {
			return fmt.Errorf("id product not found: %v", err)
		}
		return nil
	})
	if err != nil {
		return response.ProductResponse{}, err
	}
	return helper.ToProductResponse(productData), nil
}

func (service *ProductServiceImpl) Delete(productId uuid.UUID) error {
	productData := domain.Product{
		ID: productId,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		id, err := service.ProductRepository.FindById(productData.ID)
		if err != nil {
			return fmt.Errorf("not found product: %v", err)

		}
		err = service.ProductRepository.Delete(id)
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

func (service *ProductServiceImpl) FindById(productId uuid.UUID) (response.ProductResponse, error) {
	var productData domain.Product
	productData.ID = productId
	productData, err := service.ProductRepository.FindById(productId)
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("not found product: %v", err)

	}

	if err != nil {
		return helper.ToProductResponse(productData), err
	}

	return helper.ToProductResponse(productData), nil
}

func (service *ProductServiceImpl) FindAll() ([]response.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return helper.ToProductResponses(products), nil
}

func (service *ProductServiceImpl) FindByCategoryId(categoryID uuid.UUID) ([]response.ProductResponse, error) {
	products, err := service.ProductRepository.FindByCategoryId(categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by category ID: %v", err)
	}

	return helper.ToProductResponses(products), nil
}

func (service *ProductServiceImpl) FindByTypeId(typeID uuid.UUID) ([]response.ProductResponse, error) {
	products, err := service.ProductRepository.FindByTypeId(typeID)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by type ID: %v", err)
	}

	return helper.ToProductResponses(products), nil
}
