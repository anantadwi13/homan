// Code generated by go-swagger; DO NOT EDIT.

package spoe

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

// NewCreateSpoeScopeParams creates a new CreateSpoeScopeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSpoeScopeParams() *CreateSpoeScopeParams {
	return &CreateSpoeScopeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSpoeScopeParamsWithTimeout creates a new CreateSpoeScopeParams object
// with the ability to set a timeout on a request.
func NewCreateSpoeScopeParamsWithTimeout(timeout time.Duration) *CreateSpoeScopeParams {
	return &CreateSpoeScopeParams{
		timeout: timeout,
	}
}

// NewCreateSpoeScopeParamsWithContext creates a new CreateSpoeScopeParams object
// with the ability to set a context for a request.
func NewCreateSpoeScopeParamsWithContext(ctx context.Context) *CreateSpoeScopeParams {
	return &CreateSpoeScopeParams{
		Context: ctx,
	}
}

// NewCreateSpoeScopeParamsWithHTTPClient creates a new CreateSpoeScopeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSpoeScopeParamsWithHTTPClient(client *http.Client) *CreateSpoeScopeParams {
	return &CreateSpoeScopeParams{
		HTTPClient: client,
	}
}

/* CreateSpoeScopeParams contains all the parameters to send to the API endpoint
   for the create spoe scope operation.

   Typically these are written to a http.Request.
*/
type CreateSpoeScopeParams struct {

	// Data.
	Data models.SpoeScope

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

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

// WithDefaults hydrates default values in the create spoe scope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSpoeScopeParams) WithDefaults() *CreateSpoeScopeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create spoe scope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSpoeScopeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create spoe scope params
func (o *CreateSpoeScopeParams) WithTimeout(timeout time.Duration) *CreateSpoeScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create spoe scope params
func (o *CreateSpoeScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create spoe scope params
func (o *CreateSpoeScopeParams) WithContext(ctx context.Context) *CreateSpoeScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create spoe scope params
func (o *CreateSpoeScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create spoe scope params
func (o *CreateSpoeScopeParams) WithHTTPClient(client *http.Client) *CreateSpoeScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create spoe scope params
func (o *CreateSpoeScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create spoe scope params
func (o *CreateSpoeScopeParams) WithData(data models.SpoeScope) *CreateSpoeScopeParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create spoe scope params
func (o *CreateSpoeScopeParams) SetData(data models.SpoeScope) {
	o.Data = data
}

// WithSpoe adds the spoe to the create spoe scope params
func (o *CreateSpoeScopeParams) WithSpoe(spoe string) *CreateSpoeScopeParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the create spoe scope params
func (o *CreateSpoeScopeParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the create spoe scope params
func (o *CreateSpoeScopeParams) WithTransactionID(transactionID *string) *CreateSpoeScopeParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create spoe scope params
func (o *CreateSpoeScopeParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create spoe scope params
func (o *CreateSpoeScopeParams) WithVersion(version *int64) *CreateSpoeScopeParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create spoe scope params
func (o *CreateSpoeScopeParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSpoeScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
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
