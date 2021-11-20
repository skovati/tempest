package cmd

import (
	"fmt"
	// "strconv"

	"github.com/skovati/tmpst/api"
	"github.com/skovati/tmpst/ui"
)

// Run handles the main application logic
func Run(args []string) {
	ui.PrintLine()
	fmt.Println("welcome to tmpst")
	fmt.Println("grabbing forecast data...")
	ui.PrintLine()
    fmt.Println()

    // make addr struct
    addr := api.Addr{
        Street: "",
        City: "",
        State: "",
    }

	// call api
	forecasts := api.GetFore(addr)

	for _, v := range forecasts {
		ui.PrintLine()
		fmt.Println(v.String())
	}
}
