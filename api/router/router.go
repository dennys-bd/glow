package router

import (
	"github.com/gin-gonic/gin"
)

func SetApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	makeClassHandlers(api)
	makeBookingHandlers(api)
}
