package main

import "fmt"

func main() {

	weatherData := NewWeatherData()

	NewCurrentConditionsDisplay(weatherData)
	NewStatisticsDisplay(weatherData)
	NewForecastDisplay(weatherData)
	NewHeatIndexDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30)
	fmt.Println("--------")
	weatherData.SetMeasurements(82, 70, 29)
	fmt.Println("--------")
	weatherData.SetMeasurements(78, 90, 29)
}
