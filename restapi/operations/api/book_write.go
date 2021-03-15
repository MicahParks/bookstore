// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// BookWriteHandlerFunc turns a function with the right signature into a book write handler
type BookWriteHandlerFunc func(BookWriteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookWriteHandlerFunc) Handle(params BookWriteParams) middleware.Responder {
	return fn(params)
}

// BookWriteHandler interface for that can handle valid book write params
type BookWriteHandler interface {
	Handle(BookWriteParams) middleware.Responder
}

// NewBookWrite creates a new http.Handler for the book write operation
func NewBookWrite(ctx *middleware.Context, handler BookWriteHandler) *BookWrite {
	return &BookWrite{Context: ctx, Handler: handler}
}

/* BookWrite swagger:route POST /api/books/{operation} api bookWrite

Insert, update, or upsert books to the library.

*/
type BookWrite struct {
	Context *middleware.Context
	Handler BookWriteHandler
}

func (o *BookWrite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookWriteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
