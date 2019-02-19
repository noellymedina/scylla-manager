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

// NewStorageProxyReadRepairRepairedBlockingGetParams creates a new StorageProxyReadRepairRepairedBlockingGetParams object
// with the default values initialized.
func NewStorageProxyReadRepairRepairedBlockingGetParams() *StorageProxyReadRepairRepairedBlockingGetParams {

	return &StorageProxyReadRepairRepairedBlockingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyReadRepairRepairedBlockingGetParamsWithTimeout creates a new StorageProxyReadRepairRepairedBlockingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyReadRepairRepairedBlockingGetParamsWithTimeout(timeout time.Duration) *StorageProxyReadRepairRepairedBlockingGetParams {

	return &StorageProxyReadRepairRepairedBlockingGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyReadRepairRepairedBlockingGetParamsWithContext creates a new StorageProxyReadRepairRepairedBlockingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyReadRepairRepairedBlockingGetParamsWithContext(ctx context.Context) *StorageProxyReadRepairRepairedBlockingGetParams {

	return &StorageProxyReadRepairRepairedBlockingGetParams{

		Context: ctx,
	}
}

// NewStorageProxyReadRepairRepairedBlockingGetParamsWithHTTPClient creates a new StorageProxyReadRepairRepairedBlockingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyReadRepairRepairedBlockingGetParamsWithHTTPClient(client *http.Client) *StorageProxyReadRepairRepairedBlockingGetParams {

	return &StorageProxyReadRepairRepairedBlockingGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyReadRepairRepairedBlockingGetParams contains all the parameters to send to the API endpoint
for the storage proxy read repair repaired blocking get operation typically these are written to a http.Request
*/
type StorageProxyReadRepairRepairedBlockingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) WithTimeout(timeout time.Duration) *StorageProxyReadRepairRepairedBlockingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) WithContext(ctx context.Context) *StorageProxyReadRepairRepairedBlockingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) WithHTTPClient(client *http.Client) *StorageProxyReadRepairRepairedBlockingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy read repair repaired blocking get params
func (o *StorageProxyReadRepairRepairedBlockingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyReadRepairRepairedBlockingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
