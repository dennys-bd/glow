package router

import (
	"github.com/dennys-bd/glow/wire"
	"github.com/gin-gonic/gin"
)

func makeBookingHandlers(r *gin.RouterGroup) {
	controller, err := wire.InjectBookingController()
	if err != nil {
		panic(err)
	}

	r.GET("/bookings", controller.Index)
	r.GET("/bookings/:id", controller.Show)
	r.POST("/bookings", controller.Create)
	r.DELETE("/bookings/:id", controller.Delete)
}
