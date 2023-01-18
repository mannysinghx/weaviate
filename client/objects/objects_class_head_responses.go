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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsClassHeadReader is a Reader for the ObjectsClassHead structure.
type ObjectsClassHeadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsClassHeadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewObjectsClassHeadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObjectsClassHeadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsClassHeadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsClassHeadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsClassHeadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsClassHeadNoContent creates a ObjectsClassHeadNoContent with default headers values
func NewObjectsClassHeadNoContent() *ObjectsClassHeadNoContent {
	return &ObjectsClassHeadNoContent{}
}

/*
ObjectsClassHeadNoContent handles this case with default header values.

Object exists.
*/
type ObjectsClassHeadNoContent struct {
}

func (o *ObjectsClassHeadNoContent) Error() string {
	return fmt.Sprintf("[HEAD /objects/{className}/{id}][%d] objectsClassHeadNoContent ", 204)
}

func (o *ObjectsClassHeadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassHeadUnauthorized creates a ObjectsClassHeadUnauthorized with default headers values
func NewObjectsClassHeadUnauthorized() *ObjectsClassHeadUnauthorized {
	return &ObjectsClassHeadUnauthorized{}
}

/*
ObjectsClassHeadUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsClassHeadUnauthorized struct {
}

func (o *ObjectsClassHeadUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /objects/{className}/{id}][%d] objectsClassHeadUnauthorized ", 401)
}

func (o *ObjectsClassHeadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassHeadForbidden creates a ObjectsClassHeadForbidden with default headers values
func NewObjectsClassHeadForbidden() *ObjectsClassHeadForbidden {
	return &ObjectsClassHeadForbidden{}
}

/*
ObjectsClassHeadForbidden handles this case with default header values.

Forbidden
*/
type ObjectsClassHeadForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassHeadForbidden) Error() string {
	return fmt.Sprintf("[HEAD /objects/{className}/{id}][%d] objectsClassHeadForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsClassHeadForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassHeadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassHeadNotFound creates a ObjectsClassHeadNotFound with default headers values
func NewObjectsClassHeadNotFound() *ObjectsClassHeadNotFound {
	return &ObjectsClassHeadNotFound{}
}

/*
ObjectsClassHeadNotFound handles this case with default header values.

Object doesn't exist.
*/
type ObjectsClassHeadNotFound struct {
}

func (o *ObjectsClassHeadNotFound) Error() string {
	return fmt.Sprintf("[HEAD /objects/{className}/{id}][%d] objectsClassHeadNotFound ", 404)
}

func (o *ObjectsClassHeadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassHeadInternalServerError creates a ObjectsClassHeadInternalServerError with default headers values
func NewObjectsClassHeadInternalServerError() *ObjectsClassHeadInternalServerError {
	return &ObjectsClassHeadInternalServerError{}
}

/*
ObjectsClassHeadInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsClassHeadInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassHeadInternalServerError) Error() string {
	return fmt.Sprintf("[HEAD /objects/{className}/{id}][%d] objectsClassHeadInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsClassHeadInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassHeadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
