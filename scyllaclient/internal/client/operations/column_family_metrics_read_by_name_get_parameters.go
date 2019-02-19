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

// NewColumnFamilyMetricsReadByNameGetParams creates a new ColumnFamilyMetricsReadByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadByNameGetParams() *ColumnFamilyMetricsReadByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsReadByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadByNameGetParamsWithContext creates a new ColumnFamilyMetricsReadByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsReadByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics read by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) WithName(name string) *ColumnFamilyMetricsReadByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics read by name get params
func (o *ColumnFamilyMetricsReadByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
