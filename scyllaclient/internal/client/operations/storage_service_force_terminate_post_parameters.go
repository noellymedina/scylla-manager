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

// NewStorageServiceForceTerminatePostParams creates a new StorageServiceForceTerminatePostParams object
// with the default values initialized.
func NewStorageServiceForceTerminatePostParams() *StorageServiceForceTerminatePostParams {

	return &StorageServiceForceTerminatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceForceTerminatePostParamsWithTimeout creates a new StorageServiceForceTerminatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceForceTerminatePostParamsWithTimeout(timeout time.Duration) *StorageServiceForceTerminatePostParams {

	return &StorageServiceForceTerminatePostParams{

		timeout: timeout,
	}
}

// NewStorageServiceForceTerminatePostParamsWithContext creates a new StorageServiceForceTerminatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceForceTerminatePostParamsWithContext(ctx context.Context) *StorageServiceForceTerminatePostParams {

	return &StorageServiceForceTerminatePostParams{

		Context: ctx,
	}
}

// NewStorageServiceForceTerminatePostParamsWithHTTPClient creates a new StorageServiceForceTerminatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceForceTerminatePostParamsWithHTTPClient(client *http.Client) *StorageServiceForceTerminatePostParams {

	return &StorageServiceForceTerminatePostParams{
		HTTPClient: client,
	}
}

/*StorageServiceForceTerminatePostParams contains all the parameters to send to the API endpoint
for the storage service force terminate post operation typically these are written to a http.Request
*/
type StorageServiceForceTerminatePostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) WithTimeout(timeout time.Duration) *StorageServiceForceTerminatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) WithContext(ctx context.Context) *StorageServiceForceTerminatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) WithHTTPClient(client *http.Client) *StorageServiceForceTerminatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service force terminate post params
func (o *StorageServiceForceTerminatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceForceTerminatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
