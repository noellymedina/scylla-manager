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

// NewColumnFamilyMetricsWaitingOnFreeMemtableGetParams creates a new ColumnFamilyMetricsWaitingOnFreeMemtableGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsWaitingOnFreeMemtableGetParams() *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {

	return &ColumnFamilyMetricsWaitingOnFreeMemtableGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithTimeout creates a new ColumnFamilyMetricsWaitingOnFreeMemtableGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {

	return &ColumnFamilyMetricsWaitingOnFreeMemtableGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithContext creates a new ColumnFamilyMetricsWaitingOnFreeMemtableGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {

	return &ColumnFamilyMetricsWaitingOnFreeMemtableGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithHTTPClient creates a new ColumnFamilyMetricsWaitingOnFreeMemtableGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsWaitingOnFreeMemtableGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {

	return &ColumnFamilyMetricsWaitingOnFreeMemtableGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsWaitingOnFreeMemtableGetParams contains all the parameters to send to the API endpoint
for the column family metrics waiting on free memtable get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsWaitingOnFreeMemtableGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics waiting on free memtable get params
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsWaitingOnFreeMemtableGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}