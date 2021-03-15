package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/models"
	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleStatus creates a POST /api/status endpoint handler via a closure. It can perform read operations on Status
// data.
func HandleStatus(logger *zap.SugaredLogger, statusStore storage.StatusStore) api.BookStatusHandlerFunc {
	return func(params api.BookStatusParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"isbns", params.Isbns,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Read the historical status data from the bookstore.
		statuses, err := statusStore.Read(ctx, params.Isbns)
		if err != nil {

			// Log the error.
			msg := "Failed to read Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			//
			// Typically don't show internal error message, but this is for speed.
			return errorResponse(500, msg+": "+err.Error(), &api.BookStatusDefault{})
		}

		// Get the latest statuses.
		latest := make(map[string]models.Status, len(statuses))
		for _, status := range statuses {
			latest[status.Isbn] = status.History[len(status.History)-1]
		}

		return &api.BookStatusOK{
			Payload: latest,
		}
	}
}
