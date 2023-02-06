package items

import (
	"gorm.io/gorm"
)

type ItemsRepository interface {
	FindAll(Id string) ([]ItemsResponse, error)
	Create(Data Items) error
	Update(Id string, Data Items) error
	Delete(Id string) error
}

type itemsRepository struct {
	db *gorm.DB
}

func NewItemsRepository(db *gorm.DB) *itemsRepository {
	return &itemsRepository{db: db}
}

func (r *itemsRepository) FindAll(Id string) ([]ItemsResponse, error) {
	var item []ItemsResponse

	if Id != "" {
		err := r.db.Table("tbm_items as ti").Select("ti.id", "ti.nama", "ti.barcode", "ti.satuan", "ti.detail", "tk.nama as kategori_nama", "tk.id as kategori_id").Joins("LEFT JOIN tbm_kategori as tk on tk.id = ti.kategori_id").Where("ti.id", Id).Order("ti.id DESC").Find(&item).Error
		return item, err
	} else {
		err := r.db.Table("tbm_items as ti").Select("ti.id", "ti.nama", "ti.barcode", "ti.satuan", "ti.detail", "tk.nama as kategori_nama", "tk.id as kategori_id").Joins("LEFT JOIN tbm_kategori as tk on tk.id = ti.kategori_id").Find(&item).Error
		return item, err
	}

}

func (r *itemsRepository) Create(Data Items) error {
	item := Data

	err := r.db.Table("tbm_items").Create(&item).Error

	return err
}

func (r *itemsRepository) Update(Id string, Data Items) error {

	err := r.db.Table("tbm_items").Where("id  = ? ", Id).Updates(Data).Error

	return err
}
func (r *itemsRepository) Delete(Id string) error {
	err := r.db.Exec("DELETE FROM tbm_items WHERE id = ?", Id).Error

	return err
}
