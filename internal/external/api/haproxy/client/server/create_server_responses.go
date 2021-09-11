// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// CreateServerReader is a Reader for the CreateServer structure.
type CreateServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateServerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServerCreated creates a CreateServerCreated with default headers values
func NewCreateServerCreated() *CreateServerCreated {
	return &CreateServerCreated{}
}

/* CreateServerCreated describes a response with status code 201, with default header values.

Server created
*/
type CreateServerCreated struct {
	Payload *models.Server
}

func (o *CreateServerCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerCreated  %+v", 201, o.Payload)
}
func (o *CreateServerCreated) GetPayload() *models.Server {
	return o.Payload
}

func (o *CreateServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerAccepted creates a CreateServerAccepted with default headers values
func NewCreateServerAccepted() *CreateServerAccepted {
	return &CreateServerAccepted{}
}

/* CreateServerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateServerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Server
}

func (o *CreateServerAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerAccepted  %+v", 202, o.Payload)
}
func (o *CreateServerAccepted) GetPayload() *models.Server {
	return o.Payload
}

func (o *CreateServerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerBadRequest creates a CreateServerBadRequest with default headers values
func NewCreateServerBadRequest() *CreateServerBadRequest {
	return &CreateServerBadRequest{}
}

/* CreateServerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateServerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerBadRequest  %+v", 400, o.Payload)
}
func (o *CreateServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerConflict creates a CreateServerConflict with default headers values
func NewCreateServerConflict() *CreateServerConflict {
	return &CreateServerConflict{}
}

/* CreateServerConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateServerConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerConflict  %+v", 409, o.Payload)
}
func (o *CreateServerConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerDefault creates a CreateServerDefault with default headers values
func NewCreateServerDefault(code int) *CreateServerDefault {
	return &CreateServerDefault{
		_statusCode: code,
	}
}

/* CreateServerDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create server default response
func (o *CreateServerDefault) Code() int {
	return o._statusCode
}

func (o *CreateServerDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServer default  %+v", o._statusCode, o.Payload)
}
func (o *CreateServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}