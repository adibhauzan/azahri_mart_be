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

	err := repository.DB.Create(&productDetail).Error
	if err != nil {
		return domain.ProductDetail{}, err
	}

	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) FindById(id uuid.UUID) (domain.ProductDetail, error) {
	var productDetail domain.ProductDetail
	if err := repository.DB.First(&productDetail, id).Error; err != nil {
		return domain.ProductDetail{}, err
	}
	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) FindAll() ([]domain.ProductDetail, error) {
	var productDetail []domain.ProductDetail
	if err := repository.DB.Find(&productDetail).Error; err != nil{
		return nil, err
	}
	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) Update(id uuid.UUID, productDetail domain.ProductDetail) (domain.ProductDetail, error) {
	if err := repository.DB.Model(&productDetail).Where("id = ?", id).Updates(productDetail).Error; err != nil{
		return domain.ProductDetail{}, err
	}
	return productDetail, nil
}


func (repository *ProductDetailRepositoryImpl) Delete(productDetail domain.ProductDetail) error {
	existingProductDetail, err := repository.FindById(productDetail.ID)
	if err != nil{
		return err
	}
	
	if err = repository.DB.Delete(&existingProductDetail).Error; err != nil{
		return err
	}
	return nil

}
func (repository *ProductDetailRepositoryImpl) FindByProductId(productId uuid.UUID) ([]domain.ProductDetail, error) {
	var productDetails []domain.ProductDetail
	if err := repository.DB.Preload("Product").Where("product_id = ?", productId).Find(&productDetails).Error; err != nil {
		return nil, err
	}
	return productDetails, nil
}
