// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/bookstore/models"
)

// BookWriteOKCode is the HTTP code returned for type BookWriteOK
const BookWriteOKCode int = 200

/*BookWriteOK The books have been written to the library.

swagger:response bookWriteOK
*/
type BookWriteOK struct {
}

// NewBookWriteOK creates BookWriteOK with default headers values
func NewBookWriteOK() *BookWriteOK {

	return &BookWriteOK{}
}

// WriteResponse to the client
func (o *BookWriteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*BookWriteDefault Unexpected error.

swagger:response bookWriteDefault
*/
type BookWriteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBookWriteDefault creates BookWriteDefault with default headers values
func NewBookWriteDefault(code int) *BookWriteDefault {
	if code <= 0 {
		code = 500
	}

	return &BookWriteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the book write default response
func (o *BookWriteDefault) WithStatusCode(code int) *BookWriteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the book write default response
func (o *BookWriteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the book write default response
func (o *BookWriteDefault) WithPayload(payload *models.Error) *BookWriteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book write default response
func (o *BookWriteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookWriteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
