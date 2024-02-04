package main

import (
	"github.com/adibhauzan/azahri_mart_be/app"
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

func main() {

	db := app.NewDbConnection()

	// Product Category 
	productCategoryRepo := repositories.NewCategoryRepositoryImpl(db)
	productCategoryService := service.NewCategoryService(
		productCategoryRepo, 
		db,
	)
	productCategoryController := controller.NewCategoryController(productCategoryService)

	// End Product Category

	// Product Type
	productTypeRepository := repositories.NewProductTypeRepository(db)
	productTypeService := service.NewProductTypeService(
		productTypeRepository, 
		db,
	)
	productTypeController := controller.NewProductTypeController(productTypeService)
	// End Product Type 

	// Product
	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(
		productRepository,
		productCategoryRepo,
		productTypeRepository,
		db,
	)
	productController := controller.NewProductController(productService)
	// End Product

	// Product Detail
	productDetailRepository := repositories.NewProductDetailRepository(db)
	productDetailService := service.NewProductDetailService(db, 
		productDetailRepository, 
		productRepository,
	)
	productDetailController := controller.NewProductDetailController(productDetailService)
	// 
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
	productRoutes.GET("/category_id/:category_id", productController.FindByCategoryId)
	productRoutes.GET("/type_id/:type_id", productController.FindByTypeId)
	productRoutes.GET("/", productController.FindAll)



	productDetailRoutes := api.Group("/product-detail", productDetailController.Create)
	productDetailRoutes.POST("/", )

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}

}
