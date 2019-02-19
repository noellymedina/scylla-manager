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

// NewHintedHandoffHintsGetParams creates a new HintedHandoffHintsGetParams object
// with the default values initialized.
func NewHintedHandoffHintsGetParams() *HintedHandoffHintsGetParams {

	return &HintedHandoffHintsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHintedHandoffHintsGetParamsWithTimeout creates a new HintedHandoffHintsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHintedHandoffHintsGetParamsWithTimeout(timeout time.Duration) *HintedHandoffHintsGetParams {

	return &HintedHandoffHintsGetParams{

		timeout: timeout,
	}
}

// NewHintedHandoffHintsGetParamsWithContext creates a new HintedHandoffHintsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHintedHandoffHintsGetParamsWithContext(ctx context.Context) *HintedHandoffHintsGetParams {

	return &HintedHandoffHintsGetParams{

		Context: ctx,
	}
}

// NewHintedHandoffHintsGetParamsWithHTTPClient creates a new HintedHandoffHintsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHintedHandoffHintsGetParamsWithHTTPClient(client *http.Client) *HintedHandoffHintsGetParams {

	return &HintedHandoffHintsGetParams{
		HTTPClient: client,
	}
}

/*HintedHandoffHintsGetParams contains all the parameters to send to the API endpoint
for the hinted handoff hints get operation typically these are written to a http.Request
*/
type HintedHandoffHintsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) WithTimeout(timeout time.Duration) *HintedHandoffHintsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) WithContext(ctx context.Context) *HintedHandoffHintsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) WithHTTPClient(client *http.Client) *HintedHandoffHintsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hinted handoff hints get params
func (o *HintedHandoffHintsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HintedHandoffHintsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
