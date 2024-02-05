package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

type UserService interface {
	Create(request request.CreateUserRequest) (response.UserResponse, error)
	/*

		for Update and Delete methods later after implementing jwt authorization
		
		/ this method for user with role admin
			FindById(userId uuid.UUID) (response.UserResponse, error)
			FindAll() ([]response.UserResponse, error)
		/



	*/
	
}