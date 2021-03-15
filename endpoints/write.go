package endpoints

import (
	"errors"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/models"
	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleWrite creates and POST /api/{operation}/books endpoint handler via a closure. It can perform write operations
// on Book data.
func HandleWrite(logger *zap.SugaredLogger, bookStore storage.BookStore, statusStore storage.StatusStore) api.BookWriteHandlerFunc {
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
			msg := "Failed to write Book data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Default the code to a server side error.
			code := 500

			// Check to see if it was a client side error.
			if errors.Is(err, storage.ErrISBNExists) || errors.Is(err, storage.ErrISBNNotFound) {
				code = 422
			}

			// Report the error to the client.
			//
			// Typically don't show internal error message, but this is for speed.
			return errorResponse(code, msg+": "+err.Error(), &api.BookWriteDefault{})
		}

		// Get the current system time.
		now := strfmt.DateTime(time.Now())

		// Create the map of status.
		statuses := make(map[string]models.History, len(params.Books))

		// Iterate through the given books, create their historical statuses.
		for _, book := range params.Books {
			statuses[book.ISBN] = models.History{
				History: []models.Status{{
					Time: now,
					Type: models.StatusTypeAcquired,
				}},
				Isbn: book.ISBN,
			}
		}

		// Write the book statuses to the StatusStore.
		if err := statusStore.Write(ctx, statuses, storage.Upsert); err != nil { // TODO Not always an upsert...

			// Log the error.
			msg := "Failed to update the statuses for written books."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			//
			// Typically don't show internal error message, but this is for speed.
			return errorResponse(500, msg+": "+err.Error(), &api.BookWriteDefault{})
		}

		return &api.BookWriteOK{}
	}
}
