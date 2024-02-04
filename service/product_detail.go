package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

type ProductDetailService interface {
	Create(request request.CreateProductDetailRequest) (response.ProductDetailResponse, error)
	// Update(request request.UpdateProductDetailRequest) (response.ProductDetailResponse, error)
	// Delete(productDetailId uuid.UUID) error
	// FindById(productDetailId uuid.UUID) (response.ProductDetailResponse, error)
	// FindAll() ([]response.ProductDetailResponse, error)
	// FindByProductId(productId uuid.UUID) ([]response.ProductDetailResponse, error)
}
