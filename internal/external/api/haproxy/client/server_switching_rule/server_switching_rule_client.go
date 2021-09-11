// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server switching rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server switching rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServerSwitchingRule(params *CreateServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServerSwitchingRuleCreated, *CreateServerSwitchingRuleAccepted, error)

	DeleteServerSwitchingRule(params *DeleteServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServerSwitchingRuleAccepted, *DeleteServerSwitchingRuleNoContent, error)

	GetServerSwitchingRule(params *GetServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerSwitchingRuleOK, error)

	GetServerSwitchingRules(params *GetServerSwitchingRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerSwitchingRulesOK, error)

	ReplaceServerSwitchingRule(params *ReplaceServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceServerSwitchingRuleOK, *ReplaceServerSwitchingRuleAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateServerSwitchingRule adds a new server switching rule

  Adds a new Server Switching Rule of the specified type in the specified backend.
*/
func (a *Client) CreateServerSwitchingRule(params *CreateServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServerSwitchingRuleCreated, *CreateServerSwitchingRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServerSwitchingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createServerSwitchingRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/server_switching_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateServerSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateServerSwitchingRuleCreated:
		return value, nil, nil
	case *CreateServerSwitchingRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateServerSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteServerSwitchingRule deletes a server switching rule

  Deletes a Server Switching Rule configuration by it's index from the specified backend.
*/
func (a *Client) DeleteServerSwitchingRule(params *DeleteServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServerSwitchingRuleAccepted, *DeleteServerSwitchingRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServerSwitchingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteServerSwitchingRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/server_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServerSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteServerSwitchingRuleAccepted:
		return value, nil, nil
	case *DeleteServerSwitchingRuleNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServerSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetServerSwitchingRule returns one server switching rule

  Returns one Server Switching Rule configuration by it's index in the specified backend.
*/
func (a *Client) GetServerSwitchingRule(params *GetServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerSwitchingRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerSwitchingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServerSwitchingRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/server_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServerSwitchingRuleReader{formats: a.formats},
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
	success, ok := result.(*GetServerSwitchingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerSwitchingRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetServerSwitchingRules returns an array of all server switching rules

  Returns all Backend Switching Rules that are configured in specified backend.
*/
func (a *Client) GetServerSwitchingRules(params *GetServerSwitchingRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerSwitchingRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerSwitchingRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServerSwitchingRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/server_switching_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServerSwitchingRulesReader{formats: a.formats},
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
	success, ok := result.(*GetServerSwitchingRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerSwitchingRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceServerSwitchingRule replaces a server switching rule

  Replaces a Server Switching Rule configuration by it's index in the specified backend.
*/
func (a *Client) ReplaceServerSwitchingRule(params *ReplaceServerSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceServerSwitchingRuleOK, *ReplaceServerSwitchingRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceServerSwitchingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceServerSwitchingRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/server_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceServerSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceServerSwitchingRuleOK:
		return value, nil, nil
	case *ReplaceServerSwitchingRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceServerSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
