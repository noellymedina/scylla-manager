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

// NewStorageProxyMaxHintsInProgressGetParams creates a new StorageProxyMaxHintsInProgressGetParams object
// with the default values initialized.
func NewStorageProxyMaxHintsInProgressGetParams() *StorageProxyMaxHintsInProgressGetParams {

	return &StorageProxyMaxHintsInProgressGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMaxHintsInProgressGetParamsWithTimeout creates a new StorageProxyMaxHintsInProgressGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMaxHintsInProgressGetParamsWithTimeout(timeout time.Duration) *StorageProxyMaxHintsInProgressGetParams {

	return &StorageProxyMaxHintsInProgressGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMaxHintsInProgressGetParamsWithContext creates a new StorageProxyMaxHintsInProgressGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMaxHintsInProgressGetParamsWithContext(ctx context.Context) *StorageProxyMaxHintsInProgressGetParams {

	return &StorageProxyMaxHintsInProgressGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMaxHintsInProgressGetParamsWithHTTPClient creates a new StorageProxyMaxHintsInProgressGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMaxHintsInProgressGetParamsWithHTTPClient(client *http.Client) *StorageProxyMaxHintsInProgressGetParams {

	return &StorageProxyMaxHintsInProgressGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMaxHintsInProgressGetParams contains all the parameters to send to the API endpoint
for the storage proxy max hints in progress get operation typically these are written to a http.Request
*/
type StorageProxyMaxHintsInProgressGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) WithTimeout(timeout time.Duration) *StorageProxyMaxHintsInProgressGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) WithContext(ctx context.Context) *StorageProxyMaxHintsInProgressGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) WithHTTPClient(client *http.Client) *StorageProxyMaxHintsInProgressGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy max hints in progress get params
func (o *StorageProxyMaxHintsInProgressGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMaxHintsInProgressGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
