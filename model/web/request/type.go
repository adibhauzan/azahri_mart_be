package request

import "github.com/google/uuid"

type CreateProductTypeRequest struct {
	Name string `json:"name" binding:"required,min=3,max=99"`
}

type UpdateProductTypeRequest struct {
	ID   uuid.UUID `json:"id" uri:"id"`
	Name string    `json:"name" binding:"required,min=3,max=99"`
}