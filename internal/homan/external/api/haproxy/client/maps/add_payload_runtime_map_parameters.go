// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAddPayloadRuntimeMapParams creates a new AddPayloadRuntimeMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPayloadRuntimeMapParams() *AddPayloadRuntimeMapParams {
	return &AddPayloadRuntimeMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPayloadRuntimeMapParamsWithTimeout creates a new AddPayloadRuntimeMapParams object
// with the ability to set a timeout on a request.
func NewAddPayloadRuntimeMapParamsWithTimeout(timeout time.Duration) *AddPayloadRuntimeMapParams {
	return &AddPayloadRuntimeMapParams{
		timeout: timeout,
	}
}

// NewAddPayloadRuntimeMapParamsWithContext creates a new AddPayloadRuntimeMapParams object
// with the ability to set a context for a request.
func NewAddPayloadRuntimeMapParamsWithContext(ctx context.Context) *AddPayloadRuntimeMapParams {
	return &AddPayloadRuntimeMapParams{
		Context: ctx,
	}
}

// NewAddPayloadRuntimeMapParamsWithHTTPClient creates a new AddPayloadRuntimeMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPayloadRuntimeMapParamsWithHTTPClient(client *http.Client) *AddPayloadRuntimeMapParams {
	return &AddPayloadRuntimeMapParams{
		HTTPClient: client,
	}
}

/* AddPayloadRuntimeMapParams contains all the parameters to send to the API endpoint
   for the add payload runtime map operation.

   Typically these are written to a http.Request.
*/
type AddPayloadRuntimeMapParams struct {

	// Data.
	Data models.MapEntries

	/* ForceSync.

	   If true, immediately syncs changes to disk
	*/
	ForceSync *bool

	/* Name.

	   Map file name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add payload runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayloadRuntimeMapParams) WithDefaults() *AddPayloadRuntimeMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add payload runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayloadRuntimeMapParams) SetDefaults() {
	var (
		forceSyncDefault = bool(false)
	)

	val := AddPayloadRuntimeMapParams{
		ForceSync: &forceSyncDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithTimeout(timeout time.Duration) *AddPayloadRuntimeMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithContext(ctx context.Context) *AddPayloadRuntimeMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithHTTPClient(client *http.Client) *AddPayloadRuntimeMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithData(data models.MapEntries) *AddPayloadRuntimeMapParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetData(data models.MapEntries) {
	o.Data = data
}

// WithForceSync adds the forceSync to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithForceSync(forceSync *bool) *AddPayloadRuntimeMapParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithName adds the name to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) WithName(name string) *AddPayloadRuntimeMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the add payload runtime map params
func (o *AddPayloadRuntimeMapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AddPayloadRuntimeMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.ForceSync != nil {

		// query param force_sync
		var qrForceSync bool

		if o.ForceSync != nil {
			qrForceSync = *o.ForceSync
		}
		qForceSync := swag.FormatBool(qrForceSync)
		if qForceSync != "" {

			if err := r.SetQueryParam("force_sync", qForceSync); err != nil {
				return err
			}
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
