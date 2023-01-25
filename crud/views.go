package main

import "gorm.io/gorm"

func GetAllProducts(db *gorm.DB) []Product {
	var products []Product
	db.Find(&products)
	return products
}

func CreateProduct(db *gorm.DB, code string, price uint) {
	db.Create(&Product{Code: code, Price: price})
}
