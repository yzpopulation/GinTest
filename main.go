package main

import (
	"GinTest/models"
	"GinTest/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	var products []models.Product
	models.DB.Raw("select * from Products").Scan(&products)
	WebStart()
}

func WebStart() {
	r := gin.Default()
	routers.Init(r)
	//r.GET("/products", controllers.FindProduct)
	//r.GET("/products/:id", controllers.FindProduct)
	//r.POST("/products", controllers.CreateProduct)
	//r.PATCH("/products/:id", controllers.UpdateProduct)
	//r.DELETE("/products/:id", controllers.DeleteProduct)

	r.Run(":8080")
}
