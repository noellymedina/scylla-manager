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

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams creates a new ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams() *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics cas propose estimated recent histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics cas propose estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
