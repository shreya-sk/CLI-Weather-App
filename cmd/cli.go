package cmd

import (
	"flag"
	"fmt"

	"github.com/shreya-sk/CLI-Weather-App/internal/weather"
)

// RunCLI(weatherService *weather.WeatherService) - Main CLI entrypoint
// getCity() string - Gets city from command line args or user input
// displayWeather(weatherData *weather.WeatherData) - Displays weather information
// displayError(err error) - Formats and displays error messages
// printUsage() - Displays usage information

func RunCLI(weatherService *weather.WeatherService) {
	city := getCity()

	// Find matching cities
	locations, err := weatherService.FindCities(city)
	if err != nil {
		displayError(err)
		return
	}

	if len(locations) == 0 {
		fmt.Printf("No cities found matching '%s'\n", city)
		return
	}

	// If only one city found, use it directly
	if len(locations) == 1 {
		fetchAndDisplayWeather(weatherService, locations[0].Lat, locations[0].Lon)
		return
	}

	// Display city options
	fmt.Println("Multiple cities found. Please select one:")
	for i, loc := range locations {
		state := ""
		if loc.State != "" {
			state = ", " + loc.State
		}
		fmt.Printf("%d. %s, %s%s\n", i+1, loc.Name, loc.Country, state)
	}

	// Get user selection
	var selection int
	fmt.Print("Enter number: ")
	_, err = fmt.Scanf("%d", &selection)
	if err != nil || selection < 1 || selection > len(locations) {
		fmt.Println("Invalid selection")
		return
	}

	// Fetch and display weather for selected city
	selectedLocation := locations[selection-1]
	fetchAndDisplayWeather(weatherService, selectedLocation.Lat, selectedLocation.Lon)
}

func fetchAndDisplayWeather(weatherService *weather.WeatherService, lat, lon float64) {
	weatherDataBytes, err := weatherService.GetWeatherByCoordinates(lat, lon)
	if err != nil {
		displayError(err)
		return
	}

	weatherData, err := weatherService.ParseWeatherResponse(weatherDataBytes)
	if err != nil {
		displayError(err)
		return
	}

	displayWeather(weatherData)
}

func getCity() string {
	var city string
	flag.StringVar(&city, "city", "", "City to get weather for")
	flag.Parse()
	if city == "" {
		printUsage()
	}

	return city
}

func displayWeather(weatherData *weather.WeatherData) {

	formattedWeather := weather.FormatWeatherData(weatherData)
	fmt.Println(formattedWeather)

}

func displayError(err error) {
	fmt.Printf("Error: %s\n", err)

}

func printUsage() {
	fmt.Println("Usage: go run main.go -city <city_name>")

}
