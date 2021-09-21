// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddPayloadRuntimeACLReader is a Reader for the AddPayloadRuntimeACL structure.
type AddPayloadRuntimeACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPayloadRuntimeACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddPayloadRuntimeACLCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPayloadRuntimeACLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddPayloadRuntimeACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPayloadRuntimeACLCreated creates a AddPayloadRuntimeACLCreated with default headers values
func NewAddPayloadRuntimeACLCreated() *AddPayloadRuntimeACLCreated {
	return &AddPayloadRuntimeACLCreated{}
}

/* AddPayloadRuntimeACLCreated describes a response with status code 201, with default header values.

ACL payload added
*/
type AddPayloadRuntimeACLCreated struct {
	Payload models2.ACLFilesEntries
}

func (o *AddPayloadRuntimeACLCreated) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeAclCreated  %+v", 201, o.Payload)
}
func (o *AddPayloadRuntimeACLCreated) GetPayload() models2.ACLFilesEntries {
	return o.Payload
}

func (o *AddPayloadRuntimeACLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPayloadRuntimeACLBadRequest creates a AddPayloadRuntimeACLBadRequest with default headers values
func NewAddPayloadRuntimeACLBadRequest() *AddPayloadRuntimeACLBadRequest {
	return &AddPayloadRuntimeACLBadRequest{}
}

/* AddPayloadRuntimeACLBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddPayloadRuntimeACLBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *AddPayloadRuntimeACLBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeAclBadRequest  %+v", 400, o.Payload)
}
func (o *AddPayloadRuntimeACLBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeACLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddPayloadRuntimeACLDefault creates a AddPayloadRuntimeACLDefault with default headers values
func NewAddPayloadRuntimeACLDefault(code int) *AddPayloadRuntimeACLDefault {
	return &AddPayloadRuntimeACLDefault{
		_statusCode: code,
	}
}

/* AddPayloadRuntimeACLDefault describes a response with status code -1, with default header values.

General Error
*/
type AddPayloadRuntimeACLDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the add payload runtime ACL default response
func (o *AddPayloadRuntimeACLDefault) Code() int {
	return o._statusCode
}

func (o *AddPayloadRuntimeACLDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeACL default  %+v", o._statusCode, o.Payload)
}
func (o *AddPayloadRuntimeACLDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
