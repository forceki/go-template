package checkin

import "time"

type CheckinDetail struct {
	Id         int    `json:"id" gorm:"column:id; PRIMARY_KEY"`
	CheckinsId string `json:"checkins_id"`
	ItemId     string `json:"item_id"`
	Qty        int    `json:"qty"`
}

type Checkin struct {
	Id         string    `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Code       string    `json:"code"`
	Total      int       `json:"total"`
	SupplierId string    `json:"supplier_id"`
	GudangId   string    `json:"gudang_id"`
	Tanggal    time.Time `json:"taggal"`
	Keterangan string    `json:"keterangan"`
	RackId     string    `json:"rack_id"`
	Status     int       `json:"status"`
}
type CheckinResponse struct {
	Id         string    `json:"id,omitempty" gorm:"column:id; PRIMARY_KEY"`
	Code       string    `json:"code,omitempty"`
	Total      int       `json:"total,omitempty"`
	Supplier   string    `json:"supplier,omitempty"`
	SupplierId string    `json:"supplier_id,omitempty"`
	GudangId   string    `json:"gudang_id,omitempty"`
	Gudang     string    `json:"gudang,omitempty"`
	Tanggal    time.Time `json:"tanggal"`
	Keterangan string    `json:"keterangan"`
	RackId     string    `json:"rack_id"`
	RackName   string    `json:"rack_name"`
	Status     int       `json:"status"`
}

type CheckinDetailResponse struct {
	Id     int    `json:"checkins_id" gorm:"column:id; PRIMARY_KEY"`
	ItemId string `json:"id"`
	Nama   string `json:"nama"`
	Qty    int    `json:"qty"`
}

type CheckinDetailRes struct {
	Id  string `json:"id" validate:"required"`
	Qty int    `json:"qty" validate:"required"`
}

type CheckinRes struct {
	Id         string             `json:"id"`
	Code       string             `json:"code"`
	Total      int                `json:"total" validate:"required"`
	SupplierId string             `json:"supplier_id" validate:"required"`
	GudangId   string             `json:"gudang_id" validate:"required"`
	Tanggal    time.Time          `json:"tanggal"`
	Keterangan string             `json:"keterangan"`
	RackId     string             `json:"rack_id"`
	Status     int                `json:"status"`
	Details    []CheckinDetailRes `json:"details" validate:"required"`
}
