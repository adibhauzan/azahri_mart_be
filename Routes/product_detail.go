package Routes

import (
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductDetailRoutes(router *gin.Engine, db *gorm.DB) {

	productDetailRepository := repositories.NewProductDetailRepository(db)
	productDetailService := service.NewProductDetailService(db,
		productDetailRepository)
	productDetailController := controller.NewProductDetailController(productDetailService)

	productDetailRoutes := router.Group("/product-detail", productDetailController.Create)
	productDetailRoutes.POST("/")
	productDetailRoutes.PUT("/:id", productDetailController.Update)
	productDetailRoutes.DELETE("/:id", productDetailController.Delete)
	productDetailRoutes.GET("/:id", productDetailController.FindById)
	// productDetailRoutes.GET("/product_id/:product_id", productDetailController.FindByProductId)
	productDetailRoutes.GET("/", productDetailController.FindAll)
}
