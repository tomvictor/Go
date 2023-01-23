package main

import "gorm.io/gorm"

func GetAllProducts(db *gorm.DB) []Product {
	var products []Product
	db.Find(&products)
	return products
}

func CreateProduct(db *gorm.DB, code string, price int) {
	db.Create(&Product{Code: code, Price: uint(price)})
}
