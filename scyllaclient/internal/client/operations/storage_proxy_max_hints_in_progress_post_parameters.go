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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageProxyMaxHintsInProgressPostParams creates a new StorageProxyMaxHintsInProgressPostParams object
// with the default values initialized.
func NewStorageProxyMaxHintsInProgressPostParams() *StorageProxyMaxHintsInProgressPostParams {
	var ()
	return &StorageProxyMaxHintsInProgressPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMaxHintsInProgressPostParamsWithTimeout creates a new StorageProxyMaxHintsInProgressPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMaxHintsInProgressPostParamsWithTimeout(timeout time.Duration) *StorageProxyMaxHintsInProgressPostParams {
	var ()
	return &StorageProxyMaxHintsInProgressPostParams{

		timeout: timeout,
	}
}

// NewStorageProxyMaxHintsInProgressPostParamsWithContext creates a new StorageProxyMaxHintsInProgressPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMaxHintsInProgressPostParamsWithContext(ctx context.Context) *StorageProxyMaxHintsInProgressPostParams {
	var ()
	return &StorageProxyMaxHintsInProgressPostParams{

		Context: ctx,
	}
}

// NewStorageProxyMaxHintsInProgressPostParamsWithHTTPClient creates a new StorageProxyMaxHintsInProgressPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMaxHintsInProgressPostParamsWithHTTPClient(client *http.Client) *StorageProxyMaxHintsInProgressPostParams {
	var ()
	return &StorageProxyMaxHintsInProgressPostParams{
		HTTPClient: client,
	}
}

/*StorageProxyMaxHintsInProgressPostParams contains all the parameters to send to the API endpoint
for the storage proxy max hints in progress post operation typically these are written to a http.Request
*/
type StorageProxyMaxHintsInProgressPostParams struct {

	/*Qs
	  max hints in progress

	*/
	Qs int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) WithTimeout(timeout time.Duration) *StorageProxyMaxHintsInProgressPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) WithContext(ctx context.Context) *StorageProxyMaxHintsInProgressPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) WithHTTPClient(client *http.Client) *StorageProxyMaxHintsInProgressPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQs adds the qs to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) WithQs(qs int32) *StorageProxyMaxHintsInProgressPostParams {
	o.SetQs(qs)
	return o
}

// SetQs adds the qs to the storage proxy max hints in progress post params
func (o *StorageProxyMaxHintsInProgressPostParams) SetQs(qs int32) {
	o.Qs = qs
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMaxHintsInProgressPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param qs
	qrQs := o.Qs
	qQs := swag.FormatInt32(qrQs)
	if qQs != "" {
		if err := r.SetQueryParam("qs", qQs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
