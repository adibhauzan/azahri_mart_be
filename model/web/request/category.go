package request

import "github.com/google/uuid"

type CreateProductCategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=3,max=99"`
}

type UpdateProductCategoryRequest struct {
	ID       uuid.UUID `json:"id"`
	Name string    `json:"name" form:"id" binding:"required,min=3,max=99"`
}
