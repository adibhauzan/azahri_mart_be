package helper

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

func ToProductCategoryResponse(productCategory domain.ProductCategory) response.ProductCategoryResponse {
	return response.ProductCategoryResponse{
		ID:        productCategory.ID,
		Name:      productCategory.Name,
		CreatedAt: productCategory.CreatedAt,
		UpdatedAt: productCategory.UpdatedAt,
	}
}

func ToProductCategoryResponses(productCategories []domain.ProductCategory) []response.ProductCategoryResponse {
	var productCategoryResponses []response.ProductCategoryResponse
	for _, productCategoryResponse := range productCategories {
		productCategoryResponses = append(productCategoryResponses, ToProductCategoryResponse(productCategoryResponse))
	}
	return productCategoryResponses
}

func ToCreateProductTypeResponse(productType domain.ProductType) response.CreateProductTypeResponse {
	return response.CreateProductTypeResponse{
		ID:        productType.ID,
		Name:      productType.Name,
		CreatedAt: productType.CreatedAt,
		UpdatedAt: productType.UpdatedAt,
	}
}
