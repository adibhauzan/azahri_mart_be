package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{DB: db}
}

func (repository *ProductRepositoryImpl) Save(product domain.Product) domain.Product {
	err := repository.DB.Create(&product).Error
	if err != nil {
		return domain.Product{}
	}
	return product
}

func (repository *ProductRepositoryImpl) FindById(id uuid.UUID) (domain.Product, error) {
	var product domain.Product
	if err := repository.DB.First(&product, id).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	var product []domain.Product
	if err := repository.DB.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Update(id uuid.UUID, product domain.Product) (domain.Product, error) {
	err := repository.DB.Model(&product).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(product domain.Product) error {
	existingProduct, err := repository.FindById(product.ID)
	if err != nil {
		return err
	}

	err = repository.DB.Delete(&existingProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *ProductRepositoryImpl) FindByCategoryId(categoryID uuid.UUID) ([]domain.Product, error) {
	var products []domain.Product
	if err := repository.DB.Preload("Category").Preload("Type").Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *ProductRepositoryImpl) FindByTypeId(typeID uuid.UUID) ([]domain.Product, error) {
	var products []domain.Product
	if err := repository.DB.Preload("Category").Preload("Type").Where("type_id = ?", typeID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
