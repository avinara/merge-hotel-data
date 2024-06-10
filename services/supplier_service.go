package services

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/merge-hotel-data/config"
	"github.com/merge-hotel-data/errors"
	"github.com/merge-hotel-data/model"
)

type SupplierServiceAPI interface {
	GetSupplierData(searchStr string, searchValue string) ([]model.SupplierData, errors.ErrorInterface)
}

type SupplierService struct {
	Config config.Config
}

func NewSupplierService(Config config.Config) SupplierServiceAPI {
	return &SupplierService{
		Config: Config,
	}
}

func (s *SupplierService) GetSupplierData(searchStr string, searchValue string) (supplierDataArray []model.SupplierData, err errors.ErrorInterface) {
	for _, supplier := range s.Config.SupplierConfig {
		source := supplier.Source
		resp, err := http.Get(source)
		if err != nil {
			return nil, errors.UnableToGetTheSupplierDataError()
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.UnableToReadTheSupplierDataError()
		}

		if resp.StatusCode != http.StatusOK {
			return nil, errors.StatusCodeMismatchError()
		}

		var supplierAPIResponse []map[string]interface{}
		err = json.Unmarshal(body, &supplierAPIResponse)
		if err != nil {
			return nil, errors.UnableToDecodeTheSupplierDataError()
		}

		for _, m := range supplierAPIResponse {
			var supplierData model.SupplierData
			for key, value := range m {
				if value == "" {
					value = nil
				}
				if value != nil {
					switch key {
					case supplier.ResponseFormat["id"]:
						supplierData.Id = value.(string)
					case supplier.ResponseFormat["destination_id"]:
						supplierData.DestinationId = value.(float64)
					case supplier.ResponseFormat["name"]:
						supplierData.Name = value.(string)
					case supplier.ResponseFormat["lat"]:
						supplierData.Location.Latitude = value.(float64)
					case supplier.ResponseFormat["lng"]:
						supplierData.Location.Longitude = value.(float64)
					case supplier.ResponseFormat["description"]:
						supplierData.Description = value.(string)
					case supplier.ResponseFormat["address"]:
						supplierData.Location.Address = value.(string)
					case supplier.ResponseFormat["location.address"]:
						addressMap := value.(map[string]interface{})
						supplierData.Location.Address = addressMap["address"].(string)
						supplierData.Location.Country = addressMap["country"].(string)
					case supplier.ResponseFormat["city"]:
						supplierData.Location.City = value.(string)
					case supplier.ResponseFormat["country"]:
						supplierData.Location.Country = value.(string)
					case supplier.ResponseFormat["amenities"]:
						supplierData.Amenities.General = interfaceSliceToStringSlice(value.([]interface{}))
					case supplier.ResponseFormat["amenities.general"]:
						general := getDataFromArrayOfStrings(value.(map[string]interface{}), "general")
						supplierData.Amenities.General = general
						rooms := getDataFromArrayOfStrings(value.(map[string]interface{}), "room")
						supplierData.Amenities.Rooms = rooms
					case supplier.ResponseFormat["images"]:
						amenities := getImageDataFromArrayOfMaps(value.(map[string]interface{}), "amenities")
						supplierData.Images.Amenities = amenities
						site := getImageDataFromArrayOfMaps(value.(map[string]interface{}), "site")
						supplierData.Images.Site = site
						rooms := getImageDataFromArrayOfMaps(value.(map[string]interface{}), "rooms")
						supplierData.Images.Rooms = rooms
					case supplier.ResponseFormat["booking_conditions"]:
						supplierData.BookingConditions = interfaceSliceToStringSlice(value.([]interface{}))
					}
				}
			}

			if searchStr == "id" && supplierData.Id == searchValue {
				supplierDataArray = append(supplierDataArray, supplierData)
			}
			if searchStr == "destination_id" {
				destinationId, _ := strconv.ParseFloat(searchValue, 64)
				if destinationId == supplierData.DestinationId {
					supplierDataArray = append(supplierDataArray, supplierData)
				}
			}
		}
	}
	return supplierDataArray, nil
}

func getDataFromArrayOfStrings(value map[string]interface{}, searchString string) []string {
	temp := value[searchString]
	tempo := temp.([]interface{})
	var result []string
	for _, x := range tempo {
		result = append(result, x.(string))
	}
	return result
}

func getImageDataFromArrayOfMaps(value map[string]interface{}, searchString string) []model.ImageTemplate {
	temp := value[searchString]
	if temp == nil {
		return nil
	}
	tempo := temp.([]interface{})
	var imageArray []model.ImageTemplate
	for _, x := range tempo {
		midMap := x.(map[string]interface{})
		var image model.ImageTemplate
		for key, value := range midMap {
			if value != "" {
				switch key {
				case "url":
					image.Link = value.(string)
				case "link":
					image.Link = value.(string)
				case "description":
					image.Description = value.(string)
				case "caption":
					image.Description = value.(string)
				}
			}
		}
		imageArray = append(imageArray, image)
	}
	return imageArray
}

func interfaceSliceToStringSlice(interfaces []interface{}) []string {
	strings := make([]string, len(interfaces))
	for i, v := range interfaces {
		str, ok := v.(string)
		if !ok {
			return nil
		}
		strings[i] = str
	}
	return strings
}
