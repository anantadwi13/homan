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

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// NewCreateServerSwitchingRuleParams creates a new CreateServerSwitchingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServerSwitchingRuleParams() *CreateServerSwitchingRuleParams {
	return &CreateServerSwitchingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServerSwitchingRuleParamsWithTimeout creates a new CreateServerSwitchingRuleParams object
// with the ability to set a timeout on a request.
func NewCreateServerSwitchingRuleParamsWithTimeout(timeout time.Duration) *CreateServerSwitchingRuleParams {
	return &CreateServerSwitchingRuleParams{
		timeout: timeout,
	}
}

// NewCreateServerSwitchingRuleParamsWithContext creates a new CreateServerSwitchingRuleParams object
// with the ability to set a context for a request.
func NewCreateServerSwitchingRuleParamsWithContext(ctx context.Context) *CreateServerSwitchingRuleParams {
	return &CreateServerSwitchingRuleParams{
		Context: ctx,
	}
}

// NewCreateServerSwitchingRuleParamsWithHTTPClient creates a new CreateServerSwitchingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServerSwitchingRuleParamsWithHTTPClient(client *http.Client) *CreateServerSwitchingRuleParams {
	return &CreateServerSwitchingRuleParams{
		HTTPClient: client,
	}
}

/* CreateServerSwitchingRuleParams contains all the parameters to send to the API endpoint
   for the create server switching rule operation.

   Typically these are written to a http.Request.
*/
type CreateServerSwitchingRuleParams struct {

	/* Backend.

	   Backend name
	*/
	Backend string

	// Data.
	Data *models.ServerSwitchingRule

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

// WithDefaults hydrates default values in the create server switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerSwitchingRuleParams) WithDefaults() *CreateServerSwitchingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create server switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerSwitchingRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithTimeout(timeout time.Duration) *CreateServerSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithContext(ctx context.Context) *CreateServerSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithHTTPClient(client *http.Client) *CreateServerSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithBackend(backend string) *CreateServerSwitchingRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithData(data *models.ServerSwitchingRule) *CreateServerSwitchingRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetData(data *models.ServerSwitchingRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithForceReload(forceReload *bool) *CreateServerSwitchingRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithTransactionID(transactionID *string) *CreateServerSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) WithVersion(version *int64) *CreateServerSwitchingRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create server switching rule params
func (o *CreateServerSwitchingRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServerSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
