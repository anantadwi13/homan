// Code generated by go-swagger; DO NOT EDIT.

package log_target

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

// NewCreateLogTargetParams creates a new CreateLogTargetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLogTargetParams() *CreateLogTargetParams {
	return &CreateLogTargetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLogTargetParamsWithTimeout creates a new CreateLogTargetParams object
// with the ability to set a timeout on a request.
func NewCreateLogTargetParamsWithTimeout(timeout time.Duration) *CreateLogTargetParams {
	return &CreateLogTargetParams{
		timeout: timeout,
	}
}

// NewCreateLogTargetParamsWithContext creates a new CreateLogTargetParams object
// with the ability to set a context for a request.
func NewCreateLogTargetParamsWithContext(ctx context.Context) *CreateLogTargetParams {
	return &CreateLogTargetParams{
		Context: ctx,
	}
}

// NewCreateLogTargetParamsWithHTTPClient creates a new CreateLogTargetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLogTargetParamsWithHTTPClient(client *http.Client) *CreateLogTargetParams {
	return &CreateLogTargetParams{
		HTTPClient: client,
	}
}

/* CreateLogTargetParams contains all the parameters to send to the API endpoint
   for the create log target operation.

   Typically these are written to a http.Request.
*/
type CreateLogTargetParams struct {

	// Data.
	Data *models.LogTarget

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType string

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

// WithDefaults hydrates default values in the create log target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogTargetParams) WithDefaults() *CreateLogTargetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create log target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogTargetParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateLogTargetParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create log target params
func (o *CreateLogTargetParams) WithTimeout(timeout time.Duration) *CreateLogTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create log target params
func (o *CreateLogTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create log target params
func (o *CreateLogTargetParams) WithContext(ctx context.Context) *CreateLogTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create log target params
func (o *CreateLogTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create log target params
func (o *CreateLogTargetParams) WithHTTPClient(client *http.Client) *CreateLogTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create log target params
func (o *CreateLogTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create log target params
func (o *CreateLogTargetParams) WithData(data *models.LogTarget) *CreateLogTargetParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create log target params
func (o *CreateLogTargetParams) SetData(data *models.LogTarget) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create log target params
func (o *CreateLogTargetParams) WithForceReload(forceReload *bool) *CreateLogTargetParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create log target params
func (o *CreateLogTargetParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create log target params
func (o *CreateLogTargetParams) WithParentName(parentName *string) *CreateLogTargetParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create log target params
func (o *CreateLogTargetParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create log target params
func (o *CreateLogTargetParams) WithParentType(parentType string) *CreateLogTargetParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create log target params
func (o *CreateLogTargetParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create log target params
func (o *CreateLogTargetParams) WithTransactionID(transactionID *string) *CreateLogTargetParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create log target params
func (o *CreateLogTargetParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create log target params
func (o *CreateLogTargetParams) WithVersion(version *int64) *CreateLogTargetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create log target params
func (o *CreateLogTargetParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLogTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ParentName != nil {

		// query param parent_name
		var qrParentName string

		if o.ParentName != nil {
			qrParentName = *o.ParentName
		}
		qParentName := qrParentName
		if qParentName != "" {

			if err := r.SetQueryParam("parent_name", qParentName); err != nil {
				return err
			}
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
			return err
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
