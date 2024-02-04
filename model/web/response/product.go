package response

import (
	"time"

	"github.com/google/uuid"
)

type ProductResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
	TypeID     uuid.UUID `json:"type_id"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
