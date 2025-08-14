package main

import "fmt"

type ForecastDisplay struct {
	currentPressure int
	lastPressure    int

	weatherData Subject
}

func NewForecastDisplay(weatherData Subject) {
	fd := ForecastDisplay{
		weatherData: weatherData,
	}

	weatherData.RegisterObserver(&fd)
}

func (f *ForecastDisplay) Update(temperature, humidity, pressure int) {
	f.lastPressure = f.currentPressure
	f.currentPressure = pressure

	f.Display()
}

func (f *ForecastDisplay) Display() {
	fmt.Print("Forecast: ")
	if f.currentPressure > f.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if f.currentPressure == f.lastPressure {
		fmt.Println("More of the same")
	} else if f.currentPressure < f.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}
