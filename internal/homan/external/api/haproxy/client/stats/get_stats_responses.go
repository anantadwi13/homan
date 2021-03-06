// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetStatsReader is a Reader for the GetStats structure.
type GetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStatsOK creates a GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {
	return &GetStatsOK{}
}

/* GetStatsOK describes a response with status code 200, with default header values.

Success
*/
type GetStatsOK struct {
	Payload models2.NativeStats
}

func (o *GetStatsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/stats/native][%d] getStatsOK  %+v", 200, o.Payload)
}
func (o *GetStatsOK) GetPayload() models2.NativeStats {
	return o.Payload
}

func (o *GetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsInternalServerError creates a GetStatsInternalServerError with default headers values
func NewGetStatsInternalServerError() *GetStatsInternalServerError {
	return &GetStatsInternalServerError{}
}

/* GetStatsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStatsInternalServerError struct {
	Payload models2.NativeStats
}

func (o *GetStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/stats/native][%d] getStatsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetStatsInternalServerError) GetPayload() models2.NativeStats {
	return o.Payload
}

func (o *GetStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsDefault creates a GetStatsDefault with default headers values
func NewGetStatsDefault(code int) *GetStatsDefault {
	return &GetStatsDefault{
		_statusCode: code,
	}
}

/* GetStatsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetStatsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the get stats default response
func (o *GetStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetStatsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/stats/native][%d] getStats default  %+v", o._statusCode, o.Payload)
}
func (o *GetStatsDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
