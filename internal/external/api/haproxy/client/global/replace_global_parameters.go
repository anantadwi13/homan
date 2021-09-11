// Code generated by go-swagger; DO NOT EDIT.

package global

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

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// NewReplaceGlobalParams creates a new ReplaceGlobalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceGlobalParams() *ReplaceGlobalParams {
	return &ReplaceGlobalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceGlobalParamsWithTimeout creates a new ReplaceGlobalParams object
// with the ability to set a timeout on a request.
func NewReplaceGlobalParamsWithTimeout(timeout time.Duration) *ReplaceGlobalParams {
	return &ReplaceGlobalParams{
		timeout: timeout,
	}
}

// NewReplaceGlobalParamsWithContext creates a new ReplaceGlobalParams object
// with the ability to set a context for a request.
func NewReplaceGlobalParamsWithContext(ctx context.Context) *ReplaceGlobalParams {
	return &ReplaceGlobalParams{
		Context: ctx,
	}
}

// NewReplaceGlobalParamsWithHTTPClient creates a new ReplaceGlobalParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceGlobalParamsWithHTTPClient(client *http.Client) *ReplaceGlobalParams {
	return &ReplaceGlobalParams{
		HTTPClient: client,
	}
}

/* ReplaceGlobalParams contains all the parameters to send to the API endpoint
   for the replace global operation.

   Typically these are written to a http.Request.
*/
type ReplaceGlobalParams struct {

	// Data.
	Data *models.Global

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace global params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceGlobalParams) WithDefaults() *ReplaceGlobalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace global params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceGlobalParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceGlobalParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace global params
func (o *ReplaceGlobalParams) WithTimeout(timeout time.Duration) *ReplaceGlobalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace global params
func (o *ReplaceGlobalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace global params
func (o *ReplaceGlobalParams) WithContext(ctx context.Context) *ReplaceGlobalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace global params
func (o *ReplaceGlobalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace global params
func (o *ReplaceGlobalParams) WithHTTPClient(client *http.Client) *ReplaceGlobalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace global params
func (o *ReplaceGlobalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace global params
func (o *ReplaceGlobalParams) WithData(data *models.Global) *ReplaceGlobalParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace global params
func (o *ReplaceGlobalParams) SetData(data *models.Global) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace global params
func (o *ReplaceGlobalParams) WithForceReload(forceReload *bool) *ReplaceGlobalParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace global params
func (o *ReplaceGlobalParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the replace global params
func (o *ReplaceGlobalParams) WithTransactionID(transactionID *string) *ReplaceGlobalParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace global params
func (o *ReplaceGlobalParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace global params
func (o *ReplaceGlobalParams) WithVersion(version *int64) *ReplaceGlobalParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace global params
func (o *ReplaceGlobalParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceGlobalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}