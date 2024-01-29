package controller

import (
	"net/http"

	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
)

type ProductCategoryControllerImpl struct {
	ProductCategoryService service.ProductCategoryService
}

func NewCategoryController(productCategoryService service.ProductCategoryService) ProductCategoryController {
	return &ProductCategoryControllerImpl{
		ProductCategoryService: productCategoryService,
	}
}

func (controller *ProductCategoryControllerImpl) Create(ctx *gin.Context) {

	var productCategoryRequest request.CreateProductCategoryRequest
	if err := ctx.ShouldBindJSON(&productCategoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCategoryResponse := controller.ProductCategoryService.Create(productCategoryRequest)
	ctx.JSON(201, productCategoryResponse)
}
