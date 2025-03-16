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
	weatherDataBytes, err := weatherService.GetWeatherByCity(city)
	if err != nil {
		displayError(err)
	}
	weatherData, err := weatherService.ParseWeatherResponse(weatherDataBytes)
	if err != nil {
		displayError(err)
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
