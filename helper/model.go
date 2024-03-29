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
		Stock:      product.Stock,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponses(product []domain.Product) []response.ProductResponse {
	var productResponses []response.ProductResponse
	for _, productResponse := range product {
		productResponses = append(productResponses, ToProductResponse(productResponse))
	}
	return productResponses
}

func ToProductDetailResponse(productDetail domain.ProductDetail) response.ProductDetailResponse {
	return response.ProductDetailResponse{
		ID:        productDetail.ID,
		ProductID: productDetail.ProductID,
		ExpiredAt: productDetail.ExpiredAt,
		CreatedAt: productDetail.CreatedAt,
		UpdatedAt: productDetail.UpdatedAt,
	}
}

func ToProductDetailResponses(productDetail []domain.ProductDetail) []response.ProductDetailResponse {
	var productDetailResponses []response.ProductDetailResponse
	for _, productResponse := range productDetail {
		productDetailResponses = append(productDetailResponses, ToProductDetailResponse(productResponse))
	}
	return productDetailResponses
}

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Role:      string(user.Role),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(user []domain.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, userResponse := range user {
		userResponses = append(userResponses, ToUserResponse(userResponse))
	}
	return userResponses
}
