// Code generated by go-swagger; DO NOT EDIT.

package server_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReplaceServerTemplateReader is a Reader for the ReplaceServerTemplate structure.
type ReplaceServerTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceServerTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceServerTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceServerTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceServerTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceServerTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceServerTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceServerTemplateOK creates a ReplaceServerTemplateOK with default headers values
func NewReplaceServerTemplateOK() *ReplaceServerTemplateOK {
	return &ReplaceServerTemplateOK{}
}

/* ReplaceServerTemplateOK describes a response with status code 200, with default header values.

Server template replaced
*/
type ReplaceServerTemplateOK struct {
	Payload *models2.ServerTemplate
}

func (o *ReplaceServerTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_templates/{prefix}][%d] replaceServerTemplateOK  %+v", 200, o.Payload)
}
func (o *ReplaceServerTemplateOK) GetPayload() *models2.ServerTemplate {
	return o.Payload
}

func (o *ReplaceServerTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.ServerTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerTemplateAccepted creates a ReplaceServerTemplateAccepted with default headers values
func NewReplaceServerTemplateAccepted() *ReplaceServerTemplateAccepted {
	return &ReplaceServerTemplateAccepted{}
}

/* ReplaceServerTemplateAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceServerTemplateAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models2.ServerTemplate
}

func (o *ReplaceServerTemplateAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_templates/{prefix}][%d] replaceServerTemplateAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceServerTemplateAccepted) GetPayload() *models2.ServerTemplate {
	return o.Payload
}

func (o *ReplaceServerTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models2.ServerTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerTemplateBadRequest creates a ReplaceServerTemplateBadRequest with default headers values
func NewReplaceServerTemplateBadRequest() *ReplaceServerTemplateBadRequest {
	return &ReplaceServerTemplateBadRequest{}
}

/* ReplaceServerTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceServerTemplateBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceServerTemplateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_templates/{prefix}][%d] replaceServerTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceServerTemplateBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerTemplateNotFound creates a ReplaceServerTemplateNotFound with default headers values
func NewReplaceServerTemplateNotFound() *ReplaceServerTemplateNotFound {
	return &ReplaceServerTemplateNotFound{}
}

/* ReplaceServerTemplateNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceServerTemplateNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceServerTemplateNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_templates/{prefix}][%d] replaceServerTemplateNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceServerTemplateNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerTemplateDefault creates a ReplaceServerTemplateDefault with default headers values
func NewReplaceServerTemplateDefault(code int) *ReplaceServerTemplateDefault {
	return &ReplaceServerTemplateDefault{
		_statusCode: code,
	}
}

/* ReplaceServerTemplateDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceServerTemplateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the replace server template default response
func (o *ReplaceServerTemplateDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceServerTemplateDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_templates/{prefix}][%d] replaceServerTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceServerTemplateDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
