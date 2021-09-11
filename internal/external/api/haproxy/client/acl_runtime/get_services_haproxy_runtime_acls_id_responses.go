// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// GetServicesHaproxyRuntimeAclsIDReader is a Reader for the GetServicesHaproxyRuntimeAclsID structure.
type GetServicesHaproxyRuntimeAclsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesHaproxyRuntimeAclsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesHaproxyRuntimeAclsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServicesHaproxyRuntimeAclsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServicesHaproxyRuntimeAclsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesHaproxyRuntimeAclsIDOK creates a GetServicesHaproxyRuntimeAclsIDOK with default headers values
func NewGetServicesHaproxyRuntimeAclsIDOK() *GetServicesHaproxyRuntimeAclsIDOK {
	return &GetServicesHaproxyRuntimeAclsIDOK{}
}

/* GetServicesHaproxyRuntimeAclsIDOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServicesHaproxyRuntimeAclsIDOK struct {
	Payload *models.ACLFile
}

func (o *GetServicesHaproxyRuntimeAclsIDOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acls/{id}][%d] getServicesHaproxyRuntimeAclsIdOK  %+v", 200, o.Payload)
}
func (o *GetServicesHaproxyRuntimeAclsIDOK) GetPayload() *models.ACLFile {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeAclsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACLFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesHaproxyRuntimeAclsIDNotFound creates a GetServicesHaproxyRuntimeAclsIDNotFound with default headers values
func NewGetServicesHaproxyRuntimeAclsIDNotFound() *GetServicesHaproxyRuntimeAclsIDNotFound {
	return &GetServicesHaproxyRuntimeAclsIDNotFound{}
}

/* GetServicesHaproxyRuntimeAclsIDNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServicesHaproxyRuntimeAclsIDNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServicesHaproxyRuntimeAclsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acls/{id}][%d] getServicesHaproxyRuntimeAclsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeAclsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeAclsIDDefault creates a GetServicesHaproxyRuntimeAclsIDDefault with default headers values
func NewGetServicesHaproxyRuntimeAclsIDDefault(code int) *GetServicesHaproxyRuntimeAclsIDDefault {
	return &GetServicesHaproxyRuntimeAclsIDDefault{
		_statusCode: code,
	}
}

/* GetServicesHaproxyRuntimeAclsIDDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesHaproxyRuntimeAclsIDDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesHaproxyRuntimeAclsIDDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acls/{id}][%d] GetServicesHaproxyRuntimeAclsID default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesHaproxyRuntimeAclsIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeAclsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
