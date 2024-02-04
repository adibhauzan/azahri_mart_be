package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductDetail struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"column:product_id;type:uuid"`
	ExpiredAt time.Time `gorm:"column:expired_at"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}
