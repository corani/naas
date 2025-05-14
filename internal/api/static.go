package api

import (
	"io/fs"
	"net/http"

	"github.com/corani/naas/static"
)

// StaticFileServer returns a handler that serves static files from the embedded FS.
func StaticFileServer() http.Handler {
	subFS, err := fs.Sub(static.FS, "www")
	if err != nil {
		panic("failed to create sub FS for static files: " + err.Error())
	}

	return http.FileServer(http.FS(subFS))
}
