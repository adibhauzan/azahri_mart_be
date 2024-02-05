package repositories

import (
	"fmt"

	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {

	// userPassword, err := utils.HashPassword(user.Password)
	// if err != nil {
	// 	return domain.User{}, fmt.Errorf("failed to hash User's Password : %v", err)
	// }

	// dbUser := domain.User{
	// 	ID:       uuid.New(),
	// 	Role:     user.Role,
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// 	Password: userPassword,
	// }

	err := repository.DB.Create(&user).Error
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create User : %v", err)
	}
	return user, nil
}
