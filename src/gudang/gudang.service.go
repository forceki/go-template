package gudang

import "github.com/google/uuid"

type gudangRequest struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Status bool   `json:"status"`
}

type RackRequest struct {
	GudangId string `json:"gudang_id"`
	Details  []struct {
		RackName string `json:"rack_name`
	}
}

type GudangService interface {
	FindAll() (*[]Gudang, error)
	Create(Data gudangRequest) error
	Update(Id string, Data gudangRequest) error
	Delete(Id string) error
	CreateRack(Data RackRequest) error
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
	id := uuid.NewString()
	data := Gudang{
		GudangId: id,
		Nama:     Data.Nama,
		Alamat:   Data.Alamat,
		Status:   Data.Status,
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

func (s *gudangService) CreateRack(Data RackRequest) error {
	var data []Rack

	for _, item := range Data.Details {
		ko := Rack{
			RackId:   uuid.NewString(),
			RackName: item.RackName,
			GudangId: Data.GudangId,
		}

		data = append(data, ko)
	}

	err := s.gudangRepository.CreateRack(data)

	return err
}
