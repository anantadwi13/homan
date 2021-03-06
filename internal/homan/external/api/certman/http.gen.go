// Package certman provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package certman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// CertificateRes defines model for certificate-res.
type CertificateRes struct {
	// List domains (Main + Alternative domains)
	Domains    []string  `json:"domains"`
	ExpiryDate time.Time `json:"expiry_date"`
	KeyType    string    `json:"key_type"`

	// Certificate Name (Main domain)
	Name        string `json:"name"`
	PrivateCert string `json:"private_cert"`
	PublicCert  string `json:"public_cert"`

	// Serial Number
	SerialNumber string `json:"serial_number"`
}

// GeneralRes defines model for general-res.
type GeneralRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BadRequest defines model for bad-request.
type BadRequest GeneralRes

// DefaultError defines model for default-error.
type DefaultError GeneralRes

// NotFound defines model for not-found.
type NotFound GeneralRes

// CreateCertificateJSONBody defines parameters for CreateCertificate.
type CreateCertificateJSONBody struct {
	// Alternative Domain
	AltDomains *[]string `json:"alt_domains,omitempty"`

	// Domain
	Domain string `json:"domain"`

	// Email address
	Email string `json:"email"`
}

// CreateCertificateJSONRequestBody defines body for CreateCertificate for application/json ContentType.
type CreateCertificateJSONRequestBody CreateCertificateJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAllCertificates request
	GetAllCertificates(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateCertificate request with any body
	CreateCertificateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateCertificate(ctx context.Context, body CreateCertificateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RenewAllCertificates request
	RenewAllCertificates(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteCertificate request
	DeleteCertificate(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetCertificateByDomain request
	GetCertificateByDomain(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RenewCertificateByDomain request
	RenewCertificateByDomain(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllCertificates(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllCertificatesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCertificateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCertificateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCertificate(ctx context.Context, body CreateCertificateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCertificateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RenewAllCertificates(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRenewAllCertificatesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteCertificate(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteCertificateRequest(c.Server, domain)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetCertificateByDomain(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCertificateByDomainRequest(c.Server, domain)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RenewCertificateByDomain(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRenewCertificateByDomainRequest(c.Server, domain)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllCertificatesRequest generates requests for GetAllCertificates
func NewGetAllCertificatesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateCertificateRequest calls the generic CreateCertificate builder with application/json body
func NewCreateCertificateRequest(server string, body CreateCertificateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateCertificateRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateCertificateRequestWithBody generates requests for CreateCertificate with any type of body
func NewCreateCertificateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRenewAllCertificatesRequest generates requests for RenewAllCertificates
func NewRenewAllCertificatesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteCertificateRequest generates requests for DeleteCertificate
func NewDeleteCertificateRequest(server string, domain string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetCertificateByDomainRequest generates requests for GetCertificateByDomain
func NewGetCertificateByDomainRequest(server string, domain string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRenewCertificateByDomainRequest generates requests for RenewCertificateByDomain
func NewRenewCertificateByDomainRequest(server string, domain string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/certificates/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAllCertificates request
	GetAllCertificatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllCertificatesResponse, error)

	// CreateCertificate request with any body
	CreateCertificateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCertificateResponse, error)

	CreateCertificateWithResponse(ctx context.Context, body CreateCertificateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCertificateResponse, error)

	// RenewAllCertificates request
	RenewAllCertificatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RenewAllCertificatesResponse, error)

	// DeleteCertificate request
	DeleteCertificateWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*DeleteCertificateResponse, error)

	// GetCertificateByDomain request
	GetCertificateByDomainWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*GetCertificateByDomainResponse, error)

	// RenewCertificateByDomain request
	RenewCertificateByDomainWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*RenewCertificateByDomainResponse, error)
}

type GetAllCertificatesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]CertificateRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r GetAllCertificatesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllCertificatesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateCertificateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CertificateRes
	JSON400      *GeneralRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r CreateCertificateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateCertificateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RenewAllCertificatesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GeneralRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r RenewAllCertificatesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RenewAllCertificatesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteCertificateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GeneralRes
	JSON404      *GeneralRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r DeleteCertificateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteCertificateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCertificateByDomainResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CertificateRes
	JSON404      *GeneralRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r GetCertificateByDomainResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCertificateByDomainResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RenewCertificateByDomainResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GeneralRes
	JSON404      *GeneralRes
	JSONDefault  *GeneralRes
}

// Status returns HTTPResponse.Status
func (r RenewCertificateByDomainResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RenewCertificateByDomainResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllCertificatesWithResponse request returning *GetAllCertificatesResponse
func (c *ClientWithResponses) GetAllCertificatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllCertificatesResponse, error) {
	rsp, err := c.GetAllCertificates(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllCertificatesResponse(rsp)
}

// CreateCertificateWithBodyWithResponse request with arbitrary body returning *CreateCertificateResponse
func (c *ClientWithResponses) CreateCertificateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCertificateResponse, error) {
	rsp, err := c.CreateCertificateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateCertificateResponse(rsp)
}

func (c *ClientWithResponses) CreateCertificateWithResponse(ctx context.Context, body CreateCertificateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCertificateResponse, error) {
	rsp, err := c.CreateCertificate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateCertificateResponse(rsp)
}

// RenewAllCertificatesWithResponse request returning *RenewAllCertificatesResponse
func (c *ClientWithResponses) RenewAllCertificatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RenewAllCertificatesResponse, error) {
	rsp, err := c.RenewAllCertificates(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRenewAllCertificatesResponse(rsp)
}

// DeleteCertificateWithResponse request returning *DeleteCertificateResponse
func (c *ClientWithResponses) DeleteCertificateWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*DeleteCertificateResponse, error) {
	rsp, err := c.DeleteCertificate(ctx, domain, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteCertificateResponse(rsp)
}

// GetCertificateByDomainWithResponse request returning *GetCertificateByDomainResponse
func (c *ClientWithResponses) GetCertificateByDomainWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*GetCertificateByDomainResponse, error) {
	rsp, err := c.GetCertificateByDomain(ctx, domain, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCertificateByDomainResponse(rsp)
}

// RenewCertificateByDomainWithResponse request returning *RenewCertificateByDomainResponse
func (c *ClientWithResponses) RenewCertificateByDomainWithResponse(ctx context.Context, domain string, reqEditors ...RequestEditorFn) (*RenewCertificateByDomainResponse, error) {
	rsp, err := c.RenewCertificateByDomain(ctx, domain, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRenewCertificateByDomainResponse(rsp)
}

// ParseGetAllCertificatesResponse parses an HTTP response from a GetAllCertificatesWithResponse call
func ParseGetAllCertificatesResponse(rsp *http.Response) (*GetAllCertificatesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetAllCertificatesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []CertificateRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateCertificateResponse parses an HTTP response from a CreateCertificateWithResponse call
func ParseCreateCertificateResponse(rsp *http.Response) (*CreateCertificateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateCertificateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CertificateRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseRenewAllCertificatesResponse parses an HTTP response from a RenewAllCertificatesWithResponse call
func ParseRenewAllCertificatesResponse(rsp *http.Response) (*RenewAllCertificatesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &RenewAllCertificatesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteCertificateResponse parses an HTTP response from a DeleteCertificateWithResponse call
func ParseDeleteCertificateResponse(rsp *http.Response) (*DeleteCertificateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteCertificateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetCertificateByDomainResponse parses an HTTP response from a GetCertificateByDomainWithResponse call
func ParseGetCertificateByDomainResponse(rsp *http.Response) (*GetCertificateByDomainResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetCertificateByDomainResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CertificateRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseRenewCertificateByDomainResponse parses an HTTP response from a RenewCertificateByDomainWithResponse call
func ParseRenewCertificateByDomainResponse(rsp *http.Response) (*RenewCertificateByDomainResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &RenewCertificateByDomainResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GeneralRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
