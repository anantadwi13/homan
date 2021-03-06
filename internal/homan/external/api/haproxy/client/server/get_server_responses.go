// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetServerReader is a Reader for the GetServer structure.
type GetServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerOK creates a GetServerOK with default headers values
func NewGetServerOK() *GetServerOK {
	return &GetServerOK{}
}

/* GetServerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServerOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetServerOKBody
}

func (o *GetServerOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServerOK  %+v", 200, o.Payload)
}
func (o *GetServerOK) GetPayload() *GetServerOKBody {
	return o.Payload
}

func (o *GetServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetServerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerNotFound creates a GetServerNotFound with default headers values
func NewGetServerNotFound() *GetServerNotFound {
	return &GetServerNotFound{}
}

/* GetServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *GetServerNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServerNotFound  %+v", 404, o.Payload)
}
func (o *GetServerNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServerDefault creates a GetServerDefault with default headers values
func NewGetServerDefault(code int) *GetServerDefault {
	return &GetServerDefault{
		_statusCode: code,
	}
}

/* GetServerDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the get server default response
func (o *GetServerDefault) Code() int {
	return o._statusCode
}

func (o *GetServerDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServer default  %+v", o._statusCode, o.Payload)
}
func (o *GetServerDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *GetServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServerOKBody get server o k body
swagger:model GetServerOKBody
*/
type GetServerOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models2.Server `json:"data,omitempty"`
}

// Validate validates this get server o k body
func (o *GetServerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get server o k body based on the context it is used
func (o *GetServerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
