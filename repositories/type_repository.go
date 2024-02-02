package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
)

type ProductTypeRepository interface {
	Save(productType domain.ProductType) domain.ProductType
	FindById(id uuid.UUID) (domain.ProductType, error)
	FindALl() ([]domain.ProductType, error)
	Update(id uuid.UUID, productType domain.ProductType) (domain.ProductType, error)
	Delete(productType domain.ProductType) error
}
