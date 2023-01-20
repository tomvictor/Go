package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type TimeOffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"required,future" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

var ValidatorFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

func main() {
	fmt.Println("Gin Request data Demo")
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("future", ValidatorFuture)
	}
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

	// Normal method/Simple
	// r.POST("employee", func(ctx *gin.Context) {
	// 	date := ctx.PostForm("date")
	// 	amount := ctx.PostForm("amount")
	// 	username := ctx.DefaultPostForm("username", "tom")

	// 	ctx.IndentedJSON(http.StatusOK, gin.H{
	// 		"date":     date,
	// 		"amount":   amount,
	// 		"username": username,
	// 	})
	// })

	// Data Binding

	r.POST("employee", func(ctx *gin.Context) {
		var timeOffRequest TimeOffRequest

		if err := ctx.ShouldBind(&timeOffRequest); err == nil {
			ctx.JSON(http.StatusOK, timeOffRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	apiGroups := r.Group("/api")

	// Sample Payload
	// {
	// 	"date": "2023-12-01T00:00:00Z",
	// 	"amount": 5
	// }

	apiGroups.POST("timeoff", func(ctx *gin.Context) {
		var timeOffRequest TimeOffRequest

		if err := ctx.ShouldBind(&timeOffRequest); err == nil {
			ctx.JSON(http.StatusOK, timeOffRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	log.Fatal(r.Run(":3000"))
}
