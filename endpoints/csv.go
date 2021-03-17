package endpoints

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleCSV creates a POST /api/csv endpoint handler via a closure. It creates a CSV file for all Book data and current
// Status data.
func HandleCSV(logger *zap.SugaredLogger, bookStore storage.BookStore, statusStore storage.StatusStore) api.BookCSVHandlerFunc {
	return func(params api.BookCSVParams) middleware.Responder {

		// Debug info.
		logger.Debug("")

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Read the historical statuses of all the affected books.
		statuses, err := statusStore.Read(ctx, nil)
		if err != nil {

			// Log the error.
			msg := "Failed to read Status data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookCSVDefault{})
		}

		// Read the books from the BookStore.
		books, err := bookStore.Read(ctx, nil)
		if err != nil {

			// Log the error.
			msg := "Failed to read Book data."
			logger.Infow(msg,
				"error", err.Error(),
			)

			// Report the error to the client.
			return errorResponse(500, msg, &api.BookCSVDefault{})
		}

		// Create the CSV writer.
		buf := bytes.NewBuffer(nil)
		writer := csv.NewWriter(buf)

		// Write the CSV header.
		if err = writer.Write([]string{"ISBN", "Available", "Unavailable", "Title", "Author", "Description"}); err != nil {
			return cantCSV()
		}

		// Iterate through all Book data.
		for _, book := range books {

			// Get the current status for the book.
			status := mostRecent(statuses[book.ISBN])
			available := strconv.FormatUint(status.Available, 10)
			unavailable := strconv.FormatUint(status.Unavailable, 10)

			// Write the CSV record.
			if err = writer.Write([]string{book.ISBN, available, unavailable, book.Title, book.Author, book.Description}); err != nil {
				return cantCSV()
			}
		}

		// Flush to the in memory buffer.
		writer.Flush()

		return &api.BookCSVOK{
			Payload: ioutil.NopCloser(buf),
		}
	}
}

// cantCSV reports to the client that a CSV export could not be created.
func cantCSV() middleware.Responder {
	return errorResponse(500, "A CSV export could not be created.", &api.BookCheckoutDefault{})
}
