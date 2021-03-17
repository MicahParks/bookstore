package endpoints

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/models"
	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleCheckout creates a POST /api/checkout endpoint handler via a closure. It can update the Status data for ISBNs.
func HandleCheckout(logger *zap.SugaredLogger, statusStore storage.StatusStore) api.BookCheckoutHandlerFunc {
	return func(params api.BookCheckoutParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"isbns", params.Isbns,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Read the historical statuses of all the affected books.
		statuses, err := statusStore.Read(ctx, params.Isbns)
		if err != nil {

			// Log the error.
			msg := "Failed to read Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookCheckoutDefault{})
		}

		// Check to make sure all books are currently checked in or acquired.
		for _, isbn := range params.Isbns {

			// Confirm the latest status has at least one unavailable.
			if mostRecent(statuses[isbn]).Available == 0 {
				return cantCheckout()
			}
		}

		// Get the current time.
		now := strfmt.DateTime(time.Now())

		// Add checked in to the historical statuses.
		//
		// This isn't the most memory efficient, but that's okay for this coding challenge.
		updatedStatuses := make(map[string]models.History)
		for isbn, status := range statuses {

			// Get the most recent status.
			latestStatus := mostRecent(status)

			// Update the availability.
			available := latestStatus.Available - 1
			unavailable := latestStatus.Unavailable + 1

			// Add to the historical status data.
			status.History = append(status.History, models.Status{
				Available:   available,
				Time:        now,
				Type:        models.StatusTypeCheckout,
				Unavailable: unavailable,
			})

			// Update the new statuses map.
			updatedStatuses[isbn] = status
		}

		// Check in the book.
		if err = statusStore.Write(ctx, updatedStatuses, storage.Upsert); err != nil {

			// Log the error.
			msg := "Failed to write Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookCheckoutDefault{})
		}

		return &api.BookCheckinOK{}
	}
}

// cantCheckout reports to the client that a book can't be checked in if it hasn't been checked in or acquired.
func cantCheckout() middleware.Responder {
	return errorResponse(422, "No available books to check out.", &api.BookCheckoutDefault{})
}
