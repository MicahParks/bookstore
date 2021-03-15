// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/MicahParks/bookstore/endpoints"
	"github.com/MicahParks/bookstore/restapi/operations"
	"github.com/MicahParks/bookstore/storage"
)

//go:generate swagger generate server --target ../../bookstore --name Bookstore --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.BookstoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BookstoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Create a zap logger.
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create zap logger.\nError: %s", err.Error())
	}
	zapLogger.Info("Logger created. Configuring.")
	logger := zapLogger.Sugar()

	// Create the Book data storage.
	bookStore := storage.NewMemBook()

	// Create the Status data storage.
	statusStore := storage.NewMemStatus()

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Assign the endpoint handlers.
	api.SystemAliveHandler = endpoints.HandleAlive()
	api.APIBookDeleteHandler = endpoints.HandleDelete(logger.Named("DELETE /api/books"), bookStore, statusStore)
	api.APIBookReadHandler = endpoints.HandleRead(logger.Named("POST /api/books"), bookStore)
	api.APIBookWriteHandler = endpoints.HandleWrite(logger.Named("POST /api/books/{operation}"), bookStore, statusStore)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
