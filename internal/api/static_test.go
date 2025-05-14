package api

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/corani/naas/static"
	"github.com/stretchr/testify/require"
)

func TestStaticFileServer_Index(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rw := httptest.NewRecorder()

	h := StaticFileServer()
	h.ServeHTTP(rw, req)

	require.Equal(t, http.StatusOK, rw.Code,
		"expected 200 OK")
	require.Contains(t, rw.Body.String(), "Powered by",
		"expected index.html content")
}

func TestStaticFileServer_ExistingFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	rw := httptest.NewRecorder()

	h := StaticFileServer()
	h.ServeHTTP(rw, req)

	// Check that we're redirected to /
	require.Equal(t, http.StatusMovedPermanently, rw.Code,
		"expected status 301")
	require.Equal(t, "./", rw.Header().Get("Location"),
		"expected Location header to be ./")
}

func TestStaticFileServer_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/does-not-exist.html", nil)
	rw := httptest.NewRecorder()

	h := StaticFileServer()
	h.ServeHTTP(rw, req)

	require.Equal(t, http.StatusNotFound, rw.Code,
		"expected 404 Not Found, got %d", rw.Code)
	require.Contains(t, rw.Body.String(), "Powered by",
		"expected error.html content, got: %s", rw.Body.String())
}

func TestStaticFileServer_ErrorHtmlExists(t *testing.T) {
	subFS, err := fs.Sub(static.FS, "www")
	require.NoError(t, err,
		"failed to create sub FS")

	_, err = subFS.Open("error.html")
	require.NoError(t, err,
		"error.html should exist in static FS")
}
