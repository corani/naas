package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.wdf.sap.corp/I331555/naasgo/cfg"
)

func NewRouter(getter func() string, debug bool) http.Handler {
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
	router.Get("/no", noHandler(getter))

	return router
}

func Run(getter func() string, port string, debug bool) {
	router := NewRouter(getter, debug)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("listening on: http://localhost:%s/", port)

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("HTTP server error: %v", err)
			}
		}

		log.Println("server closed")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	log.Println("HTTP shutdown completed")
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

func noHandler(getter func() string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		reason := struct {
			Reason string `json:"reason"`
		}{
			Reason: getter(),
		}

		if err := json.NewEncoder(w).Encode(reason); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
