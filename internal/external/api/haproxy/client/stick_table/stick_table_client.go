// Code generated by go-swagger; DO NOT EDIT.

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stick table API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stick table API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetStickTable(params *GetStickTableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTableOK, error)

	GetStickTableEntries(params *GetStickTableEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTableEntriesOK, error)

	GetStickTables(params *GetStickTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTablesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStickTable returns stick table

  Returns one stick table from runtime.
*/
func (a *Client) GetStickTable(params *GetStickTableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStickTableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStickTable",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/stick_tables/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStickTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStickTableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStickTableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetStickTableEntries returns stick table entries

  Returns an array of all entries in a given stick tables.
*/
func (a *Client) GetStickTableEntries(params *GetStickTableEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTableEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStickTableEntriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStickTableEntries",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/stick_table_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStickTableEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStickTableEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStickTableEntriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetStickTables returns stick tables

  Returns an array of all stick tables.
*/
func (a *Client) GetStickTables(params *GetStickTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStickTablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStickTablesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStickTables",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/stick_tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStickTablesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStickTablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStickTablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}