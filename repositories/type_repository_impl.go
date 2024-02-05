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
	
	err := repository.DB.Create(&productType).Error
	if err != nil {
		return domain.ProductType{}
	}
	return productType
}

func (repository *ProductTypeRepositoryImpl) FindById(id uuid.UUID) (domain.ProductType, error) {
	var productType domain.ProductType
	if err := repository.DB.First(&productType, id).Error; err != nil {
		return domain.ProductType{}, err
	}
	return productType, nil
}

func (repository *ProductTypeRepositoryImpl) FindALl() ([]domain.ProductType, error) {
	var productType []domain.ProductType
	if err := repository.DB.Find(&productType).Error; err != nil {
		return nil, err
	}
	return productType, nil
}

func (repository *ProductTypeRepositoryImpl) Update(id uuid.UUID, productType domain.ProductType) (domain.ProductType, error) {
	err := repository.DB.Model(&productType).Where("id = ?", id).Update("name", productType.Name).Error
	if err != nil {
		return domain.ProductType{}, err
	}

	return productType, err
}

func (repository *ProductTypeRepositoryImpl) Delete(productType domain.ProductType) error {
	existingProductType, err := repository.FindById(productType.ID)
	if err != nil {
		return err
	}

	err = repository.DB.Delete(&existingProductType).Error
	if err != nil {
		return err
	}
	return nil
}
