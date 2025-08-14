package main

import "fmt"

type HeatIndexDisplay struct {
	heatIndex   int
	weatherData Subject
}

func NewHeatIndexDisplay(weatherData Subject) {

	weatherData.RegisterObserver(&HeatIndexDisplay{weatherData: weatherData})
}

func (hid *HeatIndexDisplay) Update(t, rh, pressure int) {
	hid.heatIndex = int(hid.computeHeatIndex(float64(t), float64(rh)))
	hid.Display()
}

func (hid *HeatIndexDisplay) computeHeatIndex(t, rh float64) float64 {
	index := (float64)((16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) + (0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) + (0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) + (0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 * (rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) + (0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) + 0.000000000843296*(t*t*rh*rh*rh)) - (0.0000000000481975 * (t * t * t * rh * rh * rh)))
	return index
}

func (hid *HeatIndexDisplay) Display() {
	fmt.Println("Heat index is ", hid.heatIndex)
}
