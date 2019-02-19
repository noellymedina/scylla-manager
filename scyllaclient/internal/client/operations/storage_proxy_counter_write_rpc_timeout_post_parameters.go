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

// NewStorageProxyCounterWriteRPCTimeoutPostParams creates a new StorageProxyCounterWriteRPCTimeoutPostParams object
// with the default values initialized.
func NewStorageProxyCounterWriteRPCTimeoutPostParams() *StorageProxyCounterWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyCounterWriteRPCTimeoutPostParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyCounterWriteRPCTimeoutPostParamsWithTimeout creates a new StorageProxyCounterWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyCounterWriteRPCTimeoutPostParamsWithTimeout(timeout time.Duration) *StorageProxyCounterWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyCounterWriteRPCTimeoutPostParams{

		requestTimeout: timeout,
	}
}

// NewStorageProxyCounterWriteRPCTimeoutPostParamsWithContext creates a new StorageProxyCounterWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyCounterWriteRPCTimeoutPostParamsWithContext(ctx context.Context) *StorageProxyCounterWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyCounterWriteRPCTimeoutPostParams{

		Context: ctx,
	}
}

// NewStorageProxyCounterWriteRPCTimeoutPostParamsWithHTTPClient creates a new StorageProxyCounterWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyCounterWriteRPCTimeoutPostParamsWithHTTPClient(client *http.Client) *StorageProxyCounterWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyCounterWriteRPCTimeoutPostParams{
		HTTPClient: client,
	}
}

/*StorageProxyCounterWriteRPCTimeoutPostParams contains all the parameters to send to the API endpoint
for the storage proxy counter write Rpc timeout post operation typically these are written to a http.Request
*/
type StorageProxyCounterWriteRPCTimeoutPostParams struct {

	/*Timeout
	  timeout in seconds

	*/
	Timeout string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) WithRequestTimeout(timeout time.Duration) *StorageProxyCounterWriteRPCTimeoutPostParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) WithContext(ctx context.Context) *StorageProxyCounterWriteRPCTimeoutPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) WithHTTPClient(client *http.Client) *StorageProxyCounterWriteRPCTimeoutPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) WithTimeout(timeout string) *StorageProxyCounterWriteRPCTimeoutPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy counter write Rpc timeout post params
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) SetTimeout(timeout string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyCounterWriteRPCTimeoutPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := qrTimeout
	if qTimeout != "" {
		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
