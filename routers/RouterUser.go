package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func users(r *gin.Engine) {
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK,gin.H{"name":name})
	})
}
