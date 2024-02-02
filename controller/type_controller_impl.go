package controller

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	productTypeResponse, err := controller.ProductTypeService.Create(productTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"code": 201, "message": "Create Product Type success", "data": productTypeResponse})
}

func (controller *ProductTypeControllerImpl) Update(ctx *gin.Context) {
	var productTypeRequest request.UpdateProductTypeRequest
	if err := ctx.ShouldBindJSON(&productTypeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idString := ctx.Params.ByName("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productTypeRequest.ID = id
	productTypeResponse, err := controller.ProductTypeService.Update(productTypeRequest)
	ctx.JSON(200, gin.H{"code": 200, "message": "Update Product Type success", "data": productTypeResponse})
}

func (controller *ProductTypeControllerImpl) FindById(ctx *gin.Context) {
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

	productTypeResponse, err := controller.ProductTypeService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found Id"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": productTypeResponse})
}

func (controller *ProductTypeControllerImpl) FindAll(ctx *gin.Context) {
	productTypeResponses, err := controller.ProductTypeService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productTypeResponse := productTypeResponses
	ctx.JSON(200, gin.H{"code": 200, "data": productTypeResponse})
}

func (controller *ProductTypeControllerImpl) Delete(ctx *gin.Context) {
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

	err = controller.ProductTypeService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Notfound"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "Product Type Delete Success"})
}