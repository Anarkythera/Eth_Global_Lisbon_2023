package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
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
	e.GET("/ViewConfig", ViewConfig)
	e.POST("/EditConfig", EditConfig)
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
	testnetURL := "https://api-testnet.polygonscan.com/"
	rest := "api?module=account&action=balance&address=0x3092ef862A180D0f44C5E537EfE05Cd7DCbB28A7&apikey="
	apiKey := ""

	queryURL := fmt.Sprintf("%s%s%s", testnetURL, rest, apiKey)

	fmt.Printf(queryURL)
	req, _ := http.NewRequest("POST", queryURL, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))
}
