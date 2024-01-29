package helper

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

func ToCreateProductCategoryResponse(productCategory domain.ProductCategory) response.CreateProductCategoryResponse {
	return response.CreateProductCategoryResponse{
		ID:        productCategory.ID,
		Name:      productCategory.Name,
		CreatedAt: productCategory.CreatedAt,
		UpdatedAt: productCategory.UpdatedAt,
	}
}

func ToCreateProductTypeResponse(productType domain.ProductType) response.CreateProductTypeResponse {
	return response.CreateProductTypeResponse{
		ID:        productType.ID,
		Name:      productType.Name,
		CreatedAt: productType.CreatedAt,
		UpdatedAt: productType.UpdatedAt,
	}
}
