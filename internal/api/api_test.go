package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestVersionRoute(t *testing.T) {
	getter := func() string { return "test-reason" }
	router := NewRouter(getter, false)

	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected status 200")

	var resp struct {
		Version string `json:"version"`
		Hash    string `json:"hash"`
		Build   string `json:"build"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp),
		"failed to unmarshal response")
}

func TestNoRoute(t *testing.T) {
	getter := func() string { return "test-reason" }
	router := NewRouter(getter, false)

	req := httptest.NewRequest(http.MethodGet, "/no", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected status 200")

	var resp struct {
		Reason string `json:"reason"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp),
		"failed to unmarshal response")
	require.Equal(t, "test-reason", resp.Reason,
		"expected reason 'test-reason'")
}

func TestDebugProfilerRoute(t *testing.T) {
	getter := func() string { return "test-reason" }
	router := NewRouter(getter, true)

	req := httptest.NewRequest(http.MethodGet, "/debug/pprof/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code,
		"expected status 200")
	require.NotEmpty(t, w.Body.Bytes(),
		"expected non-empty response body for /debug/pprof/")
}

func TestDebugProfilerRoute_Disabled(t *testing.T) {
	getter := func() string { return "test-reason" }
	router := NewRouter(getter, false)

	req := httptest.NewRequest("GET", "/debug/pprof/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.NotEqual(t, http.StatusOK, w.Code,
		"expected non-200 status when debug is false")
}

func TestTimeoutMiddleware(t *testing.T) {
	getter := func() string { return "test-reason" }
	router := NewRouter(getter, false)

	// Add a slow route for testing
	router.Method("GET", "/slow",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-time.After(1500 * time.Millisecond):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"ok":true}`)) //nolint:errcheck
			case <-r.Context().Done():
				// Context cancelled by timeout middleware
			}
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/slow", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusGatewayTimeout, w.Code,
		"expected status 504 due to timeout")
}
