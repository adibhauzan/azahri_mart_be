package controller

import "github.com/gin-gonic/gin"

type ProductCategoryController interface {
	Create(ctx *gin.Context)
}
