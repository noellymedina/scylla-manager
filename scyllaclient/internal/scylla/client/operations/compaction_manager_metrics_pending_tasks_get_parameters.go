// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCompactionManagerMetricsPendingTasksGetParams creates a new CompactionManagerMetricsPendingTasksGetParams object
// with the default values initialized.
func NewCompactionManagerMetricsPendingTasksGetParams() *CompactionManagerMetricsPendingTasksGetParams {

	return &CompactionManagerMetricsPendingTasksGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompactionManagerMetricsPendingTasksGetParamsWithTimeout creates a new CompactionManagerMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompactionManagerMetricsPendingTasksGetParamsWithTimeout(timeout time.Duration) *CompactionManagerMetricsPendingTasksGetParams {

	return &CompactionManagerMetricsPendingTasksGetParams{

		timeout: timeout,
	}
}

// NewCompactionManagerMetricsPendingTasksGetParamsWithContext creates a new CompactionManagerMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompactionManagerMetricsPendingTasksGetParamsWithContext(ctx context.Context) *CompactionManagerMetricsPendingTasksGetParams {

	return &CompactionManagerMetricsPendingTasksGetParams{

		Context: ctx,
	}
}

// NewCompactionManagerMetricsPendingTasksGetParamsWithHTTPClient creates a new CompactionManagerMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompactionManagerMetricsPendingTasksGetParamsWithHTTPClient(client *http.Client) *CompactionManagerMetricsPendingTasksGetParams {

	return &CompactionManagerMetricsPendingTasksGetParams{
		HTTPClient: client,
	}
}

/*CompactionManagerMetricsPendingTasksGetParams contains all the parameters to send to the API endpoint
for the compaction manager metrics pending tasks get operation typically these are written to a http.Request
*/
type CompactionManagerMetricsPendingTasksGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) WithTimeout(timeout time.Duration) *CompactionManagerMetricsPendingTasksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) WithContext(ctx context.Context) *CompactionManagerMetricsPendingTasksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) WithHTTPClient(client *http.Client) *CompactionManagerMetricsPendingTasksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the compaction manager metrics pending tasks get params
func (o *CompactionManagerMetricsPendingTasksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CompactionManagerMetricsPendingTasksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
