package repositories

import "github.com/adibhauzan/azahri_mart_be/model/domain"

type CategoryRepository interface {
	Save(category domain.Category) domain.Category
}
