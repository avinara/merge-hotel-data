package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/merge-hotel-data/model"
	"github.com/merge-hotel-data/services"
	"github.com/merge-hotel-data/utils"
)

type MergeHotelDataControllerAPI interface {
	GetHotelData(w http.ResponseWriter, r *http.Request)
}

type MergeHotelDataController struct {
	MergeHotelDataService services.MergeHotelDataServiceAPI
	SupplierService       services.SupplierServiceAPI
}

func NewMergeHotelDataController(MergeHotelDataService services.MergeHotelDataServiceAPI, SupplierService services.SupplierServiceAPI) *MergeHotelDataController {
	return &MergeHotelDataController{
		MergeHotelDataService: MergeHotelDataService,
		SupplierService:       SupplierService,
	}
}

func (m *MergeHotelDataController) GetHotelData(w http.ResponseWriter, r *http.Request) {

	searchStr := r.URL.Query().Get("searchStr")
	searchValue := r.URL.Query().Get("searchValue")
	var hotels []model.SupplierData

	if searchStr == "id" {
		searchValues := strings.Split(searchValue, ",")

		for _, searchString := range searchValues {
			supplierData, err := m.SupplierService.GetSupplierData(searchStr, searchString)

			if err != nil {
				utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
				return
			}
			hotel := m.MergeHotelDataService.MergeHotelDataForHotelList(searchString, supplierData)

			if hotel == nil {
				utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
				return
			}
			hotels = append(hotels, *hotel)
		}
	}

	if searchStr == "destination_id" {
		supplierData, err := m.SupplierService.GetSupplierData(searchStr, searchValue)

		if err != nil {
			utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
			return
		}
		destinationId, _ := strconv.ParseFloat(searchValue, 64)
		hotels = m.MergeHotelDataService.MergeHotelDataForDestinationId(&destinationId, supplierData)

		if hotels == nil {
			utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
			return
		}
	}
	json.NewEncoder(w).Encode(hotels)
}
