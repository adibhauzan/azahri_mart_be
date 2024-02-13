package controller

import (
	"net/http"

	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/adibhauzan/azahri_mart_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (controller *ProductDetailControllerImpl) Update(ctx *gin.Context) {
	var productDetailRequest request.UpdateProductDetailRequest
	if err := ctx.ShouldBindJSON(&productDetailRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idString := ctx.Params.ByName("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengonversi 'id' ke UUID"})
		return
	}

	productDetailRequest.ID = id
	productTypeResponse, err := controller.ProductDetailService.Update(productDetailRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "message": "Update Product Detail success", "data": productTypeResponse})
}

func (controller *ProductDetailControllerImpl) FindById(ctx *gin.Context) {
	idString, b := ctx.Params.Get("id")
	if !b {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'id' parameter"})
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengonversi 'id' ke UUID"})
		return
	}

	productDetailResponse, err := controller.ProductDetailService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": productDetailResponse})
}

func (controller *ProductDetailControllerImpl) FindAll(ctx *gin.Context) {
	productDetailResponses, err := controller.ProductDetailService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	productDetailResponse := productDetailResponses
	ctx.JSON(200, gin.H{"code": 200, "data": productDetailResponse})
}

func (controller *ProductDetailControllerImpl) Delete(ctx *gin.Context) {
	idString, b := ctx.Params.Get("id")
	if !b {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'id' parameter"})
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse 'id' to UUID"})
		return
	}

	err = controller.ProductDetailService.Delete(id)
	if err != nil {
		if err == utils.ErrNotFound{
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}


func (controller *ProductDetailControllerImpl) FindByProductId(ctx *gin.Context) {
	productID, err := uuid.Parse(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse 'category_id' to UUID"})
		return
	}

	products, err := controller.ProductDetailService.FindByProductId(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Product Details found by product ID", "data": products})
}
