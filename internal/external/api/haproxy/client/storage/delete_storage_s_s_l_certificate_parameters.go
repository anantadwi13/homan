// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewDeleteStorageSSLCertificateParams creates a new DeleteStorageSSLCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteStorageSSLCertificateParams() *DeleteStorageSSLCertificateParams {
	return &DeleteStorageSSLCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStorageSSLCertificateParamsWithTimeout creates a new DeleteStorageSSLCertificateParams object
// with the ability to set a timeout on a request.
func NewDeleteStorageSSLCertificateParamsWithTimeout(timeout time.Duration) *DeleteStorageSSLCertificateParams {
	return &DeleteStorageSSLCertificateParams{
		timeout: timeout,
	}
}

// NewDeleteStorageSSLCertificateParamsWithContext creates a new DeleteStorageSSLCertificateParams object
// with the ability to set a context for a request.
func NewDeleteStorageSSLCertificateParamsWithContext(ctx context.Context) *DeleteStorageSSLCertificateParams {
	return &DeleteStorageSSLCertificateParams{
		Context: ctx,
	}
}

// NewDeleteStorageSSLCertificateParamsWithHTTPClient creates a new DeleteStorageSSLCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteStorageSSLCertificateParamsWithHTTPClient(client *http.Client) *DeleteStorageSSLCertificateParams {
	return &DeleteStorageSSLCertificateParams{
		HTTPClient: client,
	}
}

/* DeleteStorageSSLCertificateParams contains all the parameters to send to the API endpoint
   for the delete storage s s l certificate operation.

   Typically these are written to a http.Request.
*/
type DeleteStorageSSLCertificateParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   SSL certificate name
	*/
	Name string

	/* SkipReload.

	   If set, no reload will be initiated after update
	*/
	SkipReload *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageSSLCertificateParams) WithDefaults() *DeleteStorageSSLCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageSSLCertificateParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)

		skipReloadDefault = bool(false)
	)

	val := DeleteStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,
		SkipReload:  &skipReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithTimeout(timeout time.Duration) *DeleteStorageSSLCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithContext(ctx context.Context) *DeleteStorageSSLCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithHTTPClient(client *http.Client) *DeleteStorageSSLCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithForceReload(forceReload *bool) *DeleteStorageSSLCertificateParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithName(name string) *DeleteStorageSSLCertificateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetName(name string) {
	o.Name = name
}

// WithSkipReload adds the skipReload to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) WithSkipReload(skipReload *bool) *DeleteStorageSSLCertificateParams {
	o.SetSkipReload(skipReload)
	return o
}

// SetSkipReload adds the skipReload to the delete storage s s l certificate params
func (o *DeleteStorageSSLCertificateParams) SetSkipReload(skipReload *bool) {
	o.SkipReload = skipReload
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStorageSSLCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool

		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {

			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.SkipReload != nil {

		// query param skip_reload
		var qrSkipReload bool

		if o.SkipReload != nil {
			qrSkipReload = *o.SkipReload
		}
		qSkipReload := swag.FormatBool(qrSkipReload)
		if qSkipReload != "" {

			if err := r.SetQueryParam("skip_reload", qSkipReload); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
