package weather

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/fatih/color"
	// In service.go:
)

// FormatWeatherData(data *WeatherData) string - Formats weather data for terminal display
// FormatTemperature(temp float64) string - Formats temperature with units
// FormatTime(timestamp int64, timezone int) string - Formats time values
// FormatWindDirection(degrees float64) string - Converts wind degrees to cardinal direction

func FormatWeatherData(weatherData *WeatherData) string {
	// In weather/formatter.go

	var output strings.Builder

	// Create color formatters
	headerColor := color.New(color.FgBlue, color.Bold)
	locationColor := color.New(color.FgYellow, color.Bold)
	sectionColor := color.New(color.FgGreen)
	valueColor := color.New(color.FgWhite)

	// Build the formatted output
	output.WriteString("\n\n")

	// Add header with color
	headerText := headerColor.Sprintf("=== WEATHER INFORMATION ===\n")
	output.WriteString(headerText)

	// Add location with color
	locationText := locationColor.Sprintf("Location: %s, %s\n\n", weatherData.Name, weatherData.Sys.Country)

	output.WriteString(locationText)

	// Add ASCII art and weather condition
	if len(weatherData.Weather) > 0 {
		iconCode := weatherData.Weather[0].Icon
		asciiArt := getWeatherASCII(iconCode)
		asciiText := color.New(color.FgHiYellow).Sprintf("%s\n", asciiArt)
		output.WriteString(asciiText)

		conditionText := valueColor.Sprintf("  %s\n\n", weatherData.Weather[0].Description)
		output.WriteString(conditionText)
	}

	// Add temperature section
	tempHeader := sectionColor.Sprintf("Temperature:\n")
	output.WriteString(tempHeader)

	tempDetails := valueColor.Sprintf("  Current: %.1f°C\n", weatherData.Main.Current) +
		valueColor.Sprintf("  Feels like: %.1f°C\n", weatherData.Main.FeelsLike) +
		valueColor.Sprintf("  Min/Max: %.1f°C/%.1f°C\n\n", weatherData.Main.Minimum, weatherData.Main.Maximum) +
		valueColor.Sprintf("  Humidity: %d%%\n\n", weatherData.Main.Humidity)
	output.WriteString(tempDetails)

	// Add conditions section
	condHeader := sectionColor.Sprintf("Conditions:\n")
	output.WriteString(condHeader)

	if len(weatherData.Weather) > 0 {
		condDetails := valueColor.Sprintf("  %s\n\n", weatherData.Weather[0].Description)
		output.WriteString(condDetails)
	}

	// Add wind section
	windHeader := sectionColor.Sprintf("Wind:\n")
	output.WriteString(windHeader)

	windDetails := valueColor.Sprintf("  Speed: %.2f kmph %s\n", weatherData.Wind.Speed, FormatWindDirection(weatherData.Wind.Degrees))
	output.WriteString(windDetails)

	return output.String()
}

func FormatTemperature(temp float64) string {
	// Format temperature with units
	return fmt.Sprintf("%.2f°C", temp)
}

func FormatTime(timestamp int64, timezone int) string {
	// Format time values
	loc := time.FixedZone("UTC", timezone*3600)
	return time.Unix(timestamp, 0).In(loc).String()
}

func FormatWindDirection(degrees float64) string {
	// Convert wind degrees to cardinal direction
	dirs := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	ix := int(math.Round(degrees / (360. / float64(len(dirs)))))
	return dirs[ix%len(dirs)]
}

func getWeatherASCII(iconCode string) string {
	// Match the icon code from OpenWeatherMap API
	switch iconCode {
	case "01d", "01n": // clear sky
		return `
    \   /
     .-.
  ― (   ) ―
     '-'
    /   \
        `
	case "02d", "02n", "03d", "03n", "04d", "04n": // clouds
		return `
      .--.
   .-(    ).
  (___.__)__)
            
        `
	case "09d", "09n", "10d", "10n": // rain
		return `
     .-.
    (   ).
   (___(__)
    ' ' ' '
   ' ' ' '
        `
	case "11d", "11n": // thunderstorm
		return `
      .-.
     (   ).
    (___(__)
     ⚡⚡⚡
      ⚡⚡
        `
	case "13d", "13n": // snow
		return `
      .-.
     (   ).
    (___(__)
     * * *
    * * *
        `
	case "50d", "50n": // mist/fog
		return `
      .-.
     (   ).
    (___(__)
    ― ― ― ―
   ― ― ― ―
        `
	default:
		return ""
	}
}
