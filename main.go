package main

import (
	"log"
	"net/http"

	config "github.com/bhargavameesala0420/go_db/config"
	routes "github.com/bhargavameesala0420/go_db/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	// Connect DB
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4748"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(router)
	routes.Routes(router)
	// log.Fatal(router.Run(":4748"))
	log.Fatal(http.ListenAndServe(":4748", handler))
}
