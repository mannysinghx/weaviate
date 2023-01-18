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

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// NewBatchReferencesCreateParams creates a new BatchReferencesCreateParams object
// with the default values initialized.
func NewBatchReferencesCreateParams() *BatchReferencesCreateParams {
	var ()
	return &BatchReferencesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchReferencesCreateParamsWithTimeout creates a new BatchReferencesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchReferencesCreateParamsWithTimeout(timeout time.Duration) *BatchReferencesCreateParams {
	var ()
	return &BatchReferencesCreateParams{

		timeout: timeout,
	}
}

// NewBatchReferencesCreateParamsWithContext creates a new BatchReferencesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchReferencesCreateParamsWithContext(ctx context.Context) *BatchReferencesCreateParams {
	var ()
	return &BatchReferencesCreateParams{

		Context: ctx,
	}
}

// NewBatchReferencesCreateParamsWithHTTPClient creates a new BatchReferencesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchReferencesCreateParamsWithHTTPClient(client *http.Client) *BatchReferencesCreateParams {
	var ()
	return &BatchReferencesCreateParams{
		HTTPClient: client,
	}
}

/*
BatchReferencesCreateParams contains all the parameters to send to the API endpoint
for the batch references create operation typically these are written to a http.Request
*/
type BatchReferencesCreateParams struct {

	/*Body
	  A list of references to be batched. The ideal size depends on the used database connector. Please see the documentation of the used connector for help

	*/
	Body []*models.BatchReference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batch references create params
func (o *BatchReferencesCreateParams) WithTimeout(timeout time.Duration) *BatchReferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch references create params
func (o *BatchReferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch references create params
func (o *BatchReferencesCreateParams) WithContext(ctx context.Context) *BatchReferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch references create params
func (o *BatchReferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch references create params
func (o *BatchReferencesCreateParams) WithHTTPClient(client *http.Client) *BatchReferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch references create params
func (o *BatchReferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch references create params
func (o *BatchReferencesCreateParams) WithBody(body []*models.BatchReference) *BatchReferencesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch references create params
func (o *BatchReferencesCreateParams) SetBody(body []*models.BatchReference) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchReferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
