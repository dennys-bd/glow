package main

import (
	"log"
	"os"

	"github.com/dennys-bd/glow/api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-gormigrate/gormigrate/v2"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	routes := gin.Default()

	// common := negroni.Classic()

	// webRouter := mux.NewRouter()
	router.SetApiRouter(routes)

	// webRouter.PathPrefix("/api").Handler(common.With(
	// 	middleware.NewCors(os.Getenv("ALLOWED_HOSTS")),
	// 	negroni.Wrap(apiRouter),
	// ))

	log.Printf("connect to http://localhost:%s/", port)
	// log.Fatal(http.ListenAndServe("localhost:"+port, webRouter))
	routes.Run(":" + port)
}
