package controller

import (
	"net/http"

	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
)

type ProductDetailControllerImpl struct {
	ProductDetailService service.ProductDetailService
}

func NewProductDetailController(productDetailService service.ProductDetailService) ProductDetailController {
	return &ProductDetailControllerImpl{
		ProductDetailService: productDetailService,
	}
}

func (controller *ProductDetailControllerImpl) Create(ctx *gin.Context) {
	var productDetailRequest request.CreateProductDetailRequest
	if err := ctx.ShouldBind(&productDetailRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	productDetailResponse, err := controller.ProductDetailService.Create(productDetailRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 201, "message": "Create Product Detail success", "data": productDetailResponse})
}
