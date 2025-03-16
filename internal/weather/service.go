package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/shreya-sk/CLI-Weather-App/internal/config"
	// In service.go:
)

// type WeatherService struct - Service for interacting with the weather API
// NewWeatherService(config *config.Config) *WeatherService - Constructor for the service
// GetWeatherByCity(city string) (*WeatherData, error) - Fetches weather for a city
// parseWeatherResponse(data []byte) (*WeatherData, error) - Parses JSON response
type WeatherService struct {
	config *config.Config
}

// Give the NewWeatherService a config, and it will create a new service using above defined config
func NewWeatherService(config *config.Config) *WeatherService {
	return &WeatherService{
		config: config,
	}
}

func (s *WeatherService) GetWeatherByCity(city string) ([]byte, error) {
	url := s.config.BaseURL + "?q=" + city + "&appid=" + s.config.APIKey + "&units=" + s.config.Units
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//{"coord":{"lon":10.99,"lat":44.34},
	// "weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],
	// "base":"stations",
	// "main":{"temp":285.51,"feels_like":284.55,"temp_min":282.71,"temp_max":285.97,"pressure":1011,"humidity":67,"sea_level":1011,"grnd_level":944},
	// "visibility":10000,
	// "wind":{"speed":0.83,"deg":176,"gust":2.74},
	// "clouds":{"all":97},
	// "dt":1741522641,
	// "sys":{"type":2,"id":2004688,"country":"IT","sunrise":1741498765,"sunset":1741540437},
	// "timezone":3600,"id":3163858,"name":"Zocca","cod":200}

	return body, nil
}

// Fetch weather data from the API

func (s *WeatherService) ParseWeatherResponse(data []byte) (*WeatherData, error) {
	// Create a new WeatherData struct
	weatherData := &WeatherData{}
	// Unmarshal the JSON response into the WeatherData struct
	err := json.Unmarshal(data, weatherData)
	if err != nil {
		return nil, err
	}

	return weatherData, nil

	// Parse the JSON response
}
func (s *WeatherService) FindCities(cityName string) ([]Location, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s",
		url.QueryEscape(cityName), s.config.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (s *WeatherService) GetWeatherByCoordinates(lat, lon float64) ([]byte, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=%s",
		s.config.BaseURL, lat, lon, s.config.APIKey, s.config.Units)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
