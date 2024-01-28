package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository *CategoryRepositoryImpl) Save(category domain.Category) domain.Category {
	category.ID = uuid.New()
	err := repository.DB.Create(&category).Error
	if err != nil {
		panic(err)
	}
	return category
}
