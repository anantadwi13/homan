// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

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

// NewDeleteServerSwitchingRuleParams creates a new DeleteServerSwitchingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteServerSwitchingRuleParams() *DeleteServerSwitchingRuleParams {
	return &DeleteServerSwitchingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServerSwitchingRuleParamsWithTimeout creates a new DeleteServerSwitchingRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteServerSwitchingRuleParamsWithTimeout(timeout time.Duration) *DeleteServerSwitchingRuleParams {
	return &DeleteServerSwitchingRuleParams{
		timeout: timeout,
	}
}

// NewDeleteServerSwitchingRuleParamsWithContext creates a new DeleteServerSwitchingRuleParams object
// with the ability to set a context for a request.
func NewDeleteServerSwitchingRuleParamsWithContext(ctx context.Context) *DeleteServerSwitchingRuleParams {
	return &DeleteServerSwitchingRuleParams{
		Context: ctx,
	}
}

// NewDeleteServerSwitchingRuleParamsWithHTTPClient creates a new DeleteServerSwitchingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteServerSwitchingRuleParamsWithHTTPClient(client *http.Client) *DeleteServerSwitchingRuleParams {
	return &DeleteServerSwitchingRuleParams{
		HTTPClient: client,
	}
}

/* DeleteServerSwitchingRuleParams contains all the parameters to send to the API endpoint
   for the delete server switching rule operation.

   Typically these are written to a http.Request.
*/
type DeleteServerSwitchingRuleParams struct {

	/* Backend.

	   Backend name
	*/
	Backend string

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   Switching Rule Index
	*/
	Index int64

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

// WithDefaults hydrates default values in the delete server switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServerSwitchingRuleParams) WithDefaults() *DeleteServerSwitchingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete server switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServerSwitchingRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithTimeout(timeout time.Duration) *DeleteServerSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithContext(ctx context.Context) *DeleteServerSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithHTTPClient(client *http.Client) *DeleteServerSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithBackend(backend string) *DeleteServerSwitchingRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithForceReload adds the forceReload to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithForceReload(forceReload *bool) *DeleteServerSwitchingRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithIndex(index int64) *DeleteServerSwitchingRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithTransactionID(transactionID *string) *DeleteServerSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) WithVersion(version *int64) *DeleteServerSwitchingRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete server switching rule params
func (o *DeleteServerSwitchingRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServerSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {

		if err := r.SetQueryParam("backend", qBackend); err != nil {
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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
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
