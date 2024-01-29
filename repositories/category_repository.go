package repositories

import "github.com/adibhauzan/azahri_mart_be/model/domain"

type ProductCategoryRepository interface {
	Save(category domain.ProductCategory) domain.ProductCategory
}
