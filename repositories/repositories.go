package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
)

type ProductCategoryRepository interface {
	Save(productCategory domain.ProductCategory) domain.ProductCategory
	FindById(id uuid.UUID) (domain.ProductCategory, error)
	FindAll() ([]domain.ProductCategory, error)
	Update(id uuid.UUID, productCategory domain.ProductCategory) (domain.ProductCategory, error)
	Delete(productCategory domain.ProductCategory) error
}

type ProductDetailRepository interface {
	Save(productDetail domain.ProductDetail) (domain.ProductDetail, error)
	FindById(id uuid.UUID) (domain.ProductDetail, error)
	FindAll() ([]domain.ProductDetail, error)
	Update(id uuid.UUID, productDetail domain.ProductDetail) (domain.ProductDetail, error)
	Delete(productDetail domain.ProductDetail) error
	FindByProductId(productId uuid.UUID) ([]domain.ProductDetail, error)
}

type ProductRepository interface {
	Save(product domain.Product) domain.Product
	FindById(id uuid.UUID) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(id uuid.UUID, product domain.Product) (domain.Product, error)
	Delete(product domain.Product) error
	FindByCategoryId(categoryID uuid.UUID) ([]domain.Product, error)
	FindByTypeId(typeID uuid.UUID) ([]domain.Product, error)
}

type ProductTypeRepository interface {
	Save(productType domain.ProductType) domain.ProductType
	FindById(id uuid.UUID) (domain.ProductType, error)
	FindALl() ([]domain.ProductType, error)
	Update(id uuid.UUID, productType domain.ProductType) (domain.ProductType, error)
	Delete(productType domain.ProductType) error
}

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
}
