package supplier

import "github.com/google/uuid"

type SupplierRequest struct {
	Nama   string `json:"nama" validate:"required"`
	Email  string `json:"email" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
	Detail string `json:"detail"`
}

type SupplierService interface {
	FindAll() (*[]Supplier, error)
	Create(Data SupplierRequest) error
	Update(Id string, Data SupplierRequest) error
	Delete(Id string) error
}

type supplierService struct {
	supplierRepository SupplierRepository
}

func NewSupplierService(supplierRepository SupplierRepository) *supplierService {
	return &supplierService{supplierRepository}
}

func (s *supplierService) FindAll() (*[]Supplier, error) {
	data, err := s.supplierRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (s *supplierService) Create(Data SupplierRequest) error {
	id := uuid.NewString()
	data := Supplier{
		ID:     id,
		Nama:   Data.Nama,
		Email:  Data.Email,
		Phone:  Data.Phone,
		Detail: Data.Detail,
	}

	err := s.supplierRepository.Create(data)

	return err
}

func (s *supplierService) Update(Id string, Data SupplierRequest) error {
	data := Supplier{
		Nama:   Data.Nama,
		Email:  Data.Email,
		Phone:  Data.Phone,
		Detail: Data.Detail,
	}

	err := s.supplierRepository.Update(Id, data)

	return err
}

func (s *supplierService) Delete(Id string) error {
	err := s.supplierRepository.Delete(Id)

	return err
}
