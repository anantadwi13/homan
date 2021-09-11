// Code generated by go-swagger; DO NOT EDIT.

package bind

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

// GetBindsReader is a Reader for the GetBinds structure.
type GetBindsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBindsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBindsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBindsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBindsOK creates a GetBindsOK with default headers values
func NewGetBindsOK() *GetBindsOK {
	return &GetBindsOK{}
}

/* GetBindsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBindsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetBindsOKBody
}

func (o *GetBindsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/binds][%d] getBindsOK  %+v", 200, o.Payload)
}
func (o *GetBindsOK) GetPayload() *GetBindsOKBody {
	return o.Payload
}

func (o *GetBindsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetBindsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBindsDefault creates a GetBindsDefault with default headers values
func NewGetBindsDefault(code int) *GetBindsDefault {
	return &GetBindsDefault{
		_statusCode: code,
	}
}

/* GetBindsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetBindsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get binds default response
func (o *GetBindsDefault) Code() int {
	return o._statusCode
}

func (o *GetBindsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/binds][%d] getBinds default  %+v", o._statusCode, o.Payload)
}
func (o *GetBindsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBindsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBindsOKBody get binds o k body
swagger:model GetBindsOKBody
*/
type GetBindsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Binds `json:"data"`
}

// Validate validates this get binds o k body
func (o *GetBindsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBindsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBindsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBindsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get binds o k body based on the context it is used
func (o *GetBindsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBindsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBindsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBindsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBindsOKBody) UnmarshalBinary(b []byte) error {
	var res GetBindsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
