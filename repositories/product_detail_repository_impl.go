package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"gorm.io/gorm"
)

type ProductDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductDetailRepository(db *gorm.DB) ProductDetailRepository {
	return &ProductDetailRepositoryImpl{DB: db}
}

func (repository *ProductDetailRepositoryImpl) Save(productDetail domain.ProductDetail) (domain.ProductDetail, error) {
	
	err := repository.DB.Create(&productDetail).Error
	if err != nil {
		return domain.ProductDetail{}, err
	}

	return productDetail, nil
}
