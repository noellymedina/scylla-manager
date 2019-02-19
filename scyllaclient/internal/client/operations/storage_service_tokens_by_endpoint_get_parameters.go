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

// NewStorageServiceTokensByEndpointGetParams creates a new StorageServiceTokensByEndpointGetParams object
// with the default values initialized.
func NewStorageServiceTokensByEndpointGetParams() *StorageServiceTokensByEndpointGetParams {
	var ()
	return &StorageServiceTokensByEndpointGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceTokensByEndpointGetParamsWithTimeout creates a new StorageServiceTokensByEndpointGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceTokensByEndpointGetParamsWithTimeout(timeout time.Duration) *StorageServiceTokensByEndpointGetParams {
	var ()
	return &StorageServiceTokensByEndpointGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceTokensByEndpointGetParamsWithContext creates a new StorageServiceTokensByEndpointGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceTokensByEndpointGetParamsWithContext(ctx context.Context) *StorageServiceTokensByEndpointGetParams {
	var ()
	return &StorageServiceTokensByEndpointGetParams{

		Context: ctx,
	}
}

// NewStorageServiceTokensByEndpointGetParamsWithHTTPClient creates a new StorageServiceTokensByEndpointGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceTokensByEndpointGetParamsWithHTTPClient(client *http.Client) *StorageServiceTokensByEndpointGetParams {
	var ()
	return &StorageServiceTokensByEndpointGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceTokensByEndpointGetParams contains all the parameters to send to the API endpoint
for the storage service tokens by endpoint get operation typically these are written to a http.Request
*/
type StorageServiceTokensByEndpointGetParams struct {

	/*Endpoint
	  The endpoint

	*/
	Endpoint string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) WithTimeout(timeout time.Duration) *StorageServiceTokensByEndpointGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) WithContext(ctx context.Context) *StorageServiceTokensByEndpointGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) WithHTTPClient(client *http.Client) *StorageServiceTokensByEndpointGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpoint adds the endpoint to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) WithEndpoint(endpoint string) *StorageServiceTokensByEndpointGetParams {
	o.SetEndpoint(endpoint)
	return o
}

// SetEndpoint adds the endpoint to the storage service tokens by endpoint get params
func (o *StorageServiceTokensByEndpointGetParams) SetEndpoint(endpoint string) {
	o.Endpoint = endpoint
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceTokensByEndpointGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpoint
	if err := r.SetPathParam("endpoint", o.Endpoint); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
