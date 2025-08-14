package main

import "fmt"

type StatisticsDisplay struct {
	maxTemp     int
	minTemp     int
	tempSum     int
	numReadings int

	weatherData Subject
}

func NewStatisticsDisplay(weatherData Subject) {
	sd := StatisticsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(&sd)
}

func (s *StatisticsDisplay) Update(temp, humidity, pressure int) {
	s.tempSum += temp
	s.numReadings++

	if temp > s.maxTemp {
		s.maxTemp = temp
	}

	if temp < s.minTemp {
		s.minTemp = temp
	}

	s.Display()
}

func (s *StatisticsDisplay) Display() {
	fmt.Println("Avg/Max/Min temperature = ", (s.tempSum / s.numReadings), "/", s.maxTemp, "/", s.minTemp)
}
