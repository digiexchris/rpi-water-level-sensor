package sensors

/**
This feeds an array of booleans into a channel as readings, meaning
each sensor is either off or on
*/

type Sensors interface {
	Connect() error
	Disconnect() error
	Run()
	Stop()
}

type Reading struct {
	Sensor int
	On     bool
	Err    error
}
