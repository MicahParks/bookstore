// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/bookstore/models"
)

// BookCheckinOKCode is the HTTP code returned for type BookCheckinOK
const BookCheckinOKCode int = 200

/*BookCheckinOK The books have been checked in.

swagger:response bookCheckinOK
*/
type BookCheckinOK struct {
}

// NewBookCheckinOK creates BookCheckinOK with default headers values
func NewBookCheckinOK() *BookCheckinOK {

	return &BookCheckinOK{}
}

// WriteResponse to the client
func (o *BookCheckinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*BookCheckinDefault Unexpected error.

swagger:response bookCheckinDefault
*/
type BookCheckinDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBookCheckinDefault creates BookCheckinDefault with default headers values
func NewBookCheckinDefault(code int) *BookCheckinDefault {
	if code <= 0 {
		code = 500
	}

	return &BookCheckinDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the book checkin default response
func (o *BookCheckinDefault) WithStatusCode(code int) *BookCheckinDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the book checkin default response
func (o *BookCheckinDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the book checkin default response
func (o *BookCheckinDefault) WithPayload(payload *models.Error) *BookCheckinDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book checkin default response
func (o *BookCheckinDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookCheckinDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
