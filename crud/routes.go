package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

/*type Data struct {
	Code  string `json:"code" binding:"required"`
	Price string `json:"price" binding:"required"`
}*/

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(context *gin.Context) {
		products := GetAllProducts(db)
		context.JSON(http.StatusOK, products)
	})

	r.POST("/create", func(context *gin.Context) {
		//data := context.Request.Body
		var newProduct Product
		err := context.ShouldBindJSON(&newProduct)
		if err != nil {
			return
		}
		fmt.Println(newProduct)
		CreateProduct(db, newProduct.Code, newProduct.Price)
		context.JSON(http.StatusOK, gin.H{
			"message": "Item created successfully",
			"code":    newProduct.Code,
			"price":   newProduct.Price,
		})
	})
}
