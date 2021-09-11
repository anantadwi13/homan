// Code generated by go-swagger; DO NOT EDIT.

package maps

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
	"github.com/go-openapi/swag"
)

// NewGetAllRuntimeMapFilesParams creates a new GetAllRuntimeMapFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllRuntimeMapFilesParams() *GetAllRuntimeMapFilesParams {
	return &GetAllRuntimeMapFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRuntimeMapFilesParamsWithTimeout creates a new GetAllRuntimeMapFilesParams object
// with the ability to set a timeout on a request.
func NewGetAllRuntimeMapFilesParamsWithTimeout(timeout time.Duration) *GetAllRuntimeMapFilesParams {
	return &GetAllRuntimeMapFilesParams{
		timeout: timeout,
	}
}

// NewGetAllRuntimeMapFilesParamsWithContext creates a new GetAllRuntimeMapFilesParams object
// with the ability to set a context for a request.
func NewGetAllRuntimeMapFilesParamsWithContext(ctx context.Context) *GetAllRuntimeMapFilesParams {
	return &GetAllRuntimeMapFilesParams{
		Context: ctx,
	}
}

// NewGetAllRuntimeMapFilesParamsWithHTTPClient creates a new GetAllRuntimeMapFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllRuntimeMapFilesParamsWithHTTPClient(client *http.Client) *GetAllRuntimeMapFilesParams {
	return &GetAllRuntimeMapFilesParams{
		HTTPClient: client,
	}
}

/* GetAllRuntimeMapFilesParams contains all the parameters to send to the API endpoint
   for the get all runtime map files operation.

   Typically these are written to a http.Request.
*/
type GetAllRuntimeMapFilesParams struct {

	/* IncludeUnmanaged.

	   If true, also show unmanaged map files loaded in haproxy
	*/
	IncludeUnmanaged *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all runtime map files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRuntimeMapFilesParams) WithDefaults() *GetAllRuntimeMapFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all runtime map files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRuntimeMapFilesParams) SetDefaults() {
	var (
		includeUnmanagedDefault = bool(false)
	)

	val := GetAllRuntimeMapFilesParams{
		IncludeUnmanaged: &includeUnmanagedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) WithTimeout(timeout time.Duration) *GetAllRuntimeMapFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) WithContext(ctx context.Context) *GetAllRuntimeMapFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) WithHTTPClient(client *http.Client) *GetAllRuntimeMapFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeUnmanaged adds the includeUnmanaged to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) WithIncludeUnmanaged(includeUnmanaged *bool) *GetAllRuntimeMapFilesParams {
	o.SetIncludeUnmanaged(includeUnmanaged)
	return o
}

// SetIncludeUnmanaged adds the includeUnmanaged to the get all runtime map files params
func (o *GetAllRuntimeMapFilesParams) SetIncludeUnmanaged(includeUnmanaged *bool) {
	o.IncludeUnmanaged = includeUnmanaged
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRuntimeMapFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeUnmanaged != nil {

		// query param include_unmanaged
		var qrIncludeUnmanaged bool

		if o.IncludeUnmanaged != nil {
			qrIncludeUnmanaged = *o.IncludeUnmanaged
		}
		qIncludeUnmanaged := swag.FormatBool(qrIncludeUnmanaged)
		if qIncludeUnmanaged != "" {

			if err := r.SetQueryParam("include_unmanaged", qIncludeUnmanaged); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
