package main

import (
	"github.com/adibhauzan/azahri_mart_be/app"
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/model/domain"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

func main() {

	db := app.NewDbConnection()
	db.AutoMigrate(&domain.Category{})

	categoryRepo := repositories.NewCategoryRepositoryImpl(db)
	categoryService := service.NewCategoryService(categoryRepo, db)
	CategoryController := controller.NewCategoryController(categoryService)

	router := gin.Default()

	router.POST("/category", CategoryController.Create)

	router.Run(":3000")

}
