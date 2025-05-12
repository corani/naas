package main

import (
	"os"

	"github.wdf.sap.corp/I331555/naasgo/internal/api"
	"github.wdf.sap.corp/I331555/naasgo/internal/cli"
	"github.wdf.sap.corp/I331555/naasgo/internal/reasons"
)

func main() {
	port := os.Getenv("NAAS_PORT")
	debug := os.Getenv("NAAS_DEBUG")
	getter := reasons.Get

	switch port {
	case "":
		// If no port is set, return one reason on the console.
		cli.Run(getter)
	default:
		api.Run(getter, port, debug == "true")
	}
}
