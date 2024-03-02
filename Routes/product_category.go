package Routes

import (
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	productCategoryRepo := repositories.NewCategoryRepositoryImpl(db)
	productCategoryService := service.NewCategoryService(
		productCategoryRepo,
		db,
	)
	productCategoryController := controller.NewCategoryController(productCategoryService)

	productCategoryRoutes := router.Group("/category")
	productCategoryRoutes.POST("/", productCategoryController.Create)
	productCategoryRoutes.PUT("/:id", productCategoryController.Update)
	productCategoryRoutes.DELETE("/:id", productCategoryController.Delete)
	productCategoryRoutes.GET("/:id", productCategoryController.FindById)
	productCategoryRoutes.GET("/", productCategoryController.FindAll)
}
