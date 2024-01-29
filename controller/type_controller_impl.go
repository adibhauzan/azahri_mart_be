package controller

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductTypeControllerImpl struct {
	ProductTypeService service.ProductTypeService
}

func NewProductTypeController(productTypeService service.ProductTypeService) ProductTypeController {
	return &ProductTypeControllerImpl{
		ProductTypeService: productTypeService,
	}
}

func (controller *ProductTypeControllerImpl) Create(ctx *gin.Context) {
	var productTypeRequest request.CreateProductTypeRequest

	if err := ctx.ShouldBindJSON(&productTypeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productTypeResponse := controller.ProductTypeService.Create(productTypeRequest)
	ctx.JSON(201, productTypeResponse)
}
