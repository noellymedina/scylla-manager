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

// NewStorageProxyMetricsRangeUnavailablesRatesGetParams creates a new StorageProxyMetricsRangeUnavailablesRatesGetParams object
// with the default values initialized.
func NewStorageProxyMetricsRangeUnavailablesRatesGetParams() *StorageProxyMetricsRangeUnavailablesRatesGetParams {

	return &StorageProxyMetricsRangeUnavailablesRatesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithTimeout creates a new StorageProxyMetricsRangeUnavailablesRatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsRangeUnavailablesRatesGetParams {

	return &StorageProxyMetricsRangeUnavailablesRatesGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithContext creates a new StorageProxyMetricsRangeUnavailablesRatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithContext(ctx context.Context) *StorageProxyMetricsRangeUnavailablesRatesGetParams {

	return &StorageProxyMetricsRangeUnavailablesRatesGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithHTTPClient creates a new StorageProxyMetricsRangeUnavailablesRatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsRangeUnavailablesRatesGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsRangeUnavailablesRatesGetParams {

	return &StorageProxyMetricsRangeUnavailablesRatesGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsRangeUnavailablesRatesGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics range unavailables rates get operation typically these are written to a http.Request
*/
type StorageProxyMetricsRangeUnavailablesRatesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsRangeUnavailablesRatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) WithContext(ctx context.Context) *StorageProxyMetricsRangeUnavailablesRatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsRangeUnavailablesRatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics range unavailables rates get params
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsRangeUnavailablesRatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
