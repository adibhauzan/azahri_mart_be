package Routes

import (
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoutes(router *gin.Engine, db *gorm.DB) {
	productRepository := repositories.NewProductRepository(db)
	productCategoryRepo := repositories.NewCategoryRepositoryImpl(db)
	productTypeRepository := repositories.NewProductTypeRepository(db)

	productService := service.NewProductService(
		productRepository,
		productCategoryRepo,
		productTypeRepository,
		db,
	)
	productController := controller.NewProductController(productService)
	productRoutes := router.Group("/product")
	productRoutes.Use()
	productRoutes.POST("/", productController.Create)
	productRoutes.PUT("/:id", productController.Update)
	productRoutes.DELETE("/:id", productController.Delete)
	productRoutes.GET("/:id", productController.FindById)
	productRoutes.GET("/category_id/:category_id", productController.FindByCategoryId)
	productRoutes.GET("/type_id/:type_id", productController.FindByTypeId)
	productRoutes.GET("/", productController.FindAll)

}
