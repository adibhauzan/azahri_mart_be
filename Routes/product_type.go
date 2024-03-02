package Routes

import (
	"github.com/adibhauzan/azahri_mart_be/controller"
	"github.com/adibhauzan/azahri_mart_be/repositories"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductTypeRoutes(router *gin.Engine, db *gorm.DB) {
	productTypeRepository := repositories.NewProductTypeRepository(db)
	productTypeService := service.NewProductTypeService(
		productTypeRepository,
		db,
	)
	productTypeController := controller.NewProductTypeController(productTypeService)

	productTypeRoutes := router.Group("/type")
	productTypeRoutes.POST("/", productTypeController.Create)
	productTypeRoutes.PUT("/:id", productTypeController.Update)
	productTypeRoutes.DELETE("/:id", productTypeController.Delete)
	productTypeRoutes.GET("/:id", productTypeController.FindById)
	productTypeRoutes.GET("/", productTypeController.FindAll)
}
