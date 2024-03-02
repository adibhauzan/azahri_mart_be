package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) *ProductCategoryRepositoryImpl {
	return &ProductCategoryRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductCategoryRepositoryImpl) Save(productCategory domain.ProductCategory) domain.ProductCategory {
	err := repository.DB.Create(&productCategory).Error
	if err != nil {
		return domain.ProductCategory{}
	}
	return productCategory
}

func (repository *ProductCategoryRepositoryImpl) FindById(id uuid.UUID) (domain.ProductCategory, error) {
	var productCategory domain.ProductCategory
	if err := repository.DB.First(&productCategory, id).Error; err != nil {
		return domain.ProductCategory{}, err
	}
	return productCategory, nil
}

func (repository *ProductCategoryRepositoryImpl) FindAll() ([]domain.ProductCategory, error) {
	var productCategory []domain.ProductCategory
	if err := repository.DB.Find(&productCategory).Error; err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (repository *ProductCategoryRepositoryImpl) Update(id uuid.UUID, productCategory domain.ProductCategory) (domain.ProductCategory, error) {
	err := repository.DB.Model(&productCategory).Where("id = ?", id).Update("name", productCategory.Name).Error
	if err != nil {
		return domain.ProductCategory{}, err
	}

	return productCategory, err
}

func (repository *ProductCategoryRepositoryImpl) Delete(productCategory domain.ProductCategory) error {
	existingProductCategory, err := repository.FindById(productCategory.ID)
	if err != nil {
		return err
	}

	err = repository.DB.Delete(&existingProductCategory).Error
	if err != nil {
		return err
	}
	return nil
}
