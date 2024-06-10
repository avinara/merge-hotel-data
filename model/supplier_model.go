package model

type SupplierData struct {
	Id                string
	DestinationId     float64
	Name              string
	Location          HotelLocation
	Description       string
	Amenities         HotelAmenities
	Images            HotelImages
	BookingConditions []string
}

type HotelLocation struct {
	Latitude  float64
	Longitude float64
	Address   string
	City      string
	Country   string
}

type HotelAmenities struct {
	General []string
	Rooms   []string
}

type HotelImages struct {
	Rooms     []ImageTemplate
	Site      []ImageTemplate
	Amenities []ImageTemplate
}

type ImageTemplate struct {
	Link        string
	Description string
}
