package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin   Role = "admin"
	Gudang  Role = "gudang"
	Cashier Role = "cashier"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	Role      Role      `gorm:"column:role"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
