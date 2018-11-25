package main

import (
	"github.com/digiexchris/water-level-monitor/Sensors"
	"github.com/digiexchris/water-level-monitor/httpserver"
)

var Reading chan Sensors.Reading

func main() {

	Reading = make(chan Sensors.Reading)

	//s := Sensors.New(Reading)
	//
	//s.Connect()
	//s.Run()

	server := httpserver.New()

	//go func(r chan Sensors.Reading) {
	//	for {
	//		reading := <-r
	//		server.SetReading(reading.Sensor, reading.On)
	//	}
	//
	//}(Reading)

	server.Run()
}