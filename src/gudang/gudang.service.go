package gudang

type gudangRequest struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Status bool   `json:"status"`
}

type GudangService interface {
	FindAll() (*[]Gudang, error)
	Create(Data gudangRequest) error
	Update(Id string, Data gudangRequest) error
	Delete(Id string) error
}

type gudangService struct {
	gudangRepository GudangRepository
}

func NewGudangService(gudangRepository GudangRepository) *gudangService {
	return &gudangService{gudangRepository}
}

func (s *gudangService) FindAll() (*[]Gudang, error) {
	data, err := s.gudangRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *gudangService) Create(Data gudangRequest) error {
	data := Gudang{
		Nama:   Data.Nama,
		Alamat: Data.Alamat,
		Status: Data.Status,
	}

	err := s.gudangRepository.Create(data)

	return err
}

func (s *gudangService) Update(Id string, Data gudangRequest) error {
	data := Gudang{
		Nama:   Data.Nama,
		Alamat: Data.Alamat,
		Status: Data.Status,
	}

	err := s.gudangRepository.Update(Id, data)

	return err
}

func (s *gudangService) Delete(Id string) error {
	err := s.gudangRepository.Delete(Id)

	return err
}
