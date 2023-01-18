//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsClassGetOKCode is the HTTP code returned for type ObjectsClassGetOK
const ObjectsClassGetOKCode int = 200

/*
ObjectsClassGetOK Successful response.

swagger:response objectsClassGetOK
*/
type ObjectsClassGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Object `json:"body,omitempty"`
}

// NewObjectsClassGetOK creates ObjectsClassGetOK with default headers values
func NewObjectsClassGetOK() *ObjectsClassGetOK {

	return &ObjectsClassGetOK{}
}

// WithPayload adds the payload to the objects class get o k response
func (o *ObjectsClassGetOK) WithPayload(payload *models.Object) *ObjectsClassGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class get o k response
func (o *ObjectsClassGetOK) SetPayload(payload *models.Object) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassGetBadRequestCode is the HTTP code returned for type ObjectsClassGetBadRequest
const ObjectsClassGetBadRequestCode int = 400

/*
ObjectsClassGetBadRequest Malformed request.

swagger:response objectsClassGetBadRequest
*/
type ObjectsClassGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassGetBadRequest creates ObjectsClassGetBadRequest with default headers values
func NewObjectsClassGetBadRequest() *ObjectsClassGetBadRequest {

	return &ObjectsClassGetBadRequest{}
}

// WithPayload adds the payload to the objects class get bad request response
func (o *ObjectsClassGetBadRequest) WithPayload(payload *models.ErrorResponse) *ObjectsClassGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class get bad request response
func (o *ObjectsClassGetBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassGetUnauthorizedCode is the HTTP code returned for type ObjectsClassGetUnauthorized
const ObjectsClassGetUnauthorizedCode int = 401

/*
ObjectsClassGetUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassGetUnauthorized
*/
type ObjectsClassGetUnauthorized struct {
}

// NewObjectsClassGetUnauthorized creates ObjectsClassGetUnauthorized with default headers values
func NewObjectsClassGetUnauthorized() *ObjectsClassGetUnauthorized {

	return &ObjectsClassGetUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassGetForbiddenCode is the HTTP code returned for type ObjectsClassGetForbidden
const ObjectsClassGetForbiddenCode int = 403

/*
ObjectsClassGetForbidden Forbidden

swagger:response objectsClassGetForbidden
*/
type ObjectsClassGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassGetForbidden creates ObjectsClassGetForbidden with default headers values
func NewObjectsClassGetForbidden() *ObjectsClassGetForbidden {

	return &ObjectsClassGetForbidden{}
}

// WithPayload adds the payload to the objects class get forbidden response
func (o *ObjectsClassGetForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class get forbidden response
func (o *ObjectsClassGetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassGetNotFoundCode is the HTTP code returned for type ObjectsClassGetNotFound
const ObjectsClassGetNotFoundCode int = 404

/*
ObjectsClassGetNotFound Successful query result but no resource was found.

swagger:response objectsClassGetNotFound
*/
type ObjectsClassGetNotFound struct {
}

// NewObjectsClassGetNotFound creates ObjectsClassGetNotFound with default headers values
func NewObjectsClassGetNotFound() *ObjectsClassGetNotFound {

	return &ObjectsClassGetNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassGetInternalServerErrorCode is the HTTP code returned for type ObjectsClassGetInternalServerError
const ObjectsClassGetInternalServerErrorCode int = 500

/*
ObjectsClassGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassGetInternalServerError
*/
type ObjectsClassGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassGetInternalServerError creates ObjectsClassGetInternalServerError with default headers values
func NewObjectsClassGetInternalServerError() *ObjectsClassGetInternalServerError {

	return &ObjectsClassGetInternalServerError{}
}

// WithPayload adds the payload to the objects class get internal server error response
func (o *ObjectsClassGetInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class get internal server error response
func (o *ObjectsClassGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
