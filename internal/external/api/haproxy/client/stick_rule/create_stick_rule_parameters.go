// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

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

// NewCreateStickRuleParams creates a new CreateStickRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStickRuleParams() *CreateStickRuleParams {
	return &CreateStickRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStickRuleParamsWithTimeout creates a new CreateStickRuleParams object
// with the ability to set a timeout on a request.
func NewCreateStickRuleParamsWithTimeout(timeout time.Duration) *CreateStickRuleParams {
	return &CreateStickRuleParams{
		timeout: timeout,
	}
}

// NewCreateStickRuleParamsWithContext creates a new CreateStickRuleParams object
// with the ability to set a context for a request.
func NewCreateStickRuleParamsWithContext(ctx context.Context) *CreateStickRuleParams {
	return &CreateStickRuleParams{
		Context: ctx,
	}
}

// NewCreateStickRuleParamsWithHTTPClient creates a new CreateStickRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStickRuleParamsWithHTTPClient(client *http.Client) *CreateStickRuleParams {
	return &CreateStickRuleParams{
		HTTPClient: client,
	}
}

/* CreateStickRuleParams contains all the parameters to send to the API endpoint
   for the create stick rule operation.

   Typically these are written to a http.Request.
*/
type CreateStickRuleParams struct {

	/* Backend.

	   Backend name
	*/
	Backend string

	// Data.
	Data *models.StickRule

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

// WithDefaults hydrates default values in the create stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStickRuleParams) WithDefaults() *CreateStickRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStickRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateStickRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create stick rule params
func (o *CreateStickRuleParams) WithTimeout(timeout time.Duration) *CreateStickRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create stick rule params
func (o *CreateStickRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create stick rule params
func (o *CreateStickRuleParams) WithContext(ctx context.Context) *CreateStickRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create stick rule params
func (o *CreateStickRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create stick rule params
func (o *CreateStickRuleParams) WithHTTPClient(client *http.Client) *CreateStickRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create stick rule params
func (o *CreateStickRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the create stick rule params
func (o *CreateStickRuleParams) WithBackend(backend string) *CreateStickRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the create stick rule params
func (o *CreateStickRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the create stick rule params
func (o *CreateStickRuleParams) WithData(data *models.StickRule) *CreateStickRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create stick rule params
func (o *CreateStickRuleParams) SetData(data *models.StickRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create stick rule params
func (o *CreateStickRuleParams) WithForceReload(forceReload *bool) *CreateStickRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create stick rule params
func (o *CreateStickRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the create stick rule params
func (o *CreateStickRuleParams) WithTransactionID(transactionID *string) *CreateStickRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create stick rule params
func (o *CreateStickRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create stick rule params
func (o *CreateStickRuleParams) WithVersion(version *int64) *CreateStickRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create stick rule params
func (o *CreateStickRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStickRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
