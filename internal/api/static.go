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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" || path == "" {
			path = "/index.html"
		}

		// Remove leading slash for fs.FS
		fsPath := path
		if len(fsPath) > 0 && fsPath[0] == '/' {
			fsPath = fsPath[1:]
		}

		// If the file doesn't exist, serve the error page with 404
		f, err := subFS.Open(fsPath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			r.URL.Path = "/error.html"
		} else {
			_ = f.Close()
		}

		http.FileServer(http.FS(subFS)).ServeHTTP(w, r)
	})
}
