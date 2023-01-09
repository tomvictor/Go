package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()

	router.StaticFile("/", "./public/index.html")
	router.Static("/static", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/index.html")
	})
	router.POST("employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New Request accepted")
	})

	log.Fatal(router.Run(":3000"))
}
