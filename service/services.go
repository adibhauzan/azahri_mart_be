package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/google/uuid"
)

type ProductCategoryService interface {
	Create(request request.CreateProductCategoryRequest) (response.ProductCategoryResponse, error)
	Update(request request.UpdateProductCategoryRequest) (response.ProductCategoryResponse, error)
	Delete(productCategoryId uuid.UUID) error
	FindById(productCategoryId uuid.UUID) (response.ProductCategoryResponse, error)
	FindAll() ([]response.ProductCategoryResponse, error)
}

type ProductDetailService interface {
	Create(request request.CreateProductDetailRequest) (response.ProductDetailResponse, error)
	Update(request request.UpdateProductDetailRequest) (response.ProductDetailResponse, error)
	Delete(productDetailId uuid.UUID) error
	FindById(productDetailId uuid.UUID) (response.ProductDetailResponse, error)
	FindAll() ([]response.ProductDetailResponse, error)
	FindByProductId(productDetailId uuid.UUID) ([]response.ProductDetailResponse, error)
}

type ProductService interface {
	Create(request request.CreateProductRequest) (response.ProductResponse, error)
	Update(request request.UpdateProductRequest) (response.ProductResponse, error)
	Delete(productId uuid.UUID) error
	FindById(productId uuid.UUID) (response.ProductResponse, error)
	FindAll() ([]response.ProductResponse, error)
	FindByCategoryId(categoryID uuid.UUID) ([]response.ProductResponse, error)
	FindByTypeId(typeID uuid.UUID) ([]response.ProductResponse, error)
}

type ProductTypeService interface {
	Create(request request.CreateProductTypeRequest) (response.ProductTypeResponse, error)
	Update(request request.UpdateProductTypeRequest) (response.ProductTypeResponse, error)
	Delete(productTypeId uuid.UUID) error
	FindById(productTypeId uuid.UUID) (response.ProductTypeResponse, error)
	FindAll() ([]response.ProductTypeResponse, error)
}

type UserService interface {
	Create(request request.CreateUserRequest) (response.UserResponse, error)
	/*

		for Update and Delete methods later after implementing jwt authorization
		
		/ this method for user with role admin
			FindById(userId uuid.UUID) (response.UserResponse, error)
			FindAll() ([]response.UserResponse, error)
		/



	*/
}