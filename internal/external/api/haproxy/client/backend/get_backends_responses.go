// Code generated by go-swagger; DO NOT EDIT.

package backend

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

// GetBackendsReader is a Reader for the GetBackends structure.
type GetBackendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackendsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackendsOK creates a GetBackendsOK with default headers values
func NewGetBackendsOK() *GetBackendsOK {
	return &GetBackendsOK{}
}

/* GetBackendsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBackendsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetBackendsOKBody
}

func (o *GetBackendsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backends][%d] getBackendsOK  %+v", 200, o.Payload)
}
func (o *GetBackendsOK) GetPayload() *GetBackendsOKBody {
	return o.Payload
}

func (o *GetBackendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetBackendsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackendsDefault creates a GetBackendsDefault with default headers values
func NewGetBackendsDefault(code int) *GetBackendsDefault {
	return &GetBackendsDefault{
		_statusCode: code,
	}
}

/* GetBackendsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetBackendsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get backends default response
func (o *GetBackendsDefault) Code() int {
	return o._statusCode
}

func (o *GetBackendsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backends][%d] getBackends default  %+v", o._statusCode, o.Payload)
}
func (o *GetBackendsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBackendsOKBody get backends o k body
swagger:model GetBackendsOKBody
*/
type GetBackendsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Backends `json:"data"`
}

// Validate validates this get backends o k body
func (o *GetBackendsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBackendsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get backends o k body based on the context it is used
func (o *GetBackendsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
