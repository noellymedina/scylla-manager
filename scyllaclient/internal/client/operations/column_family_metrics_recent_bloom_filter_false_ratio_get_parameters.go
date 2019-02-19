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

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams creates a new ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams() *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithTimeout creates a new ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithContext creates a new ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithHTTPClient creates a new ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams contains all the parameters to send to the API endpoint
for the column family metrics recent bloom filter false ratio get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics recent bloom filter false ratio get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
