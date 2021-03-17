package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleDelete creates a DELETE /api/books endpoint handler via a closure. It can delete books from the library.
func HandleDelete(logger *zap.SugaredLogger, bookStore storage.BookStore, statusStore storage.StatusStore) api.BookDeleteHandlerFunc {
	return func(params api.BookDeleteParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"isbns", params.Isbns,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Delete the ISBNs from the StatusStore.
		if err := bookStore.Delete(ctx, params.Isbns); err != nil {

			// Log the error.
			msg := "Failed to delete Book data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookWriteDefault{})
		}

		// Delete the ISBNs from the StatusStore.
		if err := statusStore.Delete(ctx, params.Isbns); err != nil {

			// Log the error.
			msg := "Failed to delete Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookWriteDefault{})
		}

		return &api.BookDeleteOK{}
	}
}
