// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRuntimeMapEntryReader is a Reader for the GetRuntimeMapEntry structure.
type GetRuntimeMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeMapEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRuntimeMapEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRuntimeMapEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeMapEntryOK creates a GetRuntimeMapEntryOK with default headers values
func NewGetRuntimeMapEntryOK() *GetRuntimeMapEntryOK {
	return &GetRuntimeMapEntryOK{}
}

/* GetRuntimeMapEntryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetRuntimeMapEntryOK struct {
	Payload *models2.MapEntry
}

func (o *GetRuntimeMapEntryOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries/{id}][%d] getRuntimeMapEntryOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeMapEntryOK) GetPayload() *models2.MapEntry {
	return o.Payload
}

func (o *GetRuntimeMapEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.MapEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeMapEntryNotFound creates a GetRuntimeMapEntryNotFound with default headers values
func NewGetRuntimeMapEntryNotFound() *GetRuntimeMapEntryNotFound {
	return &GetRuntimeMapEntryNotFound{}
}

/* GetRuntimeMapEntryNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetRuntimeMapEntryNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *GetRuntimeMapEntryNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries/{id}][%d] getRuntimeMapEntryNotFound  %+v", 404, o.Payload)
}
func (o *GetRuntimeMapEntryNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetRuntimeMapEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuntimeMapEntryDefault creates a GetRuntimeMapEntryDefault with default headers values
func NewGetRuntimeMapEntryDefault(code int) *GetRuntimeMapEntryDefault {
	return &GetRuntimeMapEntryDefault{
		_statusCode: code,
	}
}

/* GetRuntimeMapEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type GetRuntimeMapEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeMapEntryDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries/{id}][%d] getRuntimeMapEntry default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeMapEntryDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetRuntimeMapEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
