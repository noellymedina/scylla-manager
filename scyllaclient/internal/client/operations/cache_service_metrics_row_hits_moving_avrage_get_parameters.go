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

// NewCacheServiceMetricsRowHitsMovingAvrageGetParams creates a new CacheServiceMetricsRowHitsMovingAvrageGetParams object
// with the default values initialized.
func NewCacheServiceMetricsRowHitsMovingAvrageGetParams() *CacheServiceMetricsRowHitsMovingAvrageGetParams {

	return &CacheServiceMetricsRowHitsMovingAvrageGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithTimeout creates a new CacheServiceMetricsRowHitsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowHitsMovingAvrageGetParams {

	return &CacheServiceMetricsRowHitsMovingAvrageGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithContext creates a new CacheServiceMetricsRowHitsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowHitsMovingAvrageGetParams {

	return &CacheServiceMetricsRowHitsMovingAvrageGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithHTTPClient creates a new CacheServiceMetricsRowHitsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsRowHitsMovingAvrageGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowHitsMovingAvrageGetParams {

	return &CacheServiceMetricsRowHitsMovingAvrageGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsRowHitsMovingAvrageGetParams contains all the parameters to send to the API endpoint
for the cache service metrics row hits moving avrage get operation typically these are written to a http.Request
*/
type CacheServiceMetricsRowHitsMovingAvrageGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowHitsMovingAvrageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowHitsMovingAvrageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowHitsMovingAvrageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row hits moving avrage get params
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowHitsMovingAvrageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
