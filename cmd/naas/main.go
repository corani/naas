package main

import (
	"os"

	"github.com/corani/naas/internal/api"
	"github.com/corani/naas/internal/cli"
	"github.com/corani/naas/internal/reasons"
)

func main() {
	port := os.Getenv("NAAS_PORT")
	debug := os.Getenv("NAAS_DEBUG")
	getter := reasons.Get

	switch port {
	case "":
		// If no port is set, return one reason on the console.
		cli.Run(os.Stdout, getter)
	default:
		api.Run(getter, port, debug == "true")
	}
}
