package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Fetching weather data..."
	s.Color("yellow")
	s.Start()

	// Fetch weather data
	weatherDataBytes, err := weatherService.GetWeatherByCoordinates(lat, lon)

	// Stop the spinner
	s.Stop()

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
	// Get the formatted weather display as a string
	formattedWeather := weather.FormatWeatherData(weatherData)
	// Print it to the console
	fmt.Print(formattedWeather)
}

func displayError(err error) {
	fmt.Printf("Error: %s\n", err)
}

func printUsage() {

	bold := color.New(color.Bold).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(bold("=== Weather App CLI ==="))
	fmt.Println("Get current weather information for any city worldwide.")

	fmt.Println(blue("USAGE:"))
	fmt.Printf("  %s [options] -city <city_name>\n\n", os.Args[0])

	fmt.Println(blue("OPTIONS:"))
	fmt.Println("  -city <city_name>    City to get weather for")
	fmt.Println("  -i                   Run in interactive mode")
	fmt.Println("  -h                   Display this help message")

	fmt.Println(blue("EXAMPLES:"))
	fmt.Printf("  %s -city \"New York\"\n", os.Args[0])
	fmt.Printf("  %s -i\n\n", os.Args[0])

	fmt.Println(yellow("Interactive mode commands:"))
	fmt.Println("  [city name]   - Get weather for a city")
	fmt.Println("  help          - Display help information")
	fmt.Println("  exit          - Quit the application")

}

func RunInteractive(weatherService *weather.WeatherService) {
	city := getCity()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Interactive Weather App ===")
	fmt.Println("Type 'exit' to quit, 'help' for commands")

	for {
		color.New(color.FgCyan).Print("\nEnter city name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		switch strings.ToLower(input) {
		case "exit", "quit":
			fmt.Println("Goodbye!")
			return
		case "help":
			displayHelp()
			continue
		}

		// Find matching cities
		locations, err := weatherService.FindCities(input)
		if err != nil {
			displayError(err)
			continue
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
}

func displayHelp() {
	helpColor := color.New(color.FgHiWhite)
	cmdColor := color.New(color.FgHiCyan, color.Bold)

	helpColor.Println("\nAvailable commands:")
	cmdColor.Print("  [city name]")
	helpColor.Println(" - Get weather for a city")
	cmdColor.Print("  help")
	helpColor.Println(" - Display this help message")
	cmdColor.Print("  exit")
	helpColor.Println(" - Quit the application")
}
