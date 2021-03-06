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
			"books", params.BookQuantities,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Determine the type of write operation.
		var operation storage.WriteOperation
		var err error
		if operation, err = operation.FromString(params.Operation); err != nil {

			// Log the error.
			msg := "Failed to convert WriteOperation to enum. goswagger should have prevented this."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(422, msg, &api.BookWriteDefault{})
		}

		// Create the slice of books.
		books := make([]models.Book, len(params.BookQuantities))
		index := 0
		for _, bookQuantity := range params.BookQuantities {
			books[index] = bookQuantity.Book
			index++
		}

		// Write the Book data to the BookStore.
		if err = bookStore.Write(ctx, books, operation); err != nil {

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
			return errorResponse(code, msg, &api.BookWriteDefault{})
		}

		// Get the current system time.
		now := strfmt.DateTime(time.Now())

		// Create the map of status.
		statuses := make(map[string]models.History, len(params.BookQuantities))

		// Iterate through the given books, create their historical statuses.
		for _, bookQuantity := range params.BookQuantities {
			statuses[bookQuantity.Book.ISBN] = models.History{
				History: []models.Status{{
					Available: bookQuantity.Quantity,
					Time:      now,
					Type:      models.StatusTypeAcquired,
				}},
				Isbn: bookQuantity.Book.ISBN,
			}
		}

		// Write the book statuses to the StatusStore.
		//
		// Write operations other than storage.Upsert are possible.
		if err = statusStore.Write(ctx, statuses, storage.Upsert); err != nil {

			// Log the error.
			msg := "Failed to update the statuses for written books."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookWriteDefault{})
		}

		return &api.BookWriteOK{}
	}
}
