package repositories

import (
	"github.com/adibhauzan/azahri_mart_be/model/domain"
)

type ProductDetailRepository interface {
	Save(productDetail domain.ProductDetail) (domain.ProductDetail, error)
}
