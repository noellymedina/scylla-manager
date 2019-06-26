// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageProxyMetricsCasReadUnfinishedCommitGetParams creates a new StorageProxyMetricsCasReadUnfinishedCommitGetParams object
// with the default values initialized.
func NewStorageProxyMetricsCasReadUnfinishedCommitGetParams() *StorageProxyMetricsCasReadUnfinishedCommitGetParams {

	return &StorageProxyMetricsCasReadUnfinishedCommitGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithTimeout creates a new StorageProxyMetricsCasReadUnfinishedCommitGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {

	return &StorageProxyMetricsCasReadUnfinishedCommitGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithContext creates a new StorageProxyMetricsCasReadUnfinishedCommitGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithContext(ctx context.Context) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {

	return &StorageProxyMetricsCasReadUnfinishedCommitGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithHTTPClient creates a new StorageProxyMetricsCasReadUnfinishedCommitGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsCasReadUnfinishedCommitGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {

	return &StorageProxyMetricsCasReadUnfinishedCommitGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsCasReadUnfinishedCommitGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics cas read unfinished commit get operation typically these are written to a http.Request
*/
type StorageProxyMetricsCasReadUnfinishedCommitGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) WithContext(ctx context.Context) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadUnfinishedCommitGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics cas read unfinished commit get params
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsCasReadUnfinishedCommitGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
