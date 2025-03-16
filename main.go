package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shreya-sk/CLI-Weather-App/cmd"
	"github.com/shreya-sk/CLI-Weather-App/internal/config"
	"github.com/shreya-sk/CLI-Weather-App/internal/weather"
)

// main() - Application entry point
// setup() (*config.Config, *weather.WeatherService, error) - Sets up dependencies
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file")
	}

	interactive := flag.Bool("i", true, "Run in interactive mode")

	weatherService, err := setup()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	// Check which mode to run in
	if *interactive {
		cmd.RunInteractive(weatherService)
	} else {
		cmd.RunCLI(weatherService)
	}

}

func setup() (*weather.WeatherService, error) {
	// setup config

	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// setup weather service
	weatherService := weather.NewWeatherService(conf)

	return weatherService, nil
}
