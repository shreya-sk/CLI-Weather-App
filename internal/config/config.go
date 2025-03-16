package config

import (
	"errors"
	"os"
)

type Config struct {
	APIKey  string
	BaseURL string
	Units   string
}

func LoadConfig() (*Config, error) {
	// Load configuration from file
	apiKey, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	return &Config{
		APIKey:  apiKey,
		BaseURL: GetBaseURL(),
		Units:   GetUnits(),
	}, nil
	// Load configuration from environment variables
}

func GetAPIKey() (string, error) {

	key := os.Getenv("OPENWEATHER_API_KEY")
	// error check
	if key == "" {
		return "", errors.New("API key not found")
	}
	return key, nil
	// Returns the API key
}

func GetBaseURL() string {
	baseURL := os.Getenv("OPENWEATHER_API_URL")
	if baseURL == "" {
		return "API base URL not found"
	}
	return baseURL
	// Returns the API base URL
}

func GetUnits() string {
	units := os.Getenv("OPENWEATHER_UNITS")
	if units == "" {
		return "metric"
	}
	return units
	// Returns the unit system (metric/imperial)
}
