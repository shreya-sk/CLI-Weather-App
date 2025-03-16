package weather

// In models.go:

// type WeatherData struct - Main structure to hold complete weather information
// type Temperature struct - Structure for temperature data (current, min, max, feels like)
// type Wind struct - Structure for wind information (speed, direction)
// type Location struct - Structure for location data (city, country, coordinates)
// type Condition struct - Structure for weather conditions (description, icon)

type WeatherData struct {
	Coord   Coord       `json:"coord"`
	Weather []Condition `json:"weather"`
	Main    Temperature `json:"main"`
	Wind    Wind        `json:"wind"`
	Sys     Sys         `json:"sys"`
	Name    string      `json:"name"`

	// other root level fields you need
}
type Location struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names,omitempty"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
} // Coord structure for coordinates

type Condition struct {
	Description string `json:"description"`
	Main        string `json:"main"`
} // Condition structure for weather conditions (description, icon)

type Temperature struct {
	Current   float64 `json:"temp"`
	Minimum   float64 `json:"temp_min"`
	Maximum   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
} // Temperature structure for temperature data (current, min, max, feels like)

type Wind struct {
	Speed   float64 `json:"speed"`
	Degrees float64 `json:"deg"`
} // Wind structure for wind information (speed, direction)

type Sys struct {
	Country string `json:"country"`
}
