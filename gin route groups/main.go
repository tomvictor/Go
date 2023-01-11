package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin Route Groups")
	router := gin.Default()
	adminGroup := router.Group("/admin")

	adminGroup.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user route")
	})

	adminGroup.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "role route")
	})

	adminGroup.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "policies route")
	})

	log.Fatal(router.Run(":3000"))
}
