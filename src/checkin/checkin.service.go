package checkin

import (
	"time"

	"github.com/google/uuid"
)

type CheckinService interface {
	Create(Data CheckinRes) error
	FindAll() (*[]CheckinResponse, error)
	Delete(Id string) error
	FindOne(Id string) (interface{}, error)
	Update(Id string, Data CheckinRes) error
}

type checkinService struct {
	checkinRepositroy CheckinRepositroy
}

func NewCheckinService(checkinRepositroy CheckinRepositroy) *checkinService {
	return &checkinService{checkinRepositroy}
}

func (s *checkinService) Create(Data CheckinRes) error {
	id := uuid.NewString()
	data := Checkin{
		Id:         id,
		Code:       Data.Code,
		Total:      Data.Total,
		SupplierId: Data.SupplierId,
		GudangId:   Data.GudangId,
		Keterangan: Data.Keterangan,
		Tanggal:    time.Now(),
	}

	var detail []CheckinDetail

	for _, item := range Data.Details {
		key := CheckinDetail{
			ItemId: item.Id,
			Qty:    item.Qty,
		}

		detail = append(detail, key)
	}

	err := s.checkinRepositroy.Create(data, detail)

	return err
}

func (s *checkinService) FindAll() (*[]CheckinResponse, error) {
	data, err := s.checkinRepositroy.FindAll()
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *checkinService) Delete(Id string) error {
	err := s.checkinRepositroy.Delete(Id)

	return err
}

func (s *checkinService) FindOne(Id string) (interface{}, error) {

	data, err := s.checkinRepositroy.FindOne(Id)

	if err != nil {
		return nil, err
	}

	item, err := s.checkinRepositroy.FindOneDetail(Id)

	if err != nil {
		return nil, err
	}

	type response struct {
		CheckinResponse
		Items []CheckinDetailResponse `json:"items"`
	}

	res := response{}

	res.Id = data.Id
	res.Code = data.Code
	res.Gudang = data.Gudang
	res.GudangId = data.GudangId
	res.Supplier = data.Supplier
	res.SupplierId = data.SupplierId
	res.Tanggal = data.Tanggal
	res.Keterangan = data.Keterangan
	res.Total = data.Total
	res.Items = item

	return res, nil
}

func (s *checkinService) Update(Id string, Data CheckinRes) error {
	data := Checkin{
		Id:         Data.Id,
		Code:       Data.Code,
		Total:      Data.Total,
		SupplierId: Data.SupplierId,
		GudangId:   Data.GudangId,
		Keterangan: Data.Keterangan,
		Tanggal:    Data.Tanggal,
	}

	var detail []CheckinDetail

	for _, item := range Data.Details {
		key := CheckinDetail{
			CheckinsId: data.Id,
			ItemId:     item.Id,
			Qty:        item.Qty,
		}

		detail = append(detail, key)
	}

	err := s.checkinRepositroy.Update(Id, data, detail)

	return err
}
