package middleware

import (
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/bookstore"
)

// FrontendMiddleware is the middleware used to server frontend assets.
func FrontendMiddleware(next http.Handler) (handler http.HandlerFunc, err error) {

	// Check the system environment to see if an OS directory or embedded directory should be used.
	var fileSystem fs.FS
	if osDir := os.Getenv("FRONTEND_DIR"); osDir != "" {

		// Use the given directory as a file system.
		fileSystem = os.DirFS(osDir)
	} else {

		// Use the embedded frontend.
		if fileSystem, err = fs.Sub(bookstore.Frontend, bookstore.FrontendSubDir); err != nil {
			return nil, err
		}
	}

	// Create the file server.
	fileServer := http.FileServer(http.FS(fileSystem))

	// Create the HTTP handler via a closure.
	return func(writer http.ResponseWriter, request *http.Request) {

		// If the /api prefix is seen, follow the middleware pattern.
		path := request.URL.Path
		if strings.HasPrefix(path, "/api") || path == "/swagger.json" || path == "/docs" {
			next.ServeHTTP(writer, request)
		} else {

			// Serve from the embedded file system.
			fileServer.ServeHTTP(writer, request)
		}
	}, nil
}
