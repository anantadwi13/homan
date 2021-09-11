// Code generated by go-swagger; DO NOT EDIT.

package information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetHaproxyProcessInfo(params *GetHaproxyProcessInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHaproxyProcessInfoOK, error)

	GetInfo(params *GetInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInfoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetHaproxyProcessInfo returns h a proxy process information

  Return HAProxy process information
*/
func (a *Client) GetHaproxyProcessInfo(params *GetHaproxyProcessInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHaproxyProcessInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHaproxyProcessInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHaproxyProcessInfo",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHaproxyProcessInfoReader{formats: a.formats},
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
	success, ok := result.(*GetHaproxyProcessInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHaproxyProcessInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetInfo returns API hardware and o s information

  Return API, hardware and OS information
*/
func (a *Client) GetInfo(params *GetInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInfo",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
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
	success, ok := result.(*GetInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}