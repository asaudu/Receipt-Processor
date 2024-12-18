package main

import (
	"addyCodes.com/ReceiptProcessor/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
