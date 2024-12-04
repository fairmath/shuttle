// Code generated by go-swagger; DO NOT EDIT.

package query

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
	"github.com/go-openapi/swag"
)

// NewQuerySendEnabledParams creates a new QuerySendEnabledParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuerySendEnabledParams() *QuerySendEnabledParams {
	return &QuerySendEnabledParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuerySendEnabledParamsWithTimeout creates a new QuerySendEnabledParams object
// with the ability to set a timeout on a request.
func NewQuerySendEnabledParamsWithTimeout(timeout time.Duration) *QuerySendEnabledParams {
	return &QuerySendEnabledParams{
		timeout: timeout,
	}
}

// NewQuerySendEnabledParamsWithContext creates a new QuerySendEnabledParams object
// with the ability to set a context for a request.
func NewQuerySendEnabledParamsWithContext(ctx context.Context) *QuerySendEnabledParams {
	return &QuerySendEnabledParams{
		Context: ctx,
	}
}

// NewQuerySendEnabledParamsWithHTTPClient creates a new QuerySendEnabledParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuerySendEnabledParamsWithHTTPClient(client *http.Client) *QuerySendEnabledParams {
	return &QuerySendEnabledParams{
		HTTPClient: client,
	}
}

/*
QuerySendEnabledParams contains all the parameters to send to the API endpoint

	for the query send enabled operation.

	Typically these are written to a http.Request.
*/
type QuerySendEnabledParams struct {

	// Denoms.
	Denoms []string

	// PaginationCountTotal.
	PaginationCountTotal *bool

	// PaginationKey.
	//
	// Format: byte
	PaginationKey *strfmt.Base64

	// PaginationLimit.
	//
	// Format: uint64
	PaginationLimit *string

	// PaginationOffset.
	//
	// Format: uint64
	PaginationOffset *string

	// PaginationReverse.
	PaginationReverse *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query send enabled params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuerySendEnabledParams) WithDefaults() *QuerySendEnabledParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query send enabled params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuerySendEnabledParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query send enabled params
func (o *QuerySendEnabledParams) WithTimeout(timeout time.Duration) *QuerySendEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query send enabled params
func (o *QuerySendEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query send enabled params
func (o *QuerySendEnabledParams) WithContext(ctx context.Context) *QuerySendEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query send enabled params
func (o *QuerySendEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query send enabled params
func (o *QuerySendEnabledParams) WithHTTPClient(client *http.Client) *QuerySendEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query send enabled params
func (o *QuerySendEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenoms adds the denoms to the query send enabled params
func (o *QuerySendEnabledParams) WithDenoms(denoms []string) *QuerySendEnabledParams {
	o.SetDenoms(denoms)
	return o
}

// SetDenoms adds the denoms to the query send enabled params
func (o *QuerySendEnabledParams) SetDenoms(denoms []string) {
	o.Denoms = denoms
}

// WithPaginationCountTotal adds the paginationCountTotal to the query send enabled params
func (o *QuerySendEnabledParams) WithPaginationCountTotal(paginationCountTotal *bool) *QuerySendEnabledParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the query send enabled params
func (o *QuerySendEnabledParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the query send enabled params
func (o *QuerySendEnabledParams) WithPaginationKey(paginationKey *strfmt.Base64) *QuerySendEnabledParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the query send enabled params
func (o *QuerySendEnabledParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the query send enabled params
func (o *QuerySendEnabledParams) WithPaginationLimit(paginationLimit *string) *QuerySendEnabledParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the query send enabled params
func (o *QuerySendEnabledParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the query send enabled params
func (o *QuerySendEnabledParams) WithPaginationOffset(paginationOffset *string) *QuerySendEnabledParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the query send enabled params
func (o *QuerySendEnabledParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the query send enabled params
func (o *QuerySendEnabledParams) WithPaginationReverse(paginationReverse *bool) *QuerySendEnabledParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the query send enabled params
func (o *QuerySendEnabledParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WriteToRequest writes these params to a swagger request
func (o *QuerySendEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Denoms != nil {

		// binding items for denoms
		joinedDenoms := o.bindParamDenoms(reg)

		// query array param denoms
		if err := r.SetQueryParam("denoms", joinedDenoms...); err != nil {
			return err
		}
	}

	if o.PaginationCountTotal != nil {

		// query param pagination.count_total
		var qrPaginationCountTotal bool

		if o.PaginationCountTotal != nil {
			qrPaginationCountTotal = *o.PaginationCountTotal
		}
		qPaginationCountTotal := swag.FormatBool(qrPaginationCountTotal)
		if qPaginationCountTotal != "" {

			if err := r.SetQueryParam("pagination.count_total", qPaginationCountTotal); err != nil {
				return err
			}
		}
	}

	if o.PaginationKey != nil {

		// query param pagination.key
		var qrPaginationKey strfmt.Base64

		if o.PaginationKey != nil {
			qrPaginationKey = *o.PaginationKey
		}
		qPaginationKey := qrPaginationKey.String()
		if qPaginationKey != "" {

			if err := r.SetQueryParam("pagination.key", qPaginationKey); err != nil {
				return err
			}
		}
	}

	if o.PaginationLimit != nil {

		// query param pagination.limit
		var qrPaginationLimit string

		if o.PaginationLimit != nil {
			qrPaginationLimit = *o.PaginationLimit
		}
		qPaginationLimit := qrPaginationLimit
		if qPaginationLimit != "" {

			if err := r.SetQueryParam("pagination.limit", qPaginationLimit); err != nil {
				return err
			}
		}
	}

	if o.PaginationOffset != nil {

		// query param pagination.offset
		var qrPaginationOffset string

		if o.PaginationOffset != nil {
			qrPaginationOffset = *o.PaginationOffset
		}
		qPaginationOffset := qrPaginationOffset
		if qPaginationOffset != "" {

			if err := r.SetQueryParam("pagination.offset", qPaginationOffset); err != nil {
				return err
			}
		}
	}

	if o.PaginationReverse != nil {

		// query param pagination.reverse
		var qrPaginationReverse bool

		if o.PaginationReverse != nil {
			qrPaginationReverse = *o.PaginationReverse
		}
		qPaginationReverse := swag.FormatBool(qrPaginationReverse)
		if qPaginationReverse != "" {

			if err := r.SetQueryParam("pagination.reverse", qPaginationReverse); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQuerySendEnabled binds the parameter denoms
func (o *QuerySendEnabledParams) bindParamDenoms(formats strfmt.Registry) []string {
	denomsIR := o.Denoms

	var denomsIC []string
	for _, denomsIIR := range denomsIR { // explode []string

		denomsIIV := denomsIIR // string as string
		denomsIC = append(denomsIC, denomsIIV)
	}

	// items.CollectionFormat: "multi"
	denomsIS := swag.JoinByFormat(denomsIC, "multi")

	return denomsIS
}
