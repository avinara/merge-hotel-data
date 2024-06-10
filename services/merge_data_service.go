package services

import (
	"github.com/merge-hotel-data/config"
	"github.com/merge-hotel-data/errors"
	"github.com/merge-hotel-data/model"
)

type MergeHotelDataServiceAPI interface {
	MergeHotelData(supplierData []model.SupplierData) ([]model.Hotel, errors.ErrorInterface)
}

type MergeHotelDataService struct {
	Config config.Config
}

func NewMergeHotelDataService(Config config.Config) MergeHotelDataServiceAPI {
	return &MergeHotelDataService{
		Config: Config,
	}
}

func (s *MergeHotelDataService) MergeHotelData(supplierData []model.SupplierData) ([]model.Hotel, errors.ErrorInterface) {
	for _, supplier := range supplierData {

	}
	return nil, nil
}
