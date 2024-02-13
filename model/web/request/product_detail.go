package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductDetailRequest struct {
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at"`
}

type UpdateProductDetailRequest struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at" binding:"required"`
}
