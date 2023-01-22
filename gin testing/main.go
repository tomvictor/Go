package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.GET("/api/1/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{"id": 1, "country": "India"})
	})
}

func main() {
	fmt.Println("Gin Testing Framework")
	r := gin.Default()
	registerRoutes(r)
	log.Fatal(r.Run(":3000"))
}
