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

// NewStorageProxyMetricsWriteHistogramGetParams creates a new StorageProxyMetricsWriteHistogramGetParams object
// with the default values initialized.
func NewStorageProxyMetricsWriteHistogramGetParams() *StorageProxyMetricsWriteHistogramGetParams {

	return &StorageProxyMetricsWriteHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsWriteHistogramGetParamsWithTimeout creates a new StorageProxyMetricsWriteHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsWriteHistogramGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsWriteHistogramGetParams {

	return &StorageProxyMetricsWriteHistogramGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsWriteHistogramGetParamsWithContext creates a new StorageProxyMetricsWriteHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsWriteHistogramGetParamsWithContext(ctx context.Context) *StorageProxyMetricsWriteHistogramGetParams {

	return &StorageProxyMetricsWriteHistogramGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsWriteHistogramGetParamsWithHTTPClient creates a new StorageProxyMetricsWriteHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsWriteHistogramGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsWriteHistogramGetParams {

	return &StorageProxyMetricsWriteHistogramGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsWriteHistogramGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics write histogram get operation typically these are written to a http.Request
*/
type StorageProxyMetricsWriteHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsWriteHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) WithContext(ctx context.Context) *StorageProxyMetricsWriteHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsWriteHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics write histogram get params
func (o *StorageProxyMetricsWriteHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsWriteHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
