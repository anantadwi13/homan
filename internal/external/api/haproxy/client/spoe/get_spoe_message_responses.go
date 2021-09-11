// Code generated by go-swagger; DO NOT EDIT.

package spoe

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

// GetSpoeMessageReader is a Reader for the GetSpoeMessage structure.
type GetSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSpoeMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeMessageOK creates a GetSpoeMessageOK with default headers values
func NewGetSpoeMessageOK() *GetSpoeMessageOK {
	return &GetSpoeMessageOK{}
}

/* GetSpoeMessageOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeMessageOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeMessageOKBody
}

func (o *GetSpoeMessageOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_messages/{name}][%d] getSpoeMessageOK  %+v", 200, o.Payload)
}
func (o *GetSpoeMessageOK) GetPayload() *GetSpoeMessageOKBody {
	return o.Payload
}

func (o *GetSpoeMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetSpoeMessageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeMessageNotFound creates a GetSpoeMessageNotFound with default headers values
func NewGetSpoeMessageNotFound() *GetSpoeMessageNotFound {
	return &GetSpoeMessageNotFound{}
}

/* GetSpoeMessageNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetSpoeMessageNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetSpoeMessageNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_messages/{name}][%d] getSpoeMessageNotFound  %+v", 404, o.Payload)
}
func (o *GetSpoeMessageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSpoeMessageDefault creates a GetSpoeMessageDefault with default headers values
func NewGetSpoeMessageDefault(code int) *GetSpoeMessageDefault {
	return &GetSpoeMessageDefault{
		_statusCode: code,
	}
}

/* GetSpoeMessageDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeMessageDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe message default response
func (o *GetSpoeMessageDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_messages/{name}][%d] getSpoeMessage default  %+v", o._statusCode, o.Payload)
}
func (o *GetSpoeMessageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetSpoeMessageOKBody get spoe message o k body
swagger:model GetSpoeMessageOKBody
*/
type GetSpoeMessageOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data *models.SpoeMessage `json:"data"`
}

// Validate validates this get spoe message o k body
func (o *GetSpoeMessageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeMessageOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeMessageOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeMessageOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get spoe message o k body based on the context it is used
func (o *GetSpoeMessageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeMessageOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeMessageOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeMessageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeMessageOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeMessageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}