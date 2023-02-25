package gudang

import (
	"gorm.io/gorm"
)

type Gudang struct {
	GudangId string `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Status   bool   `json:"status" gorm:"type:boolean:column:status"`
	Rack     string `json:"rack"`
}

type Rack struct {
	RackId   string `json:"rack_id"`
	RackName string `json:"rack_name"`
	GudangId string `json:"gudang_id"`
}

type GudangRepository interface {
	FindAll() ([]Gudang, error)
	Create(Data Gudang) error
	Update(Id string, Data Gudang) error
	Delete(Id string) error
	CreateRack(Data []Rack) error
	DeleteRack(Id string) error
	GetRack(GudangId string) ([]Rack, error)
}

type gudangRepository struct {
	db *gorm.DB
}

func NewGudangRepository(db *gorm.DB) *gudangRepository {
	return &gudangRepository{db: db}
}

func (r *gudangRepository) FindAll() ([]Gudang, error) {
	var gudang []Gudang

	err := r.db.Raw(`select tg.*, (
		select jsonb_agg(ra)
		from (
			select r.rack_id, r.rack_name from rack as r where r.gudang_id = tg.id
		) as ra
	   )as rack from tbm_gudang as tg`).Scan(&gudang).Error

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

//rack

func (r *gudangRepository) CreateRack(Data []Rack) error {
	err := r.db.Table("rack").Create(&Data).Error

	return err
}

func (r *gudangRepository) DeleteRack(Id string) error {
	err := r.db.Exec("DELETE FROM rack WHERE rack_id = ?", Id).Error

	return err
}

func (r *gudangRepository) GetRack(GudangId string) ([]Rack, error) {

	var rack []Rack
	err := r.db.Table("rack").Where("gudang_id = ?", GudangId).Find(&rack).Error

	return rack, err
}
