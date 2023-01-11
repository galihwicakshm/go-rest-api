package models

import "time"

type Kategori struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaKategori string `gorm:"type:varchar(255)" json:"nama_product"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
