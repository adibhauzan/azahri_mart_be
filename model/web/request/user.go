package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	Role     string `json:"role" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
}
