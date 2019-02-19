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

// NewMessagingServiceMessagesSentGetParams creates a new MessagingServiceMessagesSentGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesSentGetParams() *MessagingServiceMessagesSentGetParams {

	return &MessagingServiceMessagesSentGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesSentGetParamsWithTimeout creates a new MessagingServiceMessagesSentGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesSentGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesSentGetParams {

	return &MessagingServiceMessagesSentGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesSentGetParamsWithContext creates a new MessagingServiceMessagesSentGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesSentGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesSentGetParams {

	return &MessagingServiceMessagesSentGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesSentGetParamsWithHTTPClient creates a new MessagingServiceMessagesSentGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesSentGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesSentGetParams {

	return &MessagingServiceMessagesSentGetParams{
		HTTPClient: client,
	}
}

/*MessagingServiceMessagesSentGetParams contains all the parameters to send to the API endpoint
for the messaging service messages sent get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesSentGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesSentGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesSentGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesSentGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages sent get params
func (o *MessagingServiceMessagesSentGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesSentGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
