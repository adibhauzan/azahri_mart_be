package controller

import "github.com/gin-gonic/gin"

type HandlerController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductCategoryController interface {
	HandlerController
}

type ProductTypeController interface {
	HandlerController
}

type ProductController interface {
	HandlerController
	FindByCategoryId(ctx *gin.Context)
	FindByTypeId(ctx *gin.Context)
}

type ProductDetailController interface {
	HandlerController
	FindByProductId(ctx *gin.Context)
}

type UserController interface {
	HandlerController
}
