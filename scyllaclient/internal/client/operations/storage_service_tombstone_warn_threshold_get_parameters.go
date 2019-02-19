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

// NewStorageServiceTombstoneWarnThresholdGetParams creates a new StorageServiceTombstoneWarnThresholdGetParams object
// with the default values initialized.
func NewStorageServiceTombstoneWarnThresholdGetParams() *StorageServiceTombstoneWarnThresholdGetParams {

	return &StorageServiceTombstoneWarnThresholdGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceTombstoneWarnThresholdGetParamsWithTimeout creates a new StorageServiceTombstoneWarnThresholdGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceTombstoneWarnThresholdGetParamsWithTimeout(timeout time.Duration) *StorageServiceTombstoneWarnThresholdGetParams {

	return &StorageServiceTombstoneWarnThresholdGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceTombstoneWarnThresholdGetParamsWithContext creates a new StorageServiceTombstoneWarnThresholdGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceTombstoneWarnThresholdGetParamsWithContext(ctx context.Context) *StorageServiceTombstoneWarnThresholdGetParams {

	return &StorageServiceTombstoneWarnThresholdGetParams{

		Context: ctx,
	}
}

// NewStorageServiceTombstoneWarnThresholdGetParamsWithHTTPClient creates a new StorageServiceTombstoneWarnThresholdGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceTombstoneWarnThresholdGetParamsWithHTTPClient(client *http.Client) *StorageServiceTombstoneWarnThresholdGetParams {

	return &StorageServiceTombstoneWarnThresholdGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceTombstoneWarnThresholdGetParams contains all the parameters to send to the API endpoint
for the storage service tombstone warn threshold get operation typically these are written to a http.Request
*/
type StorageServiceTombstoneWarnThresholdGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) WithTimeout(timeout time.Duration) *StorageServiceTombstoneWarnThresholdGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) WithContext(ctx context.Context) *StorageServiceTombstoneWarnThresholdGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) WithHTTPClient(client *http.Client) *StorageServiceTombstoneWarnThresholdGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service tombstone warn threshold get params
func (o *StorageServiceTombstoneWarnThresholdGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceTombstoneWarnThresholdGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
