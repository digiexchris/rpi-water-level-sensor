package main

import (
	"github.com/digiexchris/water-level-sensor/Sensors"
	"github.com/digiexchris/water-level-sensor/configuration"
	"github.com/digiexchris/water-level-sensor/httpserver"
)

var Reading chan Sensors.Reading

func main() {

	err := configuration.Load()
	if err != nil {
		panic(err)
	}

	Reading = make(chan Sensors.Reading)

	s := Sensors.New(Reading)

	err = s.Connect()
	if err != nil {
		panic(err)
	}
	s.Run()

	server := httpserver.New()

	go func(r chan Sensors.Reading) {
		for {
			reading := <-r
			server.SetReading(reading.Sensor, reading.On)
		}

	}(Reading)

	server.Run()
}
