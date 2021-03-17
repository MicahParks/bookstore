package endpoints

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MicahParks/bookstore/models"
)

const (

	// defaultTimeout is the default request context timeout.
	defaultTimeout = time.Second
)

// defaultResponse is an interface used to pass different types of default responses and return an error responder.
type defaultResponse interface {
	SetStatusCode(code int)
	SetPayload(payload *models.Error)
	WriteResponse(rw http.ResponseWriter, producer runtime.Producer)
}

// defaultCtx creates a new context and cancel function with the default timeout.
func defaultCtx() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}

// errorResponse creates a response given the required assets.
func errorResponse(code int, message string, resp defaultResponse) middleware.Responder {

	// Set the payload for the response.
	resp.SetPayload(&models.Error{
		Code:    int64(code),
		Message: message,
	})

	// Set the status code for the response.
	resp.SetStatusCode(code)

	return resp
}

// mostRecent gets the most recent status from the historical Status data.
func mostRecent(history models.History) (status models.Status) {
	return history.History[len(history.History)-1]
}
