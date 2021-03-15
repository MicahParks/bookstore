package endpoints

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleWrite creates and POST /api/{operation}/books endpoint handler via a closure. It can perform write operations on
// Book data.
func HandleWrite(logger *zap.SugaredLogger, bookStore storage.BookStore) api.BookWriteHandlerFunc {
	return func(params api.BookWriteParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"books", params.Books,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Determine the type of write operation.
		var operation storage.WriteOperation
		switch params.Operation {
		case "insert":
			operation = storage.Insert
		case "update":
			operation = storage.Update
		case "upsert":
			operation = storage.Upsert
		}

		// Write the Book data to the BookStore.
		if err := bookStore.Write(ctx, params.Books, operation); err != nil {

			// Log the error.
			logger.Infow("Failed to write Book data.",
				"error", err.Error(),
			)

			// Default the code to a server side error.
			code := 500

			// Check to see if it was a client side error.
			if errors.Is(err, storage.ErrISBNExists) || errors.Is(err, storage.ErrISBNNotFound) {
				code = 400
			}

			// Report the error to the client.
			//
			// Typically don't show internal error message, but this is for speed.
			return errorResponse(code, err.Error(), &api.BookWriteDefault{})
		}

		return &api.BookWriteOK{}
	}
}
