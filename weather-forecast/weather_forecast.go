// Package weather provides weather forecasts.
package weather

// CurrentCondition is the current weather conditions.
var CurrentCondition string

// CurrentLocation is the current location where weather conditions are occurring.
var CurrentLocation string

// Forecast returns a string representation of the given city's weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
