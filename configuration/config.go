package configuration

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

var App Config

type Config struct {
	PortName string `json:"Port"`
}

func Load() error {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file := fmt.Sprintf("%s/config.json", dir)
	log.Printf("Attempting to load configuration from %s", file)

	App = Config{}
	err = gonfig.GetConf(file, &App)

	return err
}
