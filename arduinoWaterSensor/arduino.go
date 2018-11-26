package arduinoWaterSensor

import (
	"errors"
	"github.com/digiexchris/water-level-sensor/configuration"
	"github.com/digiexchris/water-level-sensor/sensors"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"strconv"
	"strings"
)

type arduino struct {
	options serial.OpenOptions
	channel chan sensors.Reading
	port    io.ReadWriteCloser
	stop    bool
}

func New(readings chan sensors.Reading) sensors.Sensors {

	options := serial.OpenOptions{
		PortName:        configuration.App.PortName,
		BaudRate:        19200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	return &arduino{
		options: options,
		channel: readings,
	}
}

func (s *arduino) Connect() error {
	// Open the port.
	port, err := serial.Open(s.options)
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
	var done bool

	var serialResponse []byte

	buf := make([]byte, 1)
	for done == false {
		n, err := s.port.Read(buf)
		if err != nil {
			reading.Err = err
			s.channel <- reading
		}

		if n == 0 {
			reading.Err = errors.New("No data")
			s.channel <- reading
			return
		}

		if string(buf) != ";" {
			serialResponse = append(serialResponse, buf[0])
		} else {
			done = true
		}
	}

	//log.Println(string(serialResponse))

	r := strings.Split(string(serialResponse), ":")

	var err error

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
