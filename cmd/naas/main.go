package main

import (
	"os"

	"github.wdf.sap.corp/I331555/naasgo/internal/api"
	"github.wdf.sap.corp/I331555/naasgo/internal/cli"
)

func main() {
	port := os.Getenv("NAAS_PORT")
	debug := os.Getenv("NAAS_DEBUG")

	switch port {
	case "":
		// If no port is set, return one reason on the console.
		cli.Run()
	default:
		api.Run(port, debug == "true")
	}
}
