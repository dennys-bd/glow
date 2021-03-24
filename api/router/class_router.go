package router

import (
	"github.com/dennys-bd/glow/wire"
	"github.com/gin-gonic/gin"
)

func makeClassHandlers(r *gin.RouterGroup) {
	controller, err := wire.InjectClassController()
	if err != nil {
		panic(err)
	}

	r.GET("/classes", controller.Index)
	r.GET("/classes/:id", controller.Show)
	r.POST("/classes", controller.Create)
}
