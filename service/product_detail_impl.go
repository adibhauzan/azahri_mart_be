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

type ProductDetailServiceImpl struct {
	DB                      *gorm.DB
	ProductDetailRepository repositories.ProductDetailRepository
	ProductRepository       repositories.ProductRepository
}

func NewProductDetailService(db *gorm.DB, productDetailRepo repositories.ProductDetailRepository, productRepo repositories.ProductRepository) ProductDetailService {
	return &ProductDetailServiceImpl{
		DB:                      db,
		ProductDetailRepository: productDetailRepo,
		ProductRepository:       productRepo,
	}
}

func (service *ProductDetailServiceImpl) Create(request request.CreateProductDetailRequest) (response.ProductDetailResponse, error) {
	product, err := service.ProductRepository.FindById(request.ProductID)
	if err != nil {
		return response.ProductDetailResponse{}, fmt.Errorf("not found product : %v", err)
	}

	productDetail := domain.ProductDetail{
		ID:        uuid.New(),
		ProductID: request.ProductID,
		ExpiredAt: request.ExpiredAt,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		productDetail, err = service.ProductDetailRepository.Save(productDetail)
		if err != nil {
			return fmt.Errorf("failed to save product detail: %v", err)
		}

		product.Stock++
		_, err = service.ProductRepository.Update(product.ID, product)
		if err != nil {
			return fmt.Errorf("failed to update product stock: %v", err)
		}
		return nil
	})
	if err != nil {
		return response.ProductDetailResponse{}, err
	}
	return helper.ToProductDetailResponse(productDetail), nil
}

func (service *ProductDetailServiceImpl) Update(request request.UpdateProductDetailRequest) (response.ProductDetailResponse, error) {
	product, err := service.ProductRepository.FindById(request.ProductID)
	if err != nil {
		return response.ProductDetailResponse{}, fmt.Errorf("not found product: %v", err)
	}

	productDetailData := domain.ProductDetail{
		ID:        request.ID,
		ProductID: product.ID,
		ExpiredAt: request.ExpiredAt,
		// Copy other fields as needed
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		// Fetch existing product detail
		existingDetail, err := service.ProductDetailRepository.FindById(request.ID)
		if err != nil {
			return fmt.Errorf("not found product detail: %v", err)
		}

		// Update product stock
		product.Stock--
		_, err = service.ProductRepository.Update(product.ID, product)
		if err != nil {
			return fmt.Errorf("failed to update product stock: %v", err)
		}

		// Update product detail
		productDetailData, err = service.ProductDetailRepository.Update(existingDetail.ID, productDetailData)
		if err != nil {
			return fmt.Errorf("failed to update product detail: %v", err)
		}

		return nil
	})

	if err != nil {
		return response.ProductDetailResponse{}, err
	}

	return helper.ToProductDetailResponse(productDetailData), nil
}


func (service *ProductDetailServiceImpl) Delete(productDetailId uuid.UUID) error {
	product, err := service.ProductRepository.FindById(productDetailId)
	if err != nil {
		return fmt.Errorf("failed to find product: %v", err)
	}

	productData := domain.Product{
		ID:    product.ID,
		Stock: product.Stock - 1, 
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		err := service.ProductRepository.Delete(product)
		if err != nil {
			return fmt.Errorf("failed to delete product details: %v", err)
		}

		_, err = service.ProductRepository.Update(product.ID, productData)
		if err != nil {
			return fmt.Errorf("failed to update product stock: %v", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}

	return nil
}


func (service *ProductDetailServiceImpl) FindById(productDetailId uuid.UUID) (response.ProductDetailResponse, error) {
	productDetailData, err := service.ProductDetailRepository.FindById(productDetailId)
	if err != nil {
		return response.ProductDetailResponse{}, fmt.Errorf("not found product detail: %v", err)
	}

	return helper.ToProductDetailResponse(productDetailData), nil
}

func (service *ProductDetailServiceImpl) FindAll() ([]response.ProductDetailResponse, error) {
	productDetails, err := service.ProductDetailRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all product details: %v", err)
	}

	return helper.ToProductDetailResponses(productDetails), nil
}

func (service *ProductDetailServiceImpl) FindByProductId(productDetailId uuid.UUID) ([]response.ProductDetailResponse, error) {
	productsDetail, err := service.ProductDetailRepository.FindByProductId(productDetailId)
	if err != nil {
		return nil, fmt.Errorf("failed to find product details by product ID: %v", err)
	}

	return helper.ToProductDetailResponses(productsDetail), nil
}
