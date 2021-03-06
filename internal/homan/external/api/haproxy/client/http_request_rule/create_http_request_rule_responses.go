// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateHTTPRequestRuleReader is a Reader for the CreateHTTPRequestRule structure.
type CreateHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHTTPRequestRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateHTTPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHTTPRequestRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHTTPRequestRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHTTPRequestRuleCreated creates a CreateHTTPRequestRuleCreated with default headers values
func NewCreateHTTPRequestRuleCreated() *CreateHTTPRequestRuleCreated {
	return &CreateHTTPRequestRuleCreated{}
}

/* CreateHTTPRequestRuleCreated describes a response with status code 201, with default header values.

HTTP Request Rule created
*/
type CreateHTTPRequestRuleCreated struct {
	Payload *models2.HTTPRequestRule
}

func (o *CreateHTTPRequestRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateHTTPRequestRuleCreated) GetPayload() *models2.HTTPRequestRule {
	return o.Payload
}

func (o *CreateHTTPRequestRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleAccepted creates a CreateHTTPRequestRuleAccepted with default headers values
func NewCreateHTTPRequestRuleAccepted() *CreateHTTPRequestRuleAccepted {
	return &CreateHTTPRequestRuleAccepted{}
}

/* CreateHTTPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateHTTPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models2.HTTPRequestRule
}

func (o *CreateHTTPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleAccepted  %+v", 202, o.Payload)
}
func (o *CreateHTTPRequestRuleAccepted) GetPayload() *models2.HTTPRequestRule {
	return o.Payload
}

func (o *CreateHTTPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models2.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleBadRequest creates a CreateHTTPRequestRuleBadRequest with default headers values
func NewCreateHTTPRequestRuleBadRequest() *CreateHTTPRequestRuleBadRequest {
	return &CreateHTTPRequestRuleBadRequest{}
}

/* CreateHTTPRequestRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHTTPRequestRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *CreateHTTPRequestRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHTTPRequestRuleBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleConflict creates a CreateHTTPRequestRuleConflict with default headers values
func NewCreateHTTPRequestRuleConflict() *CreateHTTPRequestRuleConflict {
	return &CreateHTTPRequestRuleConflict{}
}

/* CreateHTTPRequestRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateHTTPRequestRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *CreateHTTPRequestRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleConflict  %+v", 409, o.Payload)
}
func (o *CreateHTTPRequestRuleConflict) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleDefault creates a CreateHTTPRequestRuleDefault with default headers values
func NewCreateHTTPRequestRuleDefault(code int) *CreateHTTPRequestRuleDefault {
	return &CreateHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

/* CreateHTTPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateHTTPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}
func (o *CreateHTTPRequestRuleDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
