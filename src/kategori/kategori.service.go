package kategori

type KategoriRes struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type KategoriService interface {
	FindAll() (*[]Kategori, error)
	Create(Data Kategori) error
	Update(Id string, Data Kategori) error
	Delete(Id string) error
}

type kategoriService struct {
	kategoriRepository KategoriRepository
}

func NewKategoriService(kategoriRepository KategoriRepository) *kategoriService {
	return &kategoriService{kategoriRepository}
}

func (s *kategoriService) FindAll() (*[]Kategori, error) {
	data, err := s.kategoriRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *kategoriService) Create(Data Kategori) error {
	err := s.kategoriRepository.Create(Data)

	return err
}

func (s *kategoriService) Update(Id string, Data Kategori) error {
	err := s.kategoriRepository.Update(Id, Data)

	return err
}

func (s *kategoriService) Delete(Id string) error {
	err := s.kategoriRepository.Delete(Id)

	return err
}
