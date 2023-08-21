// Code generated by go-swagger; DO NOT EDIT.

package config

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
)

// NewFindConfigLargeMemoryAllocationWarningThresholdParams creates a new FindConfigLargeMemoryAllocationWarningThresholdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigLargeMemoryAllocationWarningThresholdParams() *FindConfigLargeMemoryAllocationWarningThresholdParams {
	return &FindConfigLargeMemoryAllocationWarningThresholdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithTimeout creates a new FindConfigLargeMemoryAllocationWarningThresholdParams object
// with the ability to set a timeout on a request.
func NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithTimeout(timeout time.Duration) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	return &FindConfigLargeMemoryAllocationWarningThresholdParams{
		timeout: timeout,
	}
}

// NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithContext creates a new FindConfigLargeMemoryAllocationWarningThresholdParams object
// with the ability to set a context for a request.
func NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithContext(ctx context.Context) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	return &FindConfigLargeMemoryAllocationWarningThresholdParams{
		Context: ctx,
	}
}

// NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithHTTPClient creates a new FindConfigLargeMemoryAllocationWarningThresholdParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigLargeMemoryAllocationWarningThresholdParamsWithHTTPClient(client *http.Client) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	return &FindConfigLargeMemoryAllocationWarningThresholdParams{
		HTTPClient: client,
	}
}

/*
FindConfigLargeMemoryAllocationWarningThresholdParams contains all the parameters to send to the API endpoint

	for the find config large memory allocation warning threshold operation.

	Typically these are written to a http.Request.
*/
type FindConfigLargeMemoryAllocationWarningThresholdParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config large memory allocation warning threshold params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) WithDefaults() *FindConfigLargeMemoryAllocationWarningThresholdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config large memory allocation warning threshold params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) WithTimeout(timeout time.Duration) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) WithContext(ctx context.Context) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) WithHTTPClient(client *http.Client) *FindConfigLargeMemoryAllocationWarningThresholdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config large memory allocation warning threshold params
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigLargeMemoryAllocationWarningThresholdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}