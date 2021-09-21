// Code generated by go-swagger; DO NOT EDIT.

package discovery

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
)

// NewGetServicesEndpointsParams creates a new GetServicesEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServicesEndpointsParams() *GetServicesEndpointsParams {
	return &GetServicesEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesEndpointsParamsWithTimeout creates a new GetServicesEndpointsParams object
// with the ability to set a timeout on a request.
func NewGetServicesEndpointsParamsWithTimeout(timeout time.Duration) *GetServicesEndpointsParams {
	return &GetServicesEndpointsParams{
		timeout: timeout,
	}
}

// NewGetServicesEndpointsParamsWithContext creates a new GetServicesEndpointsParams object
// with the ability to set a context for a request.
func NewGetServicesEndpointsParamsWithContext(ctx context.Context) *GetServicesEndpointsParams {
	return &GetServicesEndpointsParams{
		Context: ctx,
	}
}

// NewGetServicesEndpointsParamsWithHTTPClient creates a new GetServicesEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServicesEndpointsParamsWithHTTPClient(client *http.Client) *GetServicesEndpointsParams {
	return &GetServicesEndpointsParams{
		HTTPClient: client,
	}
}

/* GetServicesEndpointsParams contains all the parameters to send to the API endpoint
   for the get services endpoints operation.

   Typically these are written to a http.Request.
*/
type GetServicesEndpointsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get services endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServicesEndpointsParams) WithDefaults() *GetServicesEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get services endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServicesEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get services endpoints params
func (o *GetServicesEndpointsParams) WithTimeout(timeout time.Duration) *GetServicesEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services endpoints params
func (o *GetServicesEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services endpoints params
func (o *GetServicesEndpointsParams) WithContext(ctx context.Context) *GetServicesEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services endpoints params
func (o *GetServicesEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services endpoints params
func (o *GetServicesEndpointsParams) WithHTTPClient(client *http.Client) *GetServicesEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services endpoints params
func (o *GetServicesEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}