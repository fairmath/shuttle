// Code generated by go-swagger; DO NOT EDIT.

package service

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

// NewServiceBroadcastTxParams creates a new ServiceBroadcastTxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceBroadcastTxParams() *ServiceBroadcastTxParams {
	return &ServiceBroadcastTxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBroadcastTxParamsWithTimeout creates a new ServiceBroadcastTxParams object
// with the ability to set a timeout on a request.
func NewServiceBroadcastTxParamsWithTimeout(timeout time.Duration) *ServiceBroadcastTxParams {
	return &ServiceBroadcastTxParams{
		timeout: timeout,
	}
}

// NewServiceBroadcastTxParamsWithContext creates a new ServiceBroadcastTxParams object
// with the ability to set a context for a request.
func NewServiceBroadcastTxParamsWithContext(ctx context.Context) *ServiceBroadcastTxParams {
	return &ServiceBroadcastTxParams{
		Context: ctx,
	}
}

// NewServiceBroadcastTxParamsWithHTTPClient creates a new ServiceBroadcastTxParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceBroadcastTxParamsWithHTTPClient(client *http.Client) *ServiceBroadcastTxParams {
	return &ServiceBroadcastTxParams{
		HTTPClient: client,
	}
}

/*
ServiceBroadcastTxParams contains all the parameters to send to the API endpoint

	for the service broadcast tx operation.

	Typically these are written to a http.Request.
*/
type ServiceBroadcastTxParams struct {

	// Body.
	Body ServiceBroadcastTxBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service broadcast tx params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBroadcastTxParams) WithDefaults() *ServiceBroadcastTxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service broadcast tx params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBroadcastTxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service broadcast tx params
func (o *ServiceBroadcastTxParams) WithTimeout(timeout time.Duration) *ServiceBroadcastTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broadcast tx params
func (o *ServiceBroadcastTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broadcast tx params
func (o *ServiceBroadcastTxParams) WithContext(ctx context.Context) *ServiceBroadcastTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broadcast tx params
func (o *ServiceBroadcastTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broadcast tx params
func (o *ServiceBroadcastTxParams) WithHTTPClient(client *http.Client) *ServiceBroadcastTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broadcast tx params
func (o *ServiceBroadcastTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service broadcast tx params
func (o *ServiceBroadcastTxParams) WithBody(body ServiceBroadcastTxBody) *ServiceBroadcastTxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service broadcast tx params
func (o *ServiceBroadcastTxParams) SetBody(body ServiceBroadcastTxBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBroadcastTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
