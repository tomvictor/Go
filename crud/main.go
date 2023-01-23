package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)

	r := gin.Default()

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

	log.Fatal(r.Run(":3000"))
}

func GetAllProducts(db *gorm.DB) []Product {
	var products []Product
	db.Find(&products)
	return products
}

func CreateProduct(db *gorm.DB, code string, price int) {
	db.Create(&Product{Code: code, Price: uint(price)})
}
