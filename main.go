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
	db.AutoMigrate(&domain.ProductCategory{}, &domain.ProductType{})

	productCategoryRepo := repositories.NewCategoryRepositoryImpl(db)
	productCategoryService := service.NewCategoryService(productCategoryRepo, db)
	productCategoryController := controller.NewCategoryController(productCategoryService)

	productTypeRepository := repositories.NewProductTypeRepository(db)
	productTypeService := service.NewProductTypeService(productTypeRepository, db)
	productTypeController := controller.NewProductTypeController(productTypeService)

	router := gin.Default()

	api := router.Group("/api")

	productCategoryRoutes := api.Group("/category")
	productTypeRoutes := api.Group("/type")

	productCategoryRoutes.POST("/", productCategoryController.Create)
	productTypeRoutes.POST("/", productTypeController.Create)

	router.Run(":3000")

}
