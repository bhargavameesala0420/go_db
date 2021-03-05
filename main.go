package main

import (
	"log"

	routes "github.com/bhargavameesala0420/go_db/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
