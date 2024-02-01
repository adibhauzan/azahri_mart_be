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
