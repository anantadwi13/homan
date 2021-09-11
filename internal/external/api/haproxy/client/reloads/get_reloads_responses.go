// Code generated by go-swagger; DO NOT EDIT.

package reloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// GetReloadsReader is a Reader for the GetReloads structure.
type GetReloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReloadsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReloadsOK creates a GetReloadsOK with default headers values
func NewGetReloadsOK() *GetReloadsOK {
	return &GetReloadsOK{}
}

/* GetReloadsOK describes a response with status code 200, with default header values.

Success
*/
type GetReloadsOK struct {
	Payload models.Reloads
}

func (o *GetReloadsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/reloads][%d] getReloadsOK  %+v", 200, o.Payload)
}
func (o *GetReloadsOK) GetPayload() models.Reloads {
	return o.Payload
}

func (o *GetReloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReloadsDefault creates a GetReloadsDefault with default headers values
func NewGetReloadsDefault(code int) *GetReloadsDefault {
	return &GetReloadsDefault{
		_statusCode: code,
	}
}

/* GetReloadsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetReloadsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get reloads default response
func (o *GetReloadsDefault) Code() int {
	return o._statusCode
}

func (o *GetReloadsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/reloads][%d] getReloads default  %+v", o._statusCode, o.Payload)
}
func (o *GetReloadsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReloadsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
