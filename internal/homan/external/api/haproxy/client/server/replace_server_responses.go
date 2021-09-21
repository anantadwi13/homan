// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReplaceServerReader is a Reader for the ReplaceServer structure.
type ReplaceServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceServerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceServerOK creates a ReplaceServerOK with default headers values
func NewReplaceServerOK() *ReplaceServerOK {
	return &ReplaceServerOK{}
}

/* ReplaceServerOK describes a response with status code 200, with default header values.

Server replaced
*/
type ReplaceServerOK struct {
	Payload *models2.Server
}

func (o *ReplaceServerOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/servers/{name}][%d] replaceServerOK  %+v", 200, o.Payload)
}
func (o *ReplaceServerOK) GetPayload() *models2.Server {
	return o.Payload
}

func (o *ReplaceServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerAccepted creates a ReplaceServerAccepted with default headers values
func NewReplaceServerAccepted() *ReplaceServerAccepted {
	return &ReplaceServerAccepted{}
}

/* ReplaceServerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceServerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models2.Server
}

func (o *ReplaceServerAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/servers/{name}][%d] replaceServerAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceServerAccepted) GetPayload() *models2.Server {
	return o.Payload
}

func (o *ReplaceServerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models2.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerBadRequest creates a ReplaceServerBadRequest with default headers values
func NewReplaceServerBadRequest() *ReplaceServerBadRequest {
	return &ReplaceServerBadRequest{}
}

/* ReplaceServerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceServerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceServerBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/servers/{name}][%d] replaceServerBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceServerBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerNotFound creates a ReplaceServerNotFound with default headers values
func NewReplaceServerNotFound() *ReplaceServerNotFound {
	return &ReplaceServerNotFound{}
}

/* ReplaceServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceServerNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/servers/{name}][%d] replaceServerNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceServerNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerDefault creates a ReplaceServerDefault with default headers values
func NewReplaceServerDefault(code int) *ReplaceServerDefault {
	return &ReplaceServerDefault{
		_statusCode: code,
	}
}

/* ReplaceServerDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the replace server default response
func (o *ReplaceServerDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceServerDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/servers/{name}][%d] replaceServer default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceServerDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
