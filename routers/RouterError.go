package routers

import (
	"github.com/gin-gonic/gin"
)

func errors(r *gin.Engine) {
	r.NoRoute(func(context *gin.Context) {
		context.String(404,"Not router")
	})
}
