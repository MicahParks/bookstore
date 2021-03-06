// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewBookReadParams creates a new BookReadParams object
//
// There are no default values defined in the spec.
func NewBookReadParams() BookReadParams {

	return BookReadParams{}
}

// BookReadParams contains all the bound params for the book read operation
// typically these are obtained from a http.Request
//
// swagger:parameters bookRead
type BookReadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ISBNs of the books whose data is to be read.
	  In: body
	*/
	Isbns []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBookReadParams() beforehand.
func (o *BookReadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("isbns", "body", "", err))
		} else {
			// no validation required on inline body
			o.Isbns = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
