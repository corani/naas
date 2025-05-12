package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.wdf.sap.corp/I331555/naasgo/internal"
)

func main() {
	port := os.Getenv("NAAS_PORT")

	// If no port is set, return one reason on the console.
	if port == "" {
		fmt.Println(internal.GetReason())

		return
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(httprate.LimitByIP(100, time.Minute))

	// `/ping`: health check endpoint
	router.Use(middleware.Heartbeat("/ping"))

	if os.Getenv("NAAS_DEBUG") == "true" {
		// `/debug/pprof`: pprof profiler
		router.Mount("/debug", middleware.Profiler())
	}

	// `/no`: returns a random reason
	router.Get("/no", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		reason := struct {
			Reason string `json:"reason"`
		}{
			Reason: internal.GetReason(),
		}

		if err := json.NewEncoder(w).Encode(reason); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":"+port, router)
}
