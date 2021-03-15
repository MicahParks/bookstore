// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// BookHistoryHandlerFunc turns a function with the right signature into a book history handler
type BookHistoryHandlerFunc func(BookHistoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookHistoryHandlerFunc) Handle(params BookHistoryParams) middleware.Responder {
	return fn(params)
}

// BookHistoryHandler interface for that can handle valid book history params
type BookHistoryHandler interface {
	Handle(BookHistoryParams) middleware.Responder
}

// NewBookHistory creates a new http.Handler for the book history operation
func NewBookHistory(ctx *middleware.Context, handler BookHistoryHandler) *BookHistory {
	return &BookHistory{Context: ctx, Handler: handler}
}

/* BookHistory swagger:route POST /api/history api bookHistory

Get the history for the given books.

*/
type BookHistory struct {
	Context *middleware.Context
	Handler BookHistoryHandler
}

func (o *BookHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookHistoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
