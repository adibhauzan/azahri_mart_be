package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) ProductTypeRepository {
	return &ProductTypeRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductTypeRepositoryImpl) Save(productType domain.ProductType) domain.ProductType {
	productType.ID = uuid.New()
	err := repository.DB.Create(&productType).Error
	if err != nil {
		panic(err)
	}
	return productType
}
