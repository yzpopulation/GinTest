package routers

import (
	"GinTest/controllers"
	"github.com/gin-gonic/gin"
)

func products(r *gin.Engine) {
	r.GET("/products", controllers.FindProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.POST("/products", controllers.CreateProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
