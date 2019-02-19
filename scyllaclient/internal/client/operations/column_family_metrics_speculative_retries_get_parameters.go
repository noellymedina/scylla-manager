// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsSpeculativeRetriesGetParams creates a new ColumnFamilyMetricsSpeculativeRetriesGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsSpeculativeRetriesGetParams() *ColumnFamilyMetricsSpeculativeRetriesGetParams {

	return &ColumnFamilyMetricsSpeculativeRetriesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithTimeout creates a new ColumnFamilyMetricsSpeculativeRetriesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsSpeculativeRetriesGetParams {

	return &ColumnFamilyMetricsSpeculativeRetriesGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithContext creates a new ColumnFamilyMetricsSpeculativeRetriesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsSpeculativeRetriesGetParams {

	return &ColumnFamilyMetricsSpeculativeRetriesGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithHTTPClient creates a new ColumnFamilyMetricsSpeculativeRetriesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsSpeculativeRetriesGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsSpeculativeRetriesGetParams {

	return &ColumnFamilyMetricsSpeculativeRetriesGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsSpeculativeRetriesGetParams contains all the parameters to send to the API endpoint
for the column family metrics speculative retries get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsSpeculativeRetriesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsSpeculativeRetriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsSpeculativeRetriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsSpeculativeRetriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics speculative retries get params
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsSpeculativeRetriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
