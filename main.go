package main

import (
	"github.com/digiexchris/water-level-sensor/arduinoWaterSensor"
	"github.com/digiexchris/water-level-sensor/configuration"
	"github.com/digiexchris/water-level-sensor/httpserver"
	"github.com/digiexchris/water-level-sensor/sensors"
)

var Reading chan sensors.Reading

func main() {

	err := configuration.Load()
	if err != nil {
		panic(err)
	}

	Reading = make(chan sensors.Reading)

	s := arduinoWaterSensor.New(Reading)

	err = s.Connect()
	if err != nil {
		panic(err)
	}
	s.Run()

	server := httpserver.New()

	go func(r chan sensors.Reading) {
		for {
			reading := <-r
			server.SetReading(reading.Sensor, reading.On)
		}

	}(Reading)

	server.Run()
}
