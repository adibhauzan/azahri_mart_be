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

func ToProductTypeResponse(productType domain.ProductType) response.ProductTypeResponse {
	return response.ProductTypeResponse{
		ID:        productType.ID,
		Name:      productType.Name,
		CreatedAt: productType.CreatedAt,
		UpdatedAt: productType.UpdatedAt,
	}
}

func ToProductTypeResponses(productTypes []domain.ProductType) []response.ProductTypeResponse {
	var productTypeResponses []response.ProductTypeResponse
	for _, productTypeResponse := range productTypes {
		productTypeResponses = append(productTypeResponses, ToProductTypeResponse(productTypeResponse))
	}
	return productTypeResponses
}

func ToProductResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		CategoryID: product.CategoryID,
		TypeID:     product.TypeID,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}
