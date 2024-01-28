package service

import (
	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *gorm.DB
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, db *gorm.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepo,
		DB:                 db,
	}
}

func (service *CategoryServiceImpl) Create(request request.CreateCategoryRequest) response.CreateCategoryResponse {

	category := domain.Category{
		Name: request.Name,
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		category = service.CategoryRepository.Save(category)
		return nil
	})

	if err != nil {
		panic(err)
	}
	return helper.ToCreateCategoryResponse(category)
}
