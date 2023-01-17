package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{
		"tom": "asd1",
		"don": "asd2",
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello",
		})
	})

	log.Fatal(router.Run(":3000"))
}
