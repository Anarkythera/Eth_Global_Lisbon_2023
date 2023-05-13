package main

import (
	"os"

	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
)

type config struct {
	ProdutionRateHour   int `yaml:"currentProductionRateHour"`
	ConsumptionRateHour int `yaml:"currentConsumptionRateHour"`
}

var ControllerPort = ":8082"

func main() {
	initConfig()
	go startEcho()
	select {}
}

func startEcho() {
	e := echo.New()
	e.Logger.Fatal(e.Start(ControllerPort))
}

func initConfig() (config, error) {
	var cfg config
	f, err := os.Open("config.yml")
	if err != nil {
		return config{}, err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return config{}, nil
	}

	return cfg, nil
}
