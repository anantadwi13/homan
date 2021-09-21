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

// GetServicesHaproxyRuntimeACLFileEntriesReader is a Reader for the GetServicesHaproxyRuntimeACLFileEntries structure.
type GetServicesHaproxyRuntimeACLFileEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesHaproxyRuntimeACLFileEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesHaproxyRuntimeACLFileEntriesOK creates a GetServicesHaproxyRuntimeACLFileEntriesOK with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesOK() *GetServicesHaproxyRuntimeACLFileEntriesOK {
	return &GetServicesHaproxyRuntimeACLFileEntriesOK{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServicesHaproxyRuntimeACLFileEntriesOK struct {
	Payload models2.ACLFilesEntries
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesOK  %+v", 200, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) GetPayload() models2.ACLFilesEntries {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest creates a GetServicesHaproxyRuntimeACLFileEntriesBadRequest with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest() *GetServicesHaproxyRuntimeACLFileEntriesBadRequest {
	return &GetServicesHaproxyRuntimeACLFileEntriesBadRequest{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetServicesHaproxyRuntimeACLFileEntriesBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesBadRequest  %+v", 400, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesNotFound creates a GetServicesHaproxyRuntimeACLFileEntriesNotFound with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesNotFound() *GetServicesHaproxyRuntimeACLFileEntriesNotFound {
	return &GetServicesHaproxyRuntimeACLFileEntriesNotFound{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServicesHaproxyRuntimeACLFileEntriesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesNotFound  %+v", 404, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesDefault creates a GetServicesHaproxyRuntimeACLFileEntriesDefault with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesDefault(code int) *GetServicesHaproxyRuntimeACLFileEntriesDefault {
	return &GetServicesHaproxyRuntimeACLFileEntriesDefault{
		_statusCode: code,
	}
}

/* GetServicesHaproxyRuntimeACLFileEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesHaproxyRuntimeACLFileEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the get services haproxy runtime ACL file entries default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] GetServicesHaproxyRuntimeACLFileEntries default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
