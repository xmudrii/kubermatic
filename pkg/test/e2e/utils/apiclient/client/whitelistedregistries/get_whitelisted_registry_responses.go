// Code generated by go-swagger; DO NOT EDIT.

package whitelistedregistries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetWhitelistedRegistryReader is a Reader for the GetWhitelistedRegistry structure.
type GetWhitelistedRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWhitelistedRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWhitelistedRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWhitelistedRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWhitelistedRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWhitelistedRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWhitelistedRegistryOK creates a GetWhitelistedRegistryOK with default headers values
func NewGetWhitelistedRegistryOK() *GetWhitelistedRegistryOK {
	return &GetWhitelistedRegistryOK{}
}

/*GetWhitelistedRegistryOK handles this case with default header values.

WhitelistedRegistry
*/
type GetWhitelistedRegistryOK struct {
	Payload *models.WhitelistedRegistry
}

func (o *GetWhitelistedRegistryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/whitelistedregistries/{whitelisted_registry}][%d] getWhitelistedRegistryOK  %+v", 200, o.Payload)
}

func (o *GetWhitelistedRegistryOK) GetPayload() *models.WhitelistedRegistry {
	return o.Payload
}

func (o *GetWhitelistedRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhitelistedRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWhitelistedRegistryUnauthorized creates a GetWhitelistedRegistryUnauthorized with default headers values
func NewGetWhitelistedRegistryUnauthorized() *GetWhitelistedRegistryUnauthorized {
	return &GetWhitelistedRegistryUnauthorized{}
}

/*GetWhitelistedRegistryUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetWhitelistedRegistryUnauthorized struct {
}

func (o *GetWhitelistedRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/whitelistedregistries/{whitelisted_registry}][%d] getWhitelistedRegistryUnauthorized ", 401)
}

func (o *GetWhitelistedRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWhitelistedRegistryForbidden creates a GetWhitelistedRegistryForbidden with default headers values
func NewGetWhitelistedRegistryForbidden() *GetWhitelistedRegistryForbidden {
	return &GetWhitelistedRegistryForbidden{}
}

/*GetWhitelistedRegistryForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetWhitelistedRegistryForbidden struct {
}

func (o *GetWhitelistedRegistryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/whitelistedregistries/{whitelisted_registry}][%d] getWhitelistedRegistryForbidden ", 403)
}

func (o *GetWhitelistedRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWhitelistedRegistryDefault creates a GetWhitelistedRegistryDefault with default headers values
func NewGetWhitelistedRegistryDefault(code int) *GetWhitelistedRegistryDefault {
	return &GetWhitelistedRegistryDefault{
		_statusCode: code,
	}
}

/*GetWhitelistedRegistryDefault handles this case with default header values.

errorResponse
*/
type GetWhitelistedRegistryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get whitelisted registry default response
func (o *GetWhitelistedRegistryDefault) Code() int {
	return o._statusCode
}

func (o *GetWhitelistedRegistryDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/whitelistedregistries/{whitelisted_registry}][%d] getWhitelistedRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *GetWhitelistedRegistryDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWhitelistedRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}