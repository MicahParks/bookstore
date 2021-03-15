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

// NewBookCheckinParams creates a new BookCheckinParams object
//
// There are no default values defined in the spec.
func NewBookCheckinParams() BookCheckinParams {

	return BookCheckinParams{}
}

// BookCheckinParams contains all the bound params for the book checkin operation
// typically these are obtained from a http.Request
//
// swagger:parameters bookCheckin
type BookCheckinParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ISBNs of the books to check in.
	  In: body
	*/
	Isbns []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBookCheckinParams() beforehand.
func (o *BookCheckinParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
