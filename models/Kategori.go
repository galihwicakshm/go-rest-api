package models

import "time"

type Kategori struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaKategori string `gorm:"type:varchar(300)" json:"nama_kategori"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
