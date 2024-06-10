package services

import (
	"strings"

	"github.com/merge-hotel-data/config"
	"github.com/merge-hotel-data/model"
)

type MergeHotelDataServiceAPI interface {
	MergeHotelDataForHotelList(hotelId string, supplierData []model.SupplierData) *model.SupplierData
	MergeHotelDataForDestinationId(destinationId *float64, supplierData []model.SupplierData) []model.SupplierData
}

type MergeHotelDataService struct {
	Config config.Config
}

func NewMergeHotelDataService(Config config.Config) MergeHotelDataServiceAPI {
	return &MergeHotelDataService{
		Config: Config,
	}
}

func (s *MergeHotelDataService) MergeHotelDataForHotelList(hotelId string, supplierData []model.SupplierData) *model.SupplierData {
	HotelResponse := &model.SupplierData{}
	for _, supplier := range supplierData {
		if hotelId != supplier.Id {
			continue
		}

		HotelResponse.Id = supplier.Id
		HotelResponse.DestinationId = supplier.DestinationId
		if HotelResponse.Name == "" || len(HotelResponse.Name) < len(supplier.Name) {
			HotelResponse.Name = supplier.Name
		}
		if HotelResponse.Location.Address == "" || len(HotelResponse.Location.Address) < len(supplier.Location.Address) {
			HotelResponse.Location.Address = supplier.Location.Address
		}
		if supplier.Location.Latitude != 0 {
			HotelResponse.Location.Latitude = supplier.Location.Latitude
		}
		if supplier.Location.Longitude != 0 {
			HotelResponse.Location.Longitude = supplier.Location.Longitude
		}
		if supplier.Location.City != "" && len(supplier.Location.City) > 0 && len(supplier.Location.City) > len(HotelResponse.Location.City) {
			HotelResponse.Location.City = supplier.Location.City
		}
		if supplier.Location.Country != "" && len(supplier.Location.Country) > 0 && len(supplier.Location.Country) > len(HotelResponse.Location.Country) {
			HotelResponse.Location.Country = supplier.Location.Country
		}
		if supplier.Amenities.General != nil && len(supplier.Amenities.General) > 0 {
			HotelResponse.Amenities.General = append(HotelResponse.Amenities.General, supplier.Amenities.General...)
		}
		if supplier.Amenities.Rooms != nil && len(supplier.Amenities.Rooms) > 0 {
			HotelResponse.Amenities.Rooms = append(HotelResponse.Amenities.Rooms, supplier.Amenities.Rooms...)
		}
		if len(supplier.Description) > 0 && len(supplier.Description) > len(HotelResponse.Description) {
			HotelResponse.Description = supplier.Description
		}
		if supplier.Images.Site != nil && len(supplier.Images.Site) > 0 && len(supplier.Images.Site) > len(HotelResponse.Images.Site) {
			HotelResponse.Images.Site = supplier.Images.Site
		}
		if supplier.Images.Rooms != nil && len(supplier.Images.Rooms) > 0 && len(supplier.Images.Rooms) > len(HotelResponse.Images.Rooms) {
			HotelResponse.Images.Rooms = supplier.Images.Rooms
		}
		if supplier.Images.Amenities != nil && len(supplier.Images.Amenities) > 0 && len(supplier.Images.Amenities) > len(HotelResponse.Images.Amenities) {
			HotelResponse.Images.Amenities = supplier.Images.Amenities
		}
		if supplier.BookingConditions != nil && len(supplier.BookingConditions) > 0 && len(supplier.BookingConditions) > len(HotelResponse.BookingConditions) {
			HotelResponse.BookingConditions = supplier.BookingConditions
		}
	}
	amenitiesArray := append(HotelResponse.Amenities.Rooms, HotelResponse.Amenities.General...)
	var generalArray, roomsArray []string
	for i := 0; i < len(amenitiesArray); i++ {
		if _, ok := s.Config.AmenitiesConfig.General[strings.Trim(amenitiesArray[i], "")]; ok {
			generalArray = append(generalArray, s.Config.AmenitiesConfig.General[strings.Trim(amenitiesArray[i], "")])
		} else if _, ok := s.Config.AmenitiesConfig.Rooms[strings.Trim(amenitiesArray[i], "")]; ok {
			roomsArray = append(roomsArray, s.Config.AmenitiesConfig.Rooms[strings.Trim(amenitiesArray[i], "")])
		}
	}
	roomsArray = removeDuplicates(roomsArray)
	generalArray = removeDuplicates(generalArray)
	HotelResponse.Amenities.Rooms = roomsArray
	HotelResponse.Amenities.General = generalArray
	return HotelResponse
}

func (s *MergeHotelDataService) MergeHotelDataForDestinationId(destinationId *float64, supplierData []model.SupplierData) []model.SupplierData {

	var hotelList []string
	var hotels []model.SupplierData
	for _, supplier := range supplierData {
		if *destinationId == supplier.DestinationId && !contains(hotelList, supplier.Id) {
			hotelList = append(hotelList, supplier.Id)
		}
	}

	for _, searchString := range hotelList {
		hotel := s.MergeHotelDataForHotelList(searchString, supplierData)

		if hotel == nil {
			return nil
		}
		hotels = append(hotels, *hotel)
	}

	return hotels
}

func contains(list []string, target string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		if !contains(list, item) {
			list = append(list, item)
		}
	}
	return list
}
