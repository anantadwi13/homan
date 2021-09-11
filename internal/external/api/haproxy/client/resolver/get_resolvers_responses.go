// Code generated by go-swagger; DO NOT EDIT.

package resolver

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

// GetResolversReader is a Reader for the GetResolvers structure.
type GetResolversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResolversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResolversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetResolversDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResolversOK creates a GetResolversOK with default headers values
func NewGetResolversOK() *GetResolversOK {
	return &GetResolversOK{}
}

/* GetResolversOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetResolversOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetResolversOKBody
}

func (o *GetResolversOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/resolvers][%d] getResolversOK  %+v", 200, o.Payload)
}
func (o *GetResolversOK) GetPayload() *GetResolversOKBody {
	return o.Payload
}

func (o *GetResolversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetResolversOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResolversDefault creates a GetResolversDefault with default headers values
func NewGetResolversDefault(code int) *GetResolversDefault {
	return &GetResolversDefault{
		_statusCode: code,
	}
}

/* GetResolversDefault describes a response with status code -1, with default header values.

General Error
*/
type GetResolversDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get resolvers default response
func (o *GetResolversDefault) Code() int {
	return o._statusCode
}

func (o *GetResolversDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/resolvers][%d] getResolvers default  %+v", o._statusCode, o.Payload)
}
func (o *GetResolversDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResolversDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetResolversOKBody get resolvers o k body
swagger:model GetResolversOKBody
*/
type GetResolversOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Resolvers `json:"data"`
}

// Validate validates this get resolvers o k body
func (o *GetResolversOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResolversOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getResolversOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getResolversOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get resolvers o k body based on the context it is used
func (o *GetResolversOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResolversOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getResolversOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResolversOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResolversOKBody) UnmarshalBinary(b []byte) error {
	var res GetResolversOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
