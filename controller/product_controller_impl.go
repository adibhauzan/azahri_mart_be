package controller

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{ProductService: productService}
}

func (controller *ProductControllerImpl) Create(ctx *gin.Context) {
	var productRequest request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	productResponse, err := controller.ProductService.Create(productRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 201, "message": "Create Product success", "data": productResponse})

}

func (controller *ProductControllerImpl) Update(ctx *gin.Context) {
	var productRequest request.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idString := ctx.Params.ByName("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengonversi 'id' ke UUID"})
		return
	}

	productRequest.ID = id
	productTypeResponse, err := controller.ProductService.Update(productRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "message": "Update Product success", "data": productTypeResponse})
}

func (controller *ProductControllerImpl) FindById(ctx *gin.Context) {
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

	productResponse, err := controller.ProductService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": productResponse})
}

func (controller *ProductControllerImpl) FindAll(ctx *gin.Context) {
	productResponses, err := controller.ProductService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	productResponse := productResponses
	ctx.JSON(200, gin.H{"code": 200, "data": productResponse})
}

func (controller *ProductControllerImpl) Delete(ctx *gin.Context) {
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

	err = controller.ProductService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "Product Delete Success"})
}

func (controller *ProductControllerImpl) FindByCategoryId(ctx *gin.Context) {
	categoryID, err := uuid.Parse(ctx.Param("category_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse 'category_id' to UUID"})
		return
	}

	products, err := controller.ProductService.FindByCategoryId(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Products found by category ID", "data": products})

}
func (controller *ProductControllerImpl) FindByTypeId(ctx *gin.Context) {
	typeID, err := uuid.Parse(ctx.Param("type_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse 'type_id' to UUID"})
		return
	}

	products, err := controller.ProductService.FindByTypeId(typeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Products found by type ID", "data": products})
}
