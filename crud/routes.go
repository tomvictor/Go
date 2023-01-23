package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(context *gin.Context) {
		products := GetAllProducts(db)
		context.JSONP(http.StatusOK, products)
	})

	r.GET("/create", func(context *gin.Context) {
		CreateProduct(db, "test", 12)
		context.JSONP(http.StatusOK, gin.H{
			"message": "Item created successfully",
		})
	})
}
