package repositories

import "github.com/adibhauzan/azahri_mart_be/model/domain"

type ProductTypeRepository interface {
	Save(productType domain.ProductType) domain.ProductType
}
