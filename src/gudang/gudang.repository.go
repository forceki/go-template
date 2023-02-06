package gudang

import (
	"gorm.io/gorm"
)

type Gudang struct {
	GudangId int    `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Status   bool   `json:"status" gorm:"type:boolean:column:status"`
}

type GudangRepository interface {
	FindAll() ([]Gudang, error)
	Create(Data Gudang) error
	Update(Id string, Data Gudang) error
	Delete(Id string) error
}

type gudangRepository struct {
	db *gorm.DB
}

func NewGudangRepository(db *gorm.DB) *gudangRepository {
	return &gudangRepository{db: db}
}

func (r *gudangRepository) FindAll() ([]Gudang, error) {
	var gudang []Gudang

	err := r.db.Table("tbm_gudang").Select("id", "nama", "alamat", "status").Order("id DESC").Find(&gudang).Error

	return gudang, err
}

func (r *gudangRepository) Create(Data Gudang) error {
	gudang := Data

	err := r.db.Table("tbm_gudang").Create(&gudang).Error

	return err
}

func (r *gudangRepository) Update(Id string, Data Gudang) error {
	gudang := Data

	err := r.db.Table("tbm_gudang").Where("id  = ? ", Id).Updates(map[string]interface{}{"nama": gudang.Nama, "alamat": gudang.Alamat, "status": gudang.Status}).Error

	return err
}
func (r *gudangRepository) Delete(Id string) error {
	err := r.db.Exec("DELETE FROM tbm_gudang WHERE id = ?", Id).Error

	return err
}
