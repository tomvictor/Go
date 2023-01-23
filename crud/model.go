package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
