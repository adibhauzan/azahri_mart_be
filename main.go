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
	err := db.AutoMigrate(&domain.ProductCategory{}, &domain.ProductType{}, domain.Product{})
	if err != nil {
		panic(err.Error())
	}

	productCategoryRepo := repositories.NewCategoryRepositoryImpl(db)
	productCategoryService := service.NewCategoryService(productCategoryRepo, db)
	productCategoryController := controller.NewCategoryController(productCategoryService)

	productTypeRepository := repositories.NewProductTypeRepository(db)
	productTypeService := service.NewProductTypeService(productTypeRepository, db)
	productTypeController := controller.NewProductTypeController(productTypeService)

	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepository, db)
	productController := controller.NewProductController(productService)

	router := gin.Default()

	api := router.Group("/api")

	productCategoryRoutes := api.Group("/category")

	productCategoryRoutes.POST("/", productCategoryController.Create)
	productCategoryRoutes.PUT("/:id", productCategoryController.Update)
	productCategoryRoutes.DELETE("/:id", productCategoryController.Delete)
	productCategoryRoutes.GET("/:id", productCategoryController.FindById)
	productCategoryRoutes.GET("/", productCategoryController.FindAll)

	productTypeRoutes := api.Group("/type")
	productTypeRoutes.POST("/", productTypeController.Create)
	productTypeRoutes.PUT("/:id", productTypeController.Update)
	productTypeRoutes.DELETE("/:id", productTypeController.Delete)
	productTypeRoutes.GET("/:id", productTypeController.FindById)
	productTypeRoutes.GET("/", productTypeController.FindAll)

	productRoutes := api.Group("/product")
	productRoutes.POST("/", productController.Create)
	productRoutes.PUT("/:id", productController.Update)
	productRoutes.DELETE("/:id", productController.Delete)
	productRoutes.GET("/:id", productController.FindById)
	productRoutes.GET("/", productController.FindAll)

	err = router.Run(":3000")
	if err != nil {
		panic(err)
	}

}
