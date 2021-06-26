package api

import (
    "strconv"
    "strings"
)

type Forecast struct {
    Name string
    Temp float64
    Unit string
    Desc string
}

func (f Forecast) String() string {
    return strings.ToLower(f.Name) + ":\n\ntemp: " + strconv.FormatFloat(f.Temp, 'f', 1, 64) + f.Unit + "\nforecast: " + f.Desc
}
