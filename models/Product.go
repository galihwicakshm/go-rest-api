package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NamaProduct string
	Harga       int
	Deskripsi   string
}
