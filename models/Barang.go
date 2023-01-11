package models

import "time"

type Barang struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaBarang string `gorm:"type:varchar(255)" json:"nama_barang"`
	Harga      int    `gorm:"type:integer(255)" json:"harga"`
	Deskripsi  string `gorm:"type:varchar(255)" json:"deskripsi"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
