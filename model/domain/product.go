package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	CategoryID uuid.UUID `gorm:"column:category_id;type:uuid"`
	TypeID     uuid.UUID `gorm:"column:type_id;type:uuid"`
	Name       string    `gorm:"column:name"`
	Stock      int       `gorm:"column:stock"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	Category ProductCategory `gorm:"foreignKey:CategoryID;references:ID"`
	Type     ProductType     `gorm:"foreignKey:TypeID;references:ID"`
}
