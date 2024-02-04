package request

import (
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name       string    `json:"name" form:"name" binding:"required,min=3,max=99"`
	CategoryID uuid.UUID `json:"category_id" form:"category_id" binding:"required"`
	TypeID     uuid.UUID `json:"type_id" form:"type_id" binding:"required"`
}

type UpdateProductRequest struct {
	ID         uuid.UUID `json:"id" uri:"id"`
	Name       string    `json:"name" form:"name" binding:"required,min=3,max=99"`
	CategoryID uuid.UUID `json:"category_id" form:"category_id" binding:"required"`
	TypeID     uuid.UUID `json:"type_id" form:"type_id" binding:"required"`
}
