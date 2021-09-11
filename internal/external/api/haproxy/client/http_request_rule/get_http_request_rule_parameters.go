// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

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

// NewGetHTTPRequestRuleParams creates a new GetHTTPRequestRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHTTPRequestRuleParams() *GetHTTPRequestRuleParams {
	return &GetHTTPRequestRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHTTPRequestRuleParamsWithTimeout creates a new GetHTTPRequestRuleParams object
// with the ability to set a timeout on a request.
func NewGetHTTPRequestRuleParamsWithTimeout(timeout time.Duration) *GetHTTPRequestRuleParams {
	return &GetHTTPRequestRuleParams{
		timeout: timeout,
	}
}

// NewGetHTTPRequestRuleParamsWithContext creates a new GetHTTPRequestRuleParams object
// with the ability to set a context for a request.
func NewGetHTTPRequestRuleParamsWithContext(ctx context.Context) *GetHTTPRequestRuleParams {
	return &GetHTTPRequestRuleParams{
		Context: ctx,
	}
}

// NewGetHTTPRequestRuleParamsWithHTTPClient creates a new GetHTTPRequestRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHTTPRequestRuleParamsWithHTTPClient(client *http.Client) *GetHTTPRequestRuleParams {
	return &GetHTTPRequestRuleParams{
		HTTPClient: client,
	}
}

/* GetHTTPRequestRuleParams contains all the parameters to send to the API endpoint
   for the get HTTP request rule operation.

   Typically these are written to a http.Request.
*/
type GetHTTPRequestRuleParams struct {

	/* Index.

	   HTTP Request Rule Index
	*/
	Index int64

	/* ParentName.

	   Parent name
	*/
	ParentName string

	/* ParentType.

	   Parent type
	*/
	ParentType string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPRequestRuleParams) WithDefaults() *GetHTTPRequestRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPRequestRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithTimeout(timeout time.Duration) *GetHTTPRequestRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithContext(ctx context.Context) *GetHTTPRequestRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithHTTPClient(client *http.Client) *GetHTTPRequestRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithIndex(index int64) *GetHTTPRequestRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithParentName(parentName string) *GetHTTPRequestRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithParentType(parentType string) *GetHTTPRequestRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) WithTransactionID(transactionID *string) *GetHTTPRequestRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get HTTP request rule params
func (o *GetHTTPRequestRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHTTPRequestRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {

		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
