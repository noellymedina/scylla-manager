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

// NewColumnFamilyMetricsWriteLatencyHistogramGetParams creates a new ColumnFamilyMetricsWriteLatencyHistogramGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsWriteLatencyHistogramGetParams() *ColumnFamilyMetricsWriteLatencyHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithTimeout creates a new ColumnFamilyMetricsWriteLatencyHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyHistogramGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithContext creates a new ColumnFamilyMetricsWriteLatencyHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyHistogramGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithHTTPClient creates a new ColumnFamilyMetricsWriteLatencyHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsWriteLatencyHistogramGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyHistogramGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsWriteLatencyHistogramGetParams contains all the parameters to send to the API endpoint
for the column family metrics write latency histogram get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsWriteLatencyHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics write latency histogram get params
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
