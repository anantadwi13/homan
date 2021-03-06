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

// GetRuntimeServerReader is a Reader for the GetRuntimeServer structure.
type GetRuntimeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRuntimeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRuntimeServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeServerOK creates a GetRuntimeServerOK with default headers values
func NewGetRuntimeServerOK() *GetRuntimeServerOK {
	return &GetRuntimeServerOK{}
}

/* GetRuntimeServerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetRuntimeServerOK struct {
	Payload *models2.RuntimeServer
}

func (o *GetRuntimeServerOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeServerOK) GetPayload() *models2.RuntimeServer {
	return o.Payload
}

func (o *GetRuntimeServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.RuntimeServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeServerNotFound creates a GetRuntimeServerNotFound with default headers values
func NewGetRuntimeServerNotFound() *GetRuntimeServerNotFound {
	return &GetRuntimeServerNotFound{}
}

/* GetRuntimeServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetRuntimeServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *GetRuntimeServerNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerNotFound  %+v", 404, o.Payload)
}
func (o *GetRuntimeServerNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetRuntimeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuntimeServerDefault creates a GetRuntimeServerDefault with default headers values
func NewGetRuntimeServerDefault(code int) *GetRuntimeServerDefault {
	return &GetRuntimeServerDefault{
		_statusCode: code,
	}
}

/* GetRuntimeServerDefault describes a response with status code -1, with default header values.

General Error
*/
type GetRuntimeServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the get runtime server default response
func (o *GetRuntimeServerDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeServerDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServer default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeServerDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetRuntimeServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
