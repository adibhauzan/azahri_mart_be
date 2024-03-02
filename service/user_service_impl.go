package service

import (
	"errors"
	"time"

	"github.com/adibhauzan/azahri_mart_be/helper"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB             *gorm.DB
	UserRepository *repositories.UserRepositoryImpl
}

func NewUserService(db *gorm.DB, userRepo *repositories.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{
		DB:             db,
		UserRepository: userRepo,
	}
}

func (service *UserServiceImpl) Create(request request.CreateUserRequest) (response.UserResponse, error) {
	validRoles := []domain.Role{domain.Admin, domain.Gudang, domain.Cashier}

	userRole := domain.Role(request.Role)

	if !utils.IsValidRole(userRole, validRoles) {
		return response.UserResponse{}, errors.New("invalid role")
	}
	user := domain.User{
		ID:        uuid.New(),
		Role:      userRole,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := utils.HashPassword(user.Password)
	if err != nil {
		return response.UserResponse{}, errors.New("failed to hash password")
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		if !utils.IsValidRole(user.Role, validRoles) {
			return errors.New("invalid role")
		}
		_, err := service.UserRepository.Save(user)
		if err != nil {
			return errors.New("failed to save user: " + err.Error())
		}

		return nil
	})

	if err != nil {
		return response.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}
