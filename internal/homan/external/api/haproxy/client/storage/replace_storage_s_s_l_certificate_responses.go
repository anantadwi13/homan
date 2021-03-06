// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReplaceStorageSSLCertificateReader is a Reader for the ReplaceStorageSSLCertificate structure.
type ReplaceStorageSSLCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStorageSSLCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceStorageSSLCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceStorageSSLCertificateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceStorageSSLCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceStorageSSLCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceStorageSSLCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceStorageSSLCertificateOK creates a ReplaceStorageSSLCertificateOK with default headers values
func NewReplaceStorageSSLCertificateOK() *ReplaceStorageSSLCertificateOK {
	return &ReplaceStorageSSLCertificateOK{}
}

/* ReplaceStorageSSLCertificateOK describes a response with status code 200, with default header values.

SSL certificate replaced
*/
type ReplaceStorageSSLCertificateOK struct {
	Payload *models2.SslCertificate
}

func (o *ReplaceStorageSSLCertificateOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/ssl_certificates/{name}][%d] replaceStorageSSLCertificateOK  %+v", 200, o.Payload)
}
func (o *ReplaceStorageSSLCertificateOK) GetPayload() *models2.SslCertificate {
	return o.Payload
}

func (o *ReplaceStorageSSLCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.SslCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageSSLCertificateAccepted creates a ReplaceStorageSSLCertificateAccepted with default headers values
func NewReplaceStorageSSLCertificateAccepted() *ReplaceStorageSSLCertificateAccepted {
	return &ReplaceStorageSSLCertificateAccepted{}
}

/* ReplaceStorageSSLCertificateAccepted describes a response with status code 202, with default header values.

SSL certificate replaced and reload requested
*/
type ReplaceStorageSSLCertificateAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models2.SslCertificate
}

func (o *ReplaceStorageSSLCertificateAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/ssl_certificates/{name}][%d] replaceStorageSSLCertificateAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceStorageSSLCertificateAccepted) GetPayload() *models2.SslCertificate {
	return o.Payload
}

func (o *ReplaceStorageSSLCertificateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models2.SslCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageSSLCertificateBadRequest creates a ReplaceStorageSSLCertificateBadRequest with default headers values
func NewReplaceStorageSSLCertificateBadRequest() *ReplaceStorageSSLCertificateBadRequest {
	return &ReplaceStorageSSLCertificateBadRequest{}
}

/* ReplaceStorageSSLCertificateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceStorageSSLCertificateBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceStorageSSLCertificateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/ssl_certificates/{name}][%d] replaceStorageSSLCertificateBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceStorageSSLCertificateBadRequest) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceStorageSSLCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceStorageSSLCertificateNotFound creates a ReplaceStorageSSLCertificateNotFound with default headers values
func NewReplaceStorageSSLCertificateNotFound() *ReplaceStorageSSLCertificateNotFound {
	return &ReplaceStorageSSLCertificateNotFound{}
}

/* ReplaceStorageSSLCertificateNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceStorageSSLCertificateNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

func (o *ReplaceStorageSSLCertificateNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/ssl_certificates/{name}][%d] replaceStorageSSLCertificateNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceStorageSSLCertificateNotFound) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceStorageSSLCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceStorageSSLCertificateDefault creates a ReplaceStorageSSLCertificateDefault with default headers values
func NewReplaceStorageSSLCertificateDefault(code int) *ReplaceStorageSSLCertificateDefault {
	return &ReplaceStorageSSLCertificateDefault{
		_statusCode: code,
	}
}

/* ReplaceStorageSSLCertificateDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceStorageSSLCertificateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models2.Error
}

// Code gets the status code for the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceStorageSSLCertificateDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/ssl_certificates/{name}][%d] replaceStorageSSLCertificate default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceStorageSSLCertificateDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *ReplaceStorageSSLCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
