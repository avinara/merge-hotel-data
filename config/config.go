// config/config.go
package config

import (
	"encoding/json"
	"os"
)

// Config struct represents the configuration
type Config struct {
	SupplierConfig  []SupplierConfig `json:"supplier_config"`
	AmenitiesConfig AmenitiesConfig  `json:"amenities_config"`
}

type AmenitiesConfig struct {
	General map[string]string `json:"general"`
	Rooms   map[string]string `json:"rooms"`
}
type SupplierConfig struct {
	Source         string            `json:"source"`
	Name           string            `json:"name"`
	ResponseFormat map[string]string `json:"response_format"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
