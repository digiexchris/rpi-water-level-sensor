package arduinoWaterSensor

import (
	"github.com/digiexchris/water-level-sensor/configuration"
	"github.com/digiexchris/water-level-sensor/sensors"
	"github.com/tarm/serial"
	"strconv"
	"strings"
	"time"
)

type arduino struct {
	options *serial.Config
	channel chan sensors.Reading
	port    *serial.Port
	stop    bool
}

func New(readings chan sensors.Reading) sensors.Sensors {

	options := serial.Config{
		Name:        configuration.App.PortName,
		Baud:        19200,
		StopBits:    1,
		ReadTimeout: time.Second * 30,
	}

	return &arduino{
		options: &options,
		channel: readings,
	}
}

func (s *arduino) Connect() error {
	// Open the port.
	port, err := serial.OpenPort(s.options)
	if err != nil {
		return err
	}

	s.port = port

	return nil
}

func (s *arduino) Disconnect() error {
	s.stop = true
	return s.port.Close()
}

func (s *arduino) Run() {
	s.stop = false
	go func(s *arduino) {
		for s.stop == false {
			s.read()
		}
	}(s)
}

/**
format: 1:1
port number, on or off

Will block until the arduinoWaterSensor writes or 30 seconds happens
*/
func (s *arduino) read() {
	reading := sensors.Reading{}
	buf := make([]byte, 3)
	_, err := s.port.Read(buf)
	if err != nil {
		reading.Err = err
		s.channel <- reading
	}

	r := strings.Split(string(buf), ",")

	reading.Sensor, err = strconv.Atoi(r[0])
	if err != nil {
		reading.Err = err
		s.channel <- reading
	}

	on, err := strconv.Atoi(r[1])
	if err != nil {
		reading.Err = err
		s.channel <- reading
	}

	reading.On = on != 0

	s.channel <- reading
}