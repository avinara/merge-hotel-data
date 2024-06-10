package model

type SupplierData struct {
	Id                string         `json:"id"`
	DestinationId     float64        `json:"destination_id"`
	Name              string         `json:"name"`
	Location          HotelLocation  `json:"location"`
	Description       string         `json:"description"`
	Amenities         HotelAmenities `json:"amenities"`
	Images            HotelImages    `json:"images"`
	BookingConditions []string       `json:"booking_conditions"`
}

type HotelLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
}

type HotelAmenities struct {
	General []string `json:"general"`
	Rooms   []string `json:"rooms"`
}

type HotelImages struct {
	Rooms     []ImageTemplate `json:"rooms"`
	Site      []ImageTemplate `json:"site"`
	Amenities []ImageTemplate `json:"amenities"`
}

type ImageTemplate struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
