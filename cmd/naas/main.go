package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/corani/naas/cfg"
	"github.com/corani/naas/internal/api"
	"github.com/corani/naas/internal/cli"
	"github.com/corani/naas/internal/reasons"
)

func main() {
	var (
		serveFlag   = flag.Bool("serve", false, "Serve the API on port NAAS_PORT (from environment)")
		versionFlag = flag.Bool("version", false, "Print version information")
		helpFlag    = flag.Bool("help", false, "Show usage instructions")
	)
	flag.Parse()

	if *helpFlag {
		flag.Usage()

		return
	}

	if *versionFlag {
		//nolint:errcheck
		fmt.Printf("Version: %s\nBuild: %s\nHash: %s\n", cfg.Version(), cfg.Build(), cfg.Hash())

		return
	}

	getter := reasons.Get
	port := os.Getenv("NAAS_PORT")
	debug := os.Getenv("NAAS_DEBUG")

	if *serveFlag {
		if port == "" {
			port = "3939"
		}
		api.Run(getter, port, debug == "true")
	} else {
		cli.Run(os.Stdout, getter)
	}
}
