package controller

import "github.com/gin-gonic/gin"

type ProductCategoryController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductTypeController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByCategoryId(ctx *gin.Context)
	FindByTypeId(ctx *gin.Context)
}

type ProductDetailController interface {
	Create(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// FindById(ctx *gin.Context)
	// FindAll(ctx *gin.Context)
	// Delete(ctx *gin.Context)
	// FindByProductId(ctx *gin.Context)
}
