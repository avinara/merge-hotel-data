package controllers

import (
	"encoding/json"
	"net/http"

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

// @Summary Get Hotels
// @Description Get Hotels
// @Tags Hotels
// @Produce json
// @Success 200 {array} model.Hotel
// @Failure 500 {object} model.ErrorResponse
// @Router /hotels [get]
func (m *MergeHotelDataController) GetHotelData(w http.ResponseWriter, r *http.Request) {

	searchStr := r.URL.Query().Get("searchStr")
	searchValue := r.URL.Query().Get("searchValue")
	supplierData, err := m.SupplierService.GetSupplierData(searchStr, searchValue)

	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	hotels, err := m.MergeHotelDataService.MergeHotelData(supplierData)

	if err != nil || hotels == nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	json.NewEncoder(w)
}
