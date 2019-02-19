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

// NewColumnFamilyMetricsPendingCompactionsByNameGetParams creates a new ColumnFamilyMetricsPendingCompactionsByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsPendingCompactionsByNameGetParams() *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingCompactionsByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsPendingCompactionsByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingCompactionsByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithContext creates a new ColumnFamilyMetricsPendingCompactionsByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingCompactionsByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsPendingCompactionsByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsPendingCompactionsByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	var ()
	return &ColumnFamilyMetricsPendingCompactionsByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsPendingCompactionsByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics pending compactions by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsPendingCompactionsByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) WithName(name string) *ColumnFamilyMetricsPendingCompactionsByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics pending compactions by name get params
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
