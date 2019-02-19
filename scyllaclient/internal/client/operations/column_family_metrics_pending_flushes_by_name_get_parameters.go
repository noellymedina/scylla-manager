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

// NewColumnFamilyMetricsPendingFlushesByNameGetParams creates a new ColumnFamilyMetricsPendingFlushesByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsPendingFlushesByNameGetParams() *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingFlushesByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsPendingFlushesByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingFlushesByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithContext creates a new ColumnFamilyMetricsPendingFlushesByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingFlushesByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsPendingFlushesByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsPendingFlushesByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingFlushesByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsPendingFlushesByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics pending flushes by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsPendingFlushesByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) WithName(name string) *ColumnFamilyMetricsPendingFlushesByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics pending flushes by name get params
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsPendingFlushesByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
