package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shreya-sk/CLI-Weather-App/cmd"
	"github.com/shreya-sk/CLI-Weather-App/internal/config"
	"github.com/shreya-sk/CLI-Weather-App/internal/weather"
)

// main() - Application entry point
// setup() (*config.Config, *weather.WeatherService, error) - Sets up dependencies

func main() {
	// Load .env file at application startup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// url := os.Getenv("OPENWEATHER_API_URL") + "?q=" + "sydney" + "&appid=" + os.Getenv("OPENWEATHER_API_KEY")
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	svc, err := setup()
	if err != nil {
		panic(err)
	}

	cmd.RunCLI(svc)

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
