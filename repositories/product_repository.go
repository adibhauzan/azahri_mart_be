package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(product domain.Product) domain.Product
	FindById(id uuid.UUID) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(id uuid.UUID, productCategory domain.Product) (domain.Product, error)
	Delete(product domain.Product) error
	FindByCategoryId(categoryID uuid.UUID) ([]domain.Product, error)
	FindByTypeId(typeID uuid.UUID) ([]domain.Product, error)
}
