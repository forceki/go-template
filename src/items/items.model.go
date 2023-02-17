package items

import "time"

type Items struct {
	ItemId     string    `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama       string    `json:"nama"`
	Barcode    string    `json:"barcode"`
	KategoriId int       `json:"kategori_id"`
	Satuan     string    `json:"satuan"`
	Detail     string    `json:"detail"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	CreatedBy  string    `json:"created_by,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	UpdatedBy  string    `json:"updated_by,omitempty"`
}

type ItemsRequest struct {
	Nama       string `json:"nama" validate:"required"`
	Barcode    string `json:"barcode" validate:"required"`
	KategoriId int    `json:"kategori_id" validate:"required"`
	Satuan     string `json:"satuan" validate:"required"`
	Detail     string `json:"detail"`
}

type ItemsResponse struct {
	ItemId       string `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama         string `json:"nama"`
	Barcode      string `json:"barcode"`
	KategoriNama string `json:"kategori_nama"`
	KategoriId   int    `json:"kategori_id"`
	Satuan       string `json:"satuan"`
	Detail       string `json:"detail"`
}
