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

// NewStorageServiceDrainPostParams creates a new StorageServiceDrainPostParams object
// with the default values initialized.
func NewStorageServiceDrainPostParams() *StorageServiceDrainPostParams {

	return &StorageServiceDrainPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceDrainPostParamsWithTimeout creates a new StorageServiceDrainPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceDrainPostParamsWithTimeout(timeout time.Duration) *StorageServiceDrainPostParams {

	return &StorageServiceDrainPostParams{

		timeout: timeout,
	}
}

// NewStorageServiceDrainPostParamsWithContext creates a new StorageServiceDrainPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceDrainPostParamsWithContext(ctx context.Context) *StorageServiceDrainPostParams {

	return &StorageServiceDrainPostParams{

		Context: ctx,
	}
}

// NewStorageServiceDrainPostParamsWithHTTPClient creates a new StorageServiceDrainPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceDrainPostParamsWithHTTPClient(client *http.Client) *StorageServiceDrainPostParams {

	return &StorageServiceDrainPostParams{
		HTTPClient: client,
	}
}

/*StorageServiceDrainPostParams contains all the parameters to send to the API endpoint
for the storage service drain post operation typically these are written to a http.Request
*/
type StorageServiceDrainPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service drain post params
func (o *StorageServiceDrainPostParams) WithTimeout(timeout time.Duration) *StorageServiceDrainPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service drain post params
func (o *StorageServiceDrainPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service drain post params
func (o *StorageServiceDrainPostParams) WithContext(ctx context.Context) *StorageServiceDrainPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service drain post params
func (o *StorageServiceDrainPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service drain post params
func (o *StorageServiceDrainPostParams) WithHTTPClient(client *http.Client) *StorageServiceDrainPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service drain post params
func (o *StorageServiceDrainPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceDrainPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
