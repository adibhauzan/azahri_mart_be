package controller

import (
	"net/http"

	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService *service.UserServiceImpl
}

func NewUserController(userService *service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(ctx *gin.Context) {
	var userRequest request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := controller.UserService.Create(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "Create User success", "data": userResponse})
}

func (controller *UserControllerImpl) Update(ctx *gin.Context) {
	panic("asdasd")
}

func (controller *UserControllerImpl) FindById(ctx *gin.Context) {
	panic("asdasd")
}

func (controller *UserControllerImpl) Delete(ctx *gin.Context) {
	panic("asdasd")
}

func (controller *UserControllerImpl) FindAll(ctx *gin.Context) {
	panic("asdasd")
}
