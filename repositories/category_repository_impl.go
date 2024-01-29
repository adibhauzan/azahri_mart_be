package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) ProductCategoryRepository {
	return &ProductCategoryRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductCategoryRepositoryImpl) Save(productCategory domain.ProductCategory) domain.ProductCategory {
	productCategory.ID = uuid.New()
	err := repository.DB.Create(&productCategory).Error
	if err != nil {
		panic(err)
	}
	return productCategory
}
