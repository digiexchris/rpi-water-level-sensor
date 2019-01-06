package main

import (
	"github.com/digiexchris/water-level-sensor/arduinoWaterSensor"
	"github.com/digiexchris/water-level-sensor/configuration"
	"github.com/digiexchris/water-level-sensor/httpserver"
	"github.com/digiexchris/water-level-sensor/sensors"
	"log"
	"time"
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

	go func(r chan sensors.Reading, s sensors.Sensors) {
		for {
			reading := <-r
			if reading.Err != nil {
				if reading.Err.Error() == "EOF" {
					s.Stop()
					reconnect(s)
					s.Run()
				}
			}
			server.SetReading(reading.Sensor, reading.On)
			server.SetError(reading.Err)
		}

	}(Reading, s)

	server.Run()
}

func reconnect(s sensors.Sensors) {

	time.Sleep(time.Second * 3)
	log.Println("Reconnecting")
	s.Disconnect()

	for times := 0; times != 100; times++ {
		time.Sleep(time.Second * 10)
		err := s.Connect()
		if err != nil {
			log.Println(err)
		} else {
			log.Println ("Reconnected")
			break
		}
	}
}