package weather

import (
	"fmt"
	"math"
	"strings"
	"time"
	// In service.go:
)

// FormatWeatherData(data *WeatherData) string - Formats weather data for terminal display
// FormatTemperature(temp float64) string - Formats temperature with units
// FormatTime(timestamp int64, timezone int) string - Formats time values
// FormatWindDirection(degrees float64) string - Converts wind degrees to cardinal direction

func FormatWeatherData(data *WeatherData) string {
	// Create a string builder to efficiently build our output string
	var output strings.Builder

	// Add a header with location information
	output.WriteString("=====================================\n")
	output.WriteString(fmt.Sprintf("    Weather for %s, %s\n", data.Name, data.Sys.Country))
	output.WriteString("=====================================\n")

	// Now add sections for different types of data:
	// 1. Current conditions section
	// Hint: You'll want to access data.Weather[0] if it exists
	// Don't forget to check if the Weather slice has elements before accessing
	output.WriteString("Current Conditions:\n")
	if len(data.Weather) > 0 {

		for i := 0; i < len(data.Weather); i++ {
			output.WriteString(fmt.Sprintf("   %s\n", data.Weather[i].Description))
		}
	} else {
		output.WriteString("No Information available on current conditions.")
	}
	output.WriteString("\n")

	// 2. Temperature section
	output.WriteString("Temperature:\n")
	output.WriteString(fmt.Sprintf("   Current: %s\n", FormatTemperature(data.Main.Current)))
	output.WriteString(fmt.Sprintf("   Min/Max: %s/%s\n", FormatTemperature(data.Main.Minimum), FormatTemperature(data.Main.Maximum)))
	output.WriteString(fmt.Sprintf("   Feels Like: %s\n", FormatTemperature(data.Main.FeelsLike)))
	output.WriteString(fmt.Sprintf("   Humidity: %d%%\n", data.Main.Humidity))

	output.WriteString("\n")
	// Format current, feels like, min and max from data.Main

	// 3. Wind information
	output.WriteString("Wind:\n")
	output.WriteString(fmt.Sprintf("   Speed: %.2fkmph %s\n", data.Wind.Speed, FormatWindDirection(data.Wind.Degrees)))

	// 4. Additional information
	// You could include coordinates or other data you find useful

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
