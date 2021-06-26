package api

import (
	"io/ioutil"
	"net/http"
	"time"
	// "encoding/json"
	// "fmt"
	// "reflect"

	"github.com/tidwall/gjson"
)

const testApi string = "https://api.weather.gov/gridpoints/MKX/37,63/forecast"

// makes the API request and returns a slice of custom Forecast structs
func GetFore(lat float64, long float64) []Forecast {
	var client = &http.Client{Timeout: 10 * time.Second}
	// forecast to store API response data
	var forecasts []Forecast

	// make API request over HTTPS
	resp, err := client.Get(testApi)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// covert to slice of byte
	body, err := ioutil.ReadAll(resp.Body)

	// get current, next, and next next forecasts
	periods := gjson.GetManyBytes(body, "properties.periods.0", "properties.periods.1", "properties.periods.2")

	// loop over and format as Forecast struct in memory
	for _, v := range periods {
		forecasts = append(forecasts, Forecast{
			Name: v.Get("name").String(),
			Temp: v.Get("temperature").Float(),
			Unit: v.Get("temperatureUnit").String(),
			Desc: v.Get("detailedForecast").String()})
	}

	// return near future forcasts
	return forecasts
}

func getLoc(lat float64, long float64) {

}
