package items

type ItemsService interface {
	FindAll(Id string) (*[]ItemsResponse, error)
	Create(Data ItemsRequest) error
	Update(Id string, Data ItemsRequest) error
	Delete(Id string) error
}

type itemsService struct {
	itemsRepository ItemsRepository
}

func NewItemsService(itemsRepository ItemsRepository) *itemsService {
	return &itemsService{itemsRepository}
}

func (s *itemsService) FindAll(Id string) (*[]ItemsResponse, error) {
	data, err := s.itemsRepository.FindAll(Id)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *itemsService) Create(Data ItemsRequest) error {
	data := Items{
		Nama:       Data.Nama,
		Barcode:    Data.Barcode,
		KategoriId: Data.KategoriId,
		Satuan:     Data.Satuan,
		Detail:     Data.Detail,
	}

	err := s.itemsRepository.Create(data)

	return err
}

func (s *itemsService) Update(Id string, Data ItemsRequest) error {

	data := Items{
		Nama:       Data.Nama,
		Barcode:    Data.Barcode,
		KategoriId: Data.KategoriId,
		Satuan:     Data.Satuan,
		Detail:     Data.Detail,
	}

	err := s.itemsRepository.Update(Id, data)

	return err
}

func (s *itemsService) Delete(Id string) error {
	err := s.itemsRepository.Delete(Id)

	return err
}
