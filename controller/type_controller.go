package controller

import "github.com/gin-gonic/gin"

type ProductTypeController interface {
	Create(ctx *gin.Context)
}
