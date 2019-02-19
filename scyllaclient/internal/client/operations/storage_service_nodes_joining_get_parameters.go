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

// NewStorageServiceNodesJoiningGetParams creates a new StorageServiceNodesJoiningGetParams object
// with the default values initialized.
func NewStorageServiceNodesJoiningGetParams() *StorageServiceNodesJoiningGetParams {

	return &StorageServiceNodesJoiningGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceNodesJoiningGetParamsWithTimeout creates a new StorageServiceNodesJoiningGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceNodesJoiningGetParamsWithTimeout(timeout time.Duration) *StorageServiceNodesJoiningGetParams {

	return &StorageServiceNodesJoiningGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceNodesJoiningGetParamsWithContext creates a new StorageServiceNodesJoiningGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceNodesJoiningGetParamsWithContext(ctx context.Context) *StorageServiceNodesJoiningGetParams {

	return &StorageServiceNodesJoiningGetParams{

		Context: ctx,
	}
}

// NewStorageServiceNodesJoiningGetParamsWithHTTPClient creates a new StorageServiceNodesJoiningGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceNodesJoiningGetParamsWithHTTPClient(client *http.Client) *StorageServiceNodesJoiningGetParams {

	return &StorageServiceNodesJoiningGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceNodesJoiningGetParams contains all the parameters to send to the API endpoint
for the storage service nodes joining get operation typically these are written to a http.Request
*/
type StorageServiceNodesJoiningGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) WithTimeout(timeout time.Duration) *StorageServiceNodesJoiningGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) WithContext(ctx context.Context) *StorageServiceNodesJoiningGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) WithHTTPClient(client *http.Client) *StorageServiceNodesJoiningGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service nodes joining get params
func (o *StorageServiceNodesJoiningGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceNodesJoiningGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
