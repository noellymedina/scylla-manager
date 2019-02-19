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

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetParams creates a new ColumnFamilyMetricsMemtableColumnsCountByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetParams() *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsMemtableColumnsCountByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithContext creates a new ColumnFamilyMetricsMemtableColumnsCountByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsMemtableColumnsCountByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsMemtableColumnsCountByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics memtable columns count by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsMemtableColumnsCountByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) WithName(name string) *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics memtable columns count by name get params
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
