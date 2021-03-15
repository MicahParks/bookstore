package endpoints

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/bookstore/restapi/operations/api"
	"github.com/MicahParks/bookstore/storage"
)

// HandleRead creates a POST /api/books endpoint handler via a closure. It can perform read operations on Book data.
func HandleRead(logger *zap.SugaredLogger, bookStore storage.BookStore, statusStore storage.StatusStore) api.BookReadHandlerFunc {
	return func(params api.BookReadParams) middleware.Responder {

		// Debug info.
		logger.Debugw("",
			"isbns", params.Isbns,
		)

		// Create a context for the request.
		ctx, cancel := defaultCtx()
		defer cancel()

		// Read the books from the BookStore.
		books,err := bookStore.Read(ctx, params.Isbns)
		if err != nil {
			// TODO.
		}


		return &api.BookReadOK{Payload: }
	}
}
