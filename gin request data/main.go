package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin Request data Demo")
	r := gin.Default()

	// Query data using GET method

	// Sample query
	// http://localhost:3000/query/?username=don&month=april&month=may&year=1999

	r.GET("/query/*rest", func(ctx *gin.Context) {
		username := ctx.Query("username")
		year := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := ctx.QueryArray("month")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	// Post form data handling

	r.GET("/employee", func(ctx *gin.Context) {
		ctx.File("./public/employee.html")
	})

	r.POST("employee", func(ctx *gin.Context) {
		date := ctx.PostForm("date")
		amount := ctx.PostForm("amount")
		username := ctx.DefaultPostForm("username", "tom")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		})
	})

	log.Fatal(r.Run(":3000"))
}
