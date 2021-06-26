package api

import (
	"strconv"
	"strings"
)

// Forecast models the important info from the API json reponse
type Forecast struct {
	Name string
	Temp float64
	Unit string
	Desc string
}

// returns a string representation of a Forecast
func (f Forecast) String() string {
	return strings.ToLower(f.Name) + ":\n\ntemp: " + strconv.FormatFloat(f.Temp, 'f', 1, 64) + f.Unit + "\nforecast: " + f.Desc
}
