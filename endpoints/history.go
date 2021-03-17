package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleHistory creates a POST /api/history endpoint handler via a closure. It can perform read operations on Status
// data.
func HandleHistory(logger *zap.SugaredLogger, statusStore storage.StatusStore) api.BookHistoryHandlerFunc {
	return func(params api.BookHistoryParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"isbns", params.Isbns,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Read the historical status data from the StatusStore.
		statuses, err := statusStore.Read(ctx, params.Isbns)
		if err != nil {

			// Log the error.
			msg := "Failed to read Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookHistoryDefault{})
		}

		return &api.BookHistoryOK{
			Payload: statuses,
		}
	}
}
