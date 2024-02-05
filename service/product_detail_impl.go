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
