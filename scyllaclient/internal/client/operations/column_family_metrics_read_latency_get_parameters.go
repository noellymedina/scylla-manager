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

// NewColumnFamilyMetricsReadLatencyGetParams creates a new ColumnFamilyMetricsReadLatencyGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyGetParams() *ColumnFamilyMetricsReadLatencyGetParams {

	return &ColumnFamilyMetricsReadLatencyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyGetParams {

	return &ColumnFamilyMetricsReadLatencyGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyGetParams {

	return &ColumnFamilyMetricsReadLatencyGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyGetParams {

	return &ColumnFamilyMetricsReadLatencyGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsReadLatencyGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency get params
func (o *ColumnFamilyMetricsReadLatencyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
