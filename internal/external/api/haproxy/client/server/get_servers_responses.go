// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
)

// GetServersReader is a Reader for the GetServers structure.
type GetServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServersOK creates a GetServersOK with default headers values
func NewGetServersOK() *GetServersOK {
	return &GetServersOK{}
}

/* GetServersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServersOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetServersOKBody
}

func (o *GetServersOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers][%d] getServersOK  %+v", 200, o.Payload)
}
func (o *GetServersOK) GetPayload() *GetServersOKBody {
	return o.Payload
}

func (o *GetServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetServersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServersDefault creates a GetServersDefault with default headers values
func NewGetServersDefault(code int) *GetServersDefault {
	return &GetServersDefault{
		_statusCode: code,
	}
}

/* GetServersDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServersDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get servers default response
func (o *GetServersDefault) Code() int {
	return o._statusCode
}

func (o *GetServersDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers][%d] getServers default  %+v", o._statusCode, o.Payload)
}
func (o *GetServersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServersOKBody get servers o k body
swagger:model GetServersOKBody
*/
type GetServersOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Servers `json:"data"`
}

// Validate validates this get servers o k body
func (o *GetServersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServersOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get servers o k body based on the context it is used
func (o *GetServersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServersOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServersOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServersOKBody) UnmarshalBinary(b []byte) error {
	var res GetServersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
