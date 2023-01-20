package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var MyCustomMiddleware = func(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	timeTaken := time.Since(startTime)
	fmt.Printf("It took %v time to execute the request\n", timeTaken)
}

func main() {
	fmt.Println("Custom Gin Middleware")

	r := gin.Default()

	r.Use(MyCustomMiddleware)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"first_name": "Tom",
			"last_name":  "Victor",
		})
	})

	log.Fatal(r.Run(":3000"))
}
