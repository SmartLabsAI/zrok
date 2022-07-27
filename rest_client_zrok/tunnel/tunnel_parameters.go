// Code generated by go-swagger; DO NOT EDIT.

package tunnel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// NewTunnelParams creates a new TunnelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTunnelParams() *TunnelParams {
	return &TunnelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTunnelParamsWithTimeout creates a new TunnelParams object
// with the ability to set a timeout on a request.
func NewTunnelParamsWithTimeout(timeout time.Duration) *TunnelParams {
	return &TunnelParams{
		timeout: timeout,
	}
}

// NewTunnelParamsWithContext creates a new TunnelParams object
// with the ability to set a context for a request.
func NewTunnelParamsWithContext(ctx context.Context) *TunnelParams {
	return &TunnelParams{
		Context: ctx,
	}
}

// NewTunnelParamsWithHTTPClient creates a new TunnelParams object
// with the ability to set a custom HTTPClient for a request.
func NewTunnelParamsWithHTTPClient(client *http.Client) *TunnelParams {
	return &TunnelParams{
		HTTPClient: client,
	}
}

/* TunnelParams contains all the parameters to send to the API endpoint
   for the tunnel operation.

   Typically these are written to a http.Request.
*/
type TunnelParams struct {

	// Body.
	Body *rest_model_zrok.TunnelRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tunnel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TunnelParams) WithDefaults() *TunnelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tunnel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TunnelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tunnel params
func (o *TunnelParams) WithTimeout(timeout time.Duration) *TunnelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tunnel params
func (o *TunnelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tunnel params
func (o *TunnelParams) WithContext(ctx context.Context) *TunnelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tunnel params
func (o *TunnelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tunnel params
func (o *TunnelParams) WithHTTPClient(client *http.Client) *TunnelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tunnel params
func (o *TunnelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tunnel params
func (o *TunnelParams) WithBody(body *rest_model_zrok.TunnelRequest) *TunnelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tunnel params
func (o *TunnelParams) SetBody(body *rest_model_zrok.TunnelRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TunnelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}