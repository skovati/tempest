package api

import (
	"io/ioutil"
	"net/http"
	"time"
    "errors"
    "strconv"
	// "encoding/json"
	// "fmt"
	// "reflect"

	"github.com/tidwall/gjson"
)

// base url for actual forecast API call
const foreURLPrefix string = "https://api.weather.gov/gridpoints/"

// used to get weather station from lat/long
// needs "lat,lon" after points/
const locURLPrefix string = "https://api.weather.gov/points/"

// used to get lat/long from address
// NEEDS 
// - ?street=
// AND ONE OF
// 
// - ?city=
// - ?state=
// OR
// - ?zip=
const addrURLPrefix string = "https://geocoding.geo.census.gov/geocoder/locations/address"
const addrURLSuffix string = "&benchmark=Public_AR_Current&format=json"

// GetFore makes the API request and returns a slice of custom Forecast structs
func GetFore(addr Addr) []Forecast {
	// forecast to store API response data
	var forecasts []Forecast

    // get lat and lon from address
    lat, lon, err := getLatLon(addr)
    if err != nil {
        panic(err.Error())
    }

    // and get final api.weather.gov url based on weather station
    finalURL, err := getAPIUrl(lat, lon)
    if err != nil {
        panic(err.Error())
    }

	// get current, next, and next next forecasts
	periods := gjson.GetManyBytes(makeAPIRequest(finalURL), "properties.periods.0", "properties.periods.1", "properties.periods.2")

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


// getLoc returns the correct API Url 
func getAPIUrl(lat float64, long float64) (string, error) {
    url := "https://api.weather.gov/gridpoints/LWX/96,70/forecast"
    return url, nil
}

func getLatLon(addr Addr) (float64, float64, error) {
    // if we don't have street adress, we can't do anything
    if addr.Street == "" {
        return 1000.0, 1000.0, errors.New("Error, street address incorrect")
    }

    // and then we need either city and state or zip code
    if addr.City == "" || addr.State == "" && addr.Zip == 0 {
        return 1000.0, 1000.0, errors.New("Error, need either city and state or zip code in addition to street")
    }

    // otherwise, we have the required info
    var url string = addrURLPrefix
    // if zip is not 0, use that first
    if addr.Zip != 0 {
        url += "?street=" + addr.Street + "&zip=" + strconv.Itoa(addr.Zip) + addrURLSuffix
    }
    resp := makeAPIRequest(url)
    coords := gjson.GetManyBytes(resp, "result.adressMatches.coordinates.x", "result.adressMatches.coordinates.y")
    return coords[0].Float(), coords[1].Float(), nil
}

// makeAPIRequest simply takes the url string passed as a parameter
// and makes an HTPP GET request, and returns the response as a slice of
// bytes for further processing
func makeAPIRequest(url string) []byte {
    // setup custom client that times out instead
    // of hanging if API is down
	var client = &http.Client{Timeout: 10 * time.Second}

	// make API request over HTTPS
	resp, err := client.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// covert to slice of byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
    return body
}
