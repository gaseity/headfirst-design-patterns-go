package main

import (
	"slices"
)

type WeatherData struct {
	observers []Observer

	temperature int
	humidity    int
	pressure    int
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	// for _, observer := range wd.observers {
	// 	if observer == o {
	// 		return
	// 	}
	// }
	if slices.Contains(wd.observers, o) {
		return
	}
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	// for i, observer := range wd.observers {
	// 	if observer == o {
	// 		wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
	// 		break
	// 	}
	// }
	index := slices.Index(wd.observers, o)
	if index < 0 {
		return
	}
	wd.observers = append(wd.observers[:index], wd.observers[index+1:]...)
}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure int) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}
