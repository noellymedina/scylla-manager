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

// NewSystemLoggerByNamePostParams creates a new SystemLoggerByNamePostParams object
// with the default values initialized.
func NewSystemLoggerByNamePostParams() *SystemLoggerByNamePostParams {
	var ()
	return &SystemLoggerByNamePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemLoggerByNamePostParamsWithTimeout creates a new SystemLoggerByNamePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemLoggerByNamePostParamsWithTimeout(timeout time.Duration) *SystemLoggerByNamePostParams {
	var ()
	return &SystemLoggerByNamePostParams{

		timeout: timeout,
	}
}

// NewSystemLoggerByNamePostParamsWithContext creates a new SystemLoggerByNamePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemLoggerByNamePostParamsWithContext(ctx context.Context) *SystemLoggerByNamePostParams {
	var ()
	return &SystemLoggerByNamePostParams{

		Context: ctx,
	}
}

// NewSystemLoggerByNamePostParamsWithHTTPClient creates a new SystemLoggerByNamePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemLoggerByNamePostParamsWithHTTPClient(client *http.Client) *SystemLoggerByNamePostParams {
	var ()
	return &SystemLoggerByNamePostParams{
		HTTPClient: client,
	}
}

/*SystemLoggerByNamePostParams contains all the parameters to send to the API endpoint
for the system logger by name post operation typically these are written to a http.Request
*/
type SystemLoggerByNamePostParams struct {

	/*Level
	  The new log level

	*/
	Level string
	/*Name
	  The logger to query about

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system logger by name post params
func (o *SystemLoggerByNamePostParams) WithTimeout(timeout time.Duration) *SystemLoggerByNamePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system logger by name post params
func (o *SystemLoggerByNamePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system logger by name post params
func (o *SystemLoggerByNamePostParams) WithContext(ctx context.Context) *SystemLoggerByNamePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system logger by name post params
func (o *SystemLoggerByNamePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system logger by name post params
func (o *SystemLoggerByNamePostParams) WithHTTPClient(client *http.Client) *SystemLoggerByNamePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system logger by name post params
func (o *SystemLoggerByNamePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevel adds the level to the system logger by name post params
func (o *SystemLoggerByNamePostParams) WithLevel(level string) *SystemLoggerByNamePostParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the system logger by name post params
func (o *SystemLoggerByNamePostParams) SetLevel(level string) {
	o.Level = level
}

// WithName adds the name to the system logger by name post params
func (o *SystemLoggerByNamePostParams) WithName(name string) *SystemLoggerByNamePostParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the system logger by name post params
func (o *SystemLoggerByNamePostParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SystemLoggerByNamePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param level
	qrLevel := o.Level
	qLevel := qrLevel
	if qLevel != "" {
		if err := r.SetQueryParam("level", qLevel); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
