package router

import (
	"github.com/gin-gonic/gin"
)

func SetApiRouter(r *gin.Engine) {
	// apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	// makeClassHandlers(apiRouter)
	// return apiRouter

	api := r.Group("/api")
	makeClassHandlers(api)
}
