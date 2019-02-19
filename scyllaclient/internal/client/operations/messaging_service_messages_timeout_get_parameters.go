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

// NewMessagingServiceMessagesTimeoutGetParams creates a new MessagingServiceMessagesTimeoutGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesTimeoutGetParams() *MessagingServiceMessagesTimeoutGetParams {

	return &MessagingServiceMessagesTimeoutGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesTimeoutGetParamsWithTimeout creates a new MessagingServiceMessagesTimeoutGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesTimeoutGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesTimeoutGetParams {

	return &MessagingServiceMessagesTimeoutGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesTimeoutGetParamsWithContext creates a new MessagingServiceMessagesTimeoutGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesTimeoutGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesTimeoutGetParams {

	return &MessagingServiceMessagesTimeoutGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesTimeoutGetParamsWithHTTPClient creates a new MessagingServiceMessagesTimeoutGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesTimeoutGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesTimeoutGetParams {

	return &MessagingServiceMessagesTimeoutGetParams{
		HTTPClient: client,
	}
}

/*MessagingServiceMessagesTimeoutGetParams contains all the parameters to send to the API endpoint
for the messaging service messages timeout get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesTimeoutGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesTimeoutGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesTimeoutGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesTimeoutGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages timeout get params
func (o *MessagingServiceMessagesTimeoutGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesTimeoutGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
