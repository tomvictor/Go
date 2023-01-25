package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string `json:"code" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}
