package routers

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine)  {
	products(r)
	errors(r)
	users(r)
}
