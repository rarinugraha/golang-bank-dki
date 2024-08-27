package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string
}

type Stock struct {
	ID             uint      `gorm:"primary_key"`
	NamaBarang     string    `json:"nama_barang"`
	JumlahStok     int       `json:"jumlah_stok"`
	NomorSeri      string    `json:"nomor_seri"`
	AdditionalInfo JSONB     `gorm:"type:jsonb" json:"additional_info"`
	Gambar         string    `json:"gambar"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      uint      `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      uint      `json:"updated_by"`
}
