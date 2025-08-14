package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature int
	humidity    int

	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) {

	ccd := CurrentConditionsDisplay{
		weatherData: weatherData,
	}

	weatherData.RegisterObserver(&ccd)
}

func (c *CurrentConditionsDisplay) Update(temperature, humidity, pressure int) {
	c.temperature = temperature
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Println("Current conditions: ", c.temperature, "F degrees and ", c.humidity, "% humidity")
}
