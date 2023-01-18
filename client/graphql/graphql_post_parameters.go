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

package graphql

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

// NewGraphqlPostParams creates a new GraphqlPostParams object
// with the default values initialized.
func NewGraphqlPostParams() *GraphqlPostParams {
	var ()
	return &GraphqlPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGraphqlPostParamsWithTimeout creates a new GraphqlPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGraphqlPostParamsWithTimeout(timeout time.Duration) *GraphqlPostParams {
	var ()
	return &GraphqlPostParams{

		timeout: timeout,
	}
}

// NewGraphqlPostParamsWithContext creates a new GraphqlPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGraphqlPostParamsWithContext(ctx context.Context) *GraphqlPostParams {
	var ()
	return &GraphqlPostParams{

		Context: ctx,
	}
}

// NewGraphqlPostParamsWithHTTPClient creates a new GraphqlPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGraphqlPostParamsWithHTTPClient(client *http.Client) *GraphqlPostParams {
	var ()
	return &GraphqlPostParams{
		HTTPClient: client,
	}
}

/*
GraphqlPostParams contains all the parameters to send to the API endpoint
for the graphql post operation typically these are written to a http.Request
*/
type GraphqlPostParams struct {

	/*Body
	  The GraphQL query request parameters.

	*/
	Body *models.GraphQLQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the graphql post params
func (o *GraphqlPostParams) WithTimeout(timeout time.Duration) *GraphqlPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the graphql post params
func (o *GraphqlPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the graphql post params
func (o *GraphqlPostParams) WithContext(ctx context.Context) *GraphqlPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the graphql post params
func (o *GraphqlPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the graphql post params
func (o *GraphqlPostParams) WithHTTPClient(client *http.Client) *GraphqlPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the graphql post params
func (o *GraphqlPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the graphql post params
func (o *GraphqlPostParams) WithBody(body *models.GraphQLQuery) *GraphqlPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the graphql post params
func (o *GraphqlPostParams) SetBody(body *models.GraphQLQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GraphqlPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
