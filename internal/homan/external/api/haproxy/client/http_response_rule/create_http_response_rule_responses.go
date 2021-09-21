// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateHTTPResponseRuleReader is a Reader for the CreateHTTPResponseRule structure.
type CreateHTTPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHTTPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHTTPResponseRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateHTTPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHTTPResponseRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHTTPResponseRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateHTTPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHTTPResponseRuleCreated creates a CreateHTTPResponseRuleCreated with default headers values
func NewCreateHTTPResponseRuleCreated() *CreateHTTPResponseRuleCreated {
	return &CreateHTTPResponseRuleCreated{}
}

/* CreateHTTPResponseRuleCreated describes a response with status code 201, with default header values.

HTTP Response Rule created
*/
type CreateHTTPResponseRuleCreated struct {
	Payload *models2.HTTPResponseRule
}

func (o *CreateHTTPResponseRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_response_rules][%d] createHttpResponseRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateHTTPResponseRuleCreated) GetPayload() *models2.HTTPResponseRule {
	return o.Payload
}

func (o *CreateHTTPResponseRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.HTTPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPResponseRuleAccepted creates a CreateHTTPResponseRuleAccepted with default headers values
func NewCreateHTTPResponseRuleAccepted() *CreateHTTPResponseRuleAccepted {
	return &CreateHTTPResponseRuleAccepted{}
}

/* CreateHTTPResponseRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateHTTPResponseRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models2.HTTPResponseRule
}

func (o *CreateHTTPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_response_rules][%d] createHttpResponseRuleAccepted  %+v", 202, o.Payload)
}
func (o *CreateHTTPResponseRuleAccepted) GetPayload() *models2.HTTPResponseRule {
	return o.Payload
}

func (o *CreateHTTPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models2.HTTPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPResponseRuleBadRequest creates a CreateHTTPResponseRuleBadRequest with default headers values
func NewCreateHTTPResponseRuleBadRequest() *CreateHTTPResponseRuleBadRequest {
	return &CreateHTTPResponseRuleBadRequest{}
}

/* CreateHTTPResponseRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHTTPResponseRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *CreateHTTPResponseRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_response_rules][%d] createHttpResponseRuleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHTTPResponseRuleBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPResponseRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPResponseRuleConflict creates a CreateHTTPResponseRuleConflict with default headers values
func NewCreateHTTPResponseRuleConflict() *CreateHTTPResponseRuleConflict {
	return &CreateHTTPResponseRuleConflict{}
}

/* CreateHTTPResponseRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateHTTPResponseRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *CreateHTTPResponseRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_response_rules][%d] createHttpResponseRuleConflict  %+v", 409, o.Payload)
}
func (o *CreateHTTPResponseRuleConflict) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPResponseRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPResponseRuleDefault creates a CreateHTTPResponseRuleDefault with default headers values
func NewCreateHTTPResponseRuleDefault(code int) *CreateHTTPResponseRuleDefault {
	return &CreateHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

/* CreateHTTPResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateHTTPResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the create HTTP response rule default response
func (o *CreateHTTPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateHTTPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_response_rules][%d] createHTTPResponseRule default  %+v", o._statusCode, o.Payload)
}
func (o *CreateHTTPResponseRuleDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *CreateHTTPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
