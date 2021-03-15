package middleware

import (
	"net/http"
	"strings"

	"github.com/MicahParks/bookstore"
)

// FrontendMiddleware is the middleware used to server frontend assets.
func FrontendMiddleware(next http.Handler) (handler http.HandlerFunc) {

	// Create the file server.
	fileServer := http.FileServer(http.FS(bookstore.Frontend))

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
	}
}
