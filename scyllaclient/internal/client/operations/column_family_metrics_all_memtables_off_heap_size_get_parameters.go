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

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams() *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {

	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithTimeout creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {

	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithContext creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {

	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithHTTPClient creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {

	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams contains all the parameters to send to the API endpoint
for the column family metrics all memtables off heap size get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics all memtables off heap size get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
