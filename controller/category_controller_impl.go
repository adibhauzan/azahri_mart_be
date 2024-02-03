package controller

import (
	"github.com/google/uuid"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	productCategoryResponse, err := controller.ProductCategoryService.Create(productCategoryRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "message": "Create Product Category success", "data": productCategoryResponse})

}

func (controller *ProductCategoryControllerImpl) Update(ctx *gin.Context) {
	var productCategoryRequest request.UpdateProductCategoryRequest
	if err := ctx.ShouldBindJSON(&productCategoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idString := ctx.Params.ByName("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productCategoryRequest.ID = id
	productCategoryResponse, err := controller.ProductCategoryService.Update(productCategoryRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	
	ctx.JSON(200, gin.H{"code": 200, "message": "Update Product Category success", "data": productCategoryResponse})
}
func (controller *ProductCategoryControllerImpl) FindById(ctx *gin.Context) {
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

	productCategoryResponse, err := controller.ProductCategoryService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found Id"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": productCategoryResponse})

}

func (controller *ProductCategoryControllerImpl) FindAll(ctx *gin.Context) {
	categoryResponses, err := controller.ProductCategoryService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	productCategoryResponse := categoryResponses
	ctx.JSON(200, gin.H{"code": 200, "data": productCategoryResponse})

}

func (controller *ProductCategoryControllerImpl) Delete(ctx *gin.Context) {
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

	err = controller.ProductCategoryService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "Product Category Delete Success"})
}
