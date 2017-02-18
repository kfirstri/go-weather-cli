package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const urlPath = "http://api.openweathermap.org/data/2.5/weather?units=metric&q=%s&appid=%s"
const appID = "8f704952bd05684b4a6c36b1ccdc417e"

type weatherResponse struct {
	queryCity string
	Name      string
	Weather   []map[string]interface{}
	Main      map[string]float32
}

func getData(city string) (*http.Response, error) {
    return http.Get(fmt.Sprintf(urlPath, url.QueryEscape(city), appID))
}

func (wr *weatherResponse) loadCurrentWeather() (err error) {
    var response *http.Response

	// Getting the data from the openweathermap API
    if response, err = getData(wr.queryCity); err != nil {
		return err
	}

	// Read the json response and unmarshal it to the current weather object
	outputAsBytes, err := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(outputAsBytes, wr); err != nil {
		return err
	}

	return nil
}

func main() {
	var args = os.Args
	var cityWeather weatherResponse

	if len(args) == 1 {
		fmt.Println("Missing city argument")
		return
	}

	// Get the city from the arguments list
	cityWeather.queryCity = args[1]

	err := cityWeather.loadCurrentWeather()

	if err != nil {
		fmt.Println("there was an error", err)
	}

	fmt.Printf("=== Weather in %s ===\n", cityWeather.Name)
	fmt.Printf("%s with tempature of %.2f celsius", cityWeather.Weather[0]["description"], cityWeather.Main["temp"])
}
