package main

import (
	"github.com/adibhauzan/azahri_mart_be/Routes"
	"github.com/adibhauzan/azahri_mart_be/app"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

func main() {

	db := app.NewDbConnection()

	router := gin.Default()

	router.Group("/api")

	Routes.NewProductTypeRoutes(router, db)

	Routes.NewProductCategoryRoutes(router, db)

	Routes.NewProductRoutes(router, db)

	Routes.NewProductDetailRoutes(router, db)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
