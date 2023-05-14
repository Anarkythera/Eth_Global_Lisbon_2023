package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v3"
)

var ControllerPort = ":8082"
var CurrentConsumption = 0
var CurrentProduction = 0

type config struct {
	ProdutionRateHour   int `yaml:"currentProductionRateHour"`
	ConsumptionRateHour int `yaml:"currentConsumptionRateHour"`
}

func main() {

	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", cfg)
	go startEcho()
	readContract()
	go produceCounter()
	go consumptionCounter()
	select {}
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

func startEcho() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/config", ViewConfig)
	e.POST("/config", EditConfig)
	e.Logger.Fatal(e.Start(ControllerPort))

}

func produceCounter() {
	for {
		fmt.Printf("\n Consuming \n")
		time.Sleep(10 * time.Second)
		CurrentConsumption += 1
		fmt.Printf("\n CurrentConsumption %+v \n", CurrentConsumption)
	}
}

func consumptionCounter() {

	for {
		fmt.Printf("\n Producing \n")
		time.Sleep(10 * time.Second)
		CurrentProduction += 1
		fmt.Printf("\n CurrentProduction %+v \n", CurrentProduction)
	}
}

func readContract() {

	client, err := ethclient.Dial("https://api-testnet.polygonscan.com/")
	if err != nil {
		// handle error

	}

	address := common.HexToAddress("0x4648a43B2C14Da09FdF82B161150d3F634f40491")
	instance, err := NewZap(address, client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n balanceOf %+v \n\n", instance.BalanceOf)

}
