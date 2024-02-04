package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductDetailRepository(db *gorm.DB) ProductDetailRepository {
	return &ProductDetailRepositoryImpl{DB: db}
}

func (repository *ProductDetailRepositoryImpl) Save(productDetail domain.ProductDetail) (domain.ProductDetail, error) {
	productDetail.ID = uuid.New()

	err := repository.DB.Create(&productDetail).Error
	if err != nil {
		return domain.ProductDetail{}, err
	}

	err = repository.DB.Model(&domain.Product{}).Where("id = ?", productDetail.ProductID).UpdateColumn("stock", gorm.Expr("stock + ?", 1)).Error
	if err != nil {
		return domain.ProductDetail{}, err
	}

	return productDetail, nil
}
