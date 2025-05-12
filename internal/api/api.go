package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.wdf.sap.corp/I331555/naasgo/cfg"
	"github.wdf.sap.corp/I331555/naasgo/internal/reasons"
)

func Run(port string, debug bool) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(httprate.LimitByIP(100, time.Minute))

	// `/ping`: health check endpoint
	router.Use(middleware.Heartbeat("/ping"))

	if debug {
		// `/debug/pprof`: pprof profiler
		router.Mount("/debug", middleware.Profiler())
	}

	// `/version`: returns the version of the service
	router.Get("/version", versionHandler)

	// `/no`: returns a random reason
	router.Get("/no", noHandler)

	http.ListenAndServe(":"+port, router)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	version := struct {
		Version string `json:"version"`
		Hash    string `json:"hash"`
		Build   string `json:"build"`
	}{
		Version: cfg.Version(),
		Hash:    cfg.Hash(),
		Build:   cfg.Build(),
	}

	if err := json.NewEncoder(w).Encode(version); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func noHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reason := struct {
		Reason string `json:"reason"`
	}{
		Reason: reasons.Get(),
	}

	if err := json.NewEncoder(w).Encode(reason); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
