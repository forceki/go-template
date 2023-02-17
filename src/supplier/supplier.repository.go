package supplier

import "gorm.io/gorm"

type Supplier struct {
	ID     string `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Detail string `json:"detail"`
}

type SupplierRepository interface {
	FindAll() ([]Supplier, error)
	Create(Data Supplier) error
	Update(Id string, Data Supplier) error
	Delete(Id string) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *supplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) FindAll() ([]Supplier, error) {
	var supplier []Supplier

	err := r.db.Table("tbm_suppliers").Select("id", "nama", "email", "phone", "detail").Order("id DESC").Find(&supplier).Error

	return supplier, err
}

func (r *supplierRepository) Create(Data Supplier) error {
	supplier := Data

	err := r.db.Table("tbm_suppliers").Create(&supplier).Error

	return err
}

func (r *supplierRepository) Update(Id string, Data Supplier) error {
	supplier := Data

	err := r.db.Table("tbm_suppliers").Where("id  = ? ", Id).Updates(&supplier).Error

	return err
}

func (r *supplierRepository) Delete(Id string) error {
	err := r.db.Exec("DELETE FROM tbm_suppliers WHERE id = ?", Id).Error

	return err
}
