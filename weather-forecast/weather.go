// Package weather: Gives the current weather for a specific location.
package weather

var (
	// Instantiate condition string
	CurrentCondition string
	// Instantiate location string
	CurrentLocation string
)

// Log take a city and condition string and return the current weather conditions for that location.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
