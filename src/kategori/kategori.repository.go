package kategori

import "gorm.io/gorm"

type Kategori struct {
	Id   string `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama string `json:"nama"`
}

type KategoriRepository interface {
	FindAll() ([]Kategori, error)
	Create(Data Kategori) error
	Update(Id string, Data Kategori) error
	Delete(Id string) error
}

type kategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) *kategoriRepository {
	return &kategoriRepository{db: db}
}

func (r *kategoriRepository) FindAll() ([]Kategori, error) {
	var data []Kategori

	err := r.db.Table("tbm_kategori").Select("id", "nama").Order("id DESC").Find(&data).Error

	return data, err
}

func (r *kategoriRepository) Create(Data Kategori) error {
	data := Data

	err := r.db.Table("tbm_kategori").Create(&data).Error

	return err
}

func (r *kategoriRepository) Update(Id string, Data Kategori) error {
	data := Data

	err := r.db.Table("tbm_kategori").Where("id = ?", Id).Updates(&data).Error

	return err
}

func (r *kategoriRepository) Delete(Id string) error {
	err := r.db.Exec("DELETE FROM tbm_kategori WHERE id = ?", Id).Error

	return err
}
