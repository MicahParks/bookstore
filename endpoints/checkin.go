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

// HandleCheckin creates a POST /api/checkin endpoint handler via a closure. It can read Status data for ISBNs.
func HandleCheckin(logger *zap.SugaredLogger, statusStore storage.StatusStore) api.BookCheckinHandlerFunc {
	return func(params api.BookCheckinParams) middleware.Responder {

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
			// TODO
		}

		// Check to make sure all books are currently checked out.
		for _, isbn := range params.Isbns {

			// Confirm the ISBN has historical statuses.
			history := statuses[isbn].History
			if len(history) > 0 {

				// Confirm the latest status has it checked out.
				if history[len(history)-1].Type != models.StatusTypeCheckout {
					// TODO
				}
			} else {
				// TODO
			}
		}

		// Get the current time.
		now := strfmt.DateTime(time.Now())

		// Add checked in to the historical statuses.
		for _, status := range statuses {
			status.History = append(status.History, models.Status{
				Time: now,
				Type: "",
			})
		}

		// Check in the book.

		return &api.BookCheckinOK{}
	}
}
