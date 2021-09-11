// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// GetOneSpoeFileReader is a Reader for the GetOneSpoeFile structure.
type GetOneSpoeFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneSpoeFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneSpoeFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOneSpoeFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOneSpoeFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOneSpoeFileOK creates a GetOneSpoeFileOK with default headers values
func NewGetOneSpoeFileOK() *GetOneSpoeFileOK {
	return &GetOneSpoeFileOK{}
}

/* GetOneSpoeFileOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOneSpoeFileOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetOneSpoeFileOKBody
}

func (o *GetOneSpoeFileOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files/{name}][%d] getOneSpoeFileOK  %+v", 200, o.Payload)
}
func (o *GetOneSpoeFileOK) GetPayload() *GetOneSpoeFileOKBody {
	return o.Payload
}

func (o *GetOneSpoeFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetOneSpoeFileOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneSpoeFileNotFound creates a GetOneSpoeFileNotFound with default headers values
func NewGetOneSpoeFileNotFound() *GetOneSpoeFileNotFound {
	return &GetOneSpoeFileNotFound{}
}

/* GetOneSpoeFileNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetOneSpoeFileNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetOneSpoeFileNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files/{name}][%d] getOneSpoeFileNotFound  %+v", 404, o.Payload)
}
func (o *GetOneSpoeFileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneSpoeFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOneSpoeFileDefault creates a GetOneSpoeFileDefault with default headers values
func NewGetOneSpoeFileDefault(code int) *GetOneSpoeFileDefault {
	return &GetOneSpoeFileDefault{
		_statusCode: code,
	}
}

/* GetOneSpoeFileDefault describes a response with status code -1, with default header values.

General Error
*/
type GetOneSpoeFileDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get one spoe file default response
func (o *GetOneSpoeFileDefault) Code() int {
	return o._statusCode
}

func (o *GetOneSpoeFileDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files/{name}][%d] getOneSpoeFile default  %+v", o._statusCode, o.Payload)
}
func (o *GetOneSpoeFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneSpoeFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetOneSpoeFileOKBody get one spoe file o k body
swagger:model GetOneSpoeFileOKBody
*/
type GetOneSpoeFileOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data string `json:"data,omitempty"`
}

// Validate validates this get one spoe file o k body
func (o *GetOneSpoeFileOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get one spoe file o k body based on context it is used
func (o *GetOneSpoeFileOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOneSpoeFileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOneSpoeFileOKBody) UnmarshalBinary(b []byte) error {
	var res GetOneSpoeFileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
