package cmd

import (
    "fmt"
    "strconv"

    "github.com/skovati/tmpst/api"
    "github.com/skovati/tmpst/ui"
)

func Run(args []string) {
    // convert cli args to floats
    lat, err := strconv.ParseFloat(args[1], 64)
    if err != nil {
        panic(err.Error)
    }
    long, err := strconv.ParseFloat(args[2], 64)
    if err != nil {
        panic(err.Error)
    }

    ui.PrintLine()
    fmt.Println("welcome to tmpst")
    ui.PrintLine()

    // call api
    forecasts := api.GetFore(lat, long)

    for _, v := range forecasts {
        ui.PrintLine()
        fmt.Println(v.String())
    }
}
