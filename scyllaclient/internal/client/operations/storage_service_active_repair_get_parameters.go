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

// NewStorageServiceActiveRepairGetParams creates a new StorageServiceActiveRepairGetParams object
// with the default values initialized.
func NewStorageServiceActiveRepairGetParams() *StorageServiceActiveRepairGetParams {

	return &StorageServiceActiveRepairGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceActiveRepairGetParamsWithTimeout creates a new StorageServiceActiveRepairGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceActiveRepairGetParamsWithTimeout(timeout time.Duration) *StorageServiceActiveRepairGetParams {

	return &StorageServiceActiveRepairGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceActiveRepairGetParamsWithContext creates a new StorageServiceActiveRepairGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceActiveRepairGetParamsWithContext(ctx context.Context) *StorageServiceActiveRepairGetParams {

	return &StorageServiceActiveRepairGetParams{

		Context: ctx,
	}
}

// NewStorageServiceActiveRepairGetParamsWithHTTPClient creates a new StorageServiceActiveRepairGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceActiveRepairGetParamsWithHTTPClient(client *http.Client) *StorageServiceActiveRepairGetParams {

	return &StorageServiceActiveRepairGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceActiveRepairGetParams contains all the parameters to send to the API endpoint
for the storage service active repair get operation typically these are written to a http.Request
*/
type StorageServiceActiveRepairGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) WithTimeout(timeout time.Duration) *StorageServiceActiveRepairGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) WithContext(ctx context.Context) *StorageServiceActiveRepairGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) WithHTTPClient(client *http.Client) *StorageServiceActiveRepairGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service active repair get params
func (o *StorageServiceActiveRepairGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceActiveRepairGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
