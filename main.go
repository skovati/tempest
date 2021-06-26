package main

import (
    "os"

    "github.com/skovati/tmpst/cmd"
)

func main() {
    // run cmd package with command line args
    cmd.Run(os.Args)
}
