package Routes

import (
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := service.NewUserService(db, userRepository)
	userController := controller.NewUserController(userService)
	userRoutes := router.Group("/user")
	userRoutes.POST("/register", userController.Create)
}
