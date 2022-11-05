// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// CircuitsProvidersBulkUpdateReader is a Reader for the CircuitsProvidersBulkUpdate structure.
type CircuitsProvidersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersBulkUpdateOK creates a CircuitsProvidersBulkUpdateOK with default headers values
func NewCircuitsProvidersBulkUpdateOK() *CircuitsProvidersBulkUpdateOK {
	return &CircuitsProvidersBulkUpdateOK{}
}

/*
CircuitsProvidersBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersBulkUpdateOK circuits providers bulk update o k
*/
type CircuitsProvidersBulkUpdateOK struct {
	Payload []*models.Provider
}

// IsSuccess returns true when this circuits providers bulk update o k response has a 2xx status code
func (o *CircuitsProvidersBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers bulk update o k response has a 3xx status code
func (o *CircuitsProvidersBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers bulk update o k response has a 4xx status code
func (o *CircuitsProvidersBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers bulk update o k response has a 5xx status code
func (o *CircuitsProvidersBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers bulk update o k response a status code equal to that given
func (o *CircuitsProvidersBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProvidersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/][%d] circuitsProvidersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /circuits/providers/][%d] circuitsProvidersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersBulkUpdateOK) GetPayload() []*models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersBulkUpdateDefault creates a CircuitsProvidersBulkUpdateDefault with default headers values
func NewCircuitsProvidersBulkUpdateDefault(code int) *CircuitsProvidersBulkUpdateDefault {
	return &CircuitsProvidersBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersBulkUpdateDefault describes a response with status code -1, with default header values.

CircuitsProvidersBulkUpdateDefault circuits providers bulk update default
*/
type CircuitsProvidersBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers bulk update default response
func (o *CircuitsProvidersBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers bulk update default response has a 2xx status code
func (o *CircuitsProvidersBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers bulk update default response has a 3xx status code
func (o *CircuitsProvidersBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers bulk update default response has a 4xx status code
func (o *CircuitsProvidersBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers bulk update default response has a 5xx status code
func (o *CircuitsProvidersBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers bulk update default response a status code equal to that given
func (o *CircuitsProvidersBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/][%d] circuits_providers_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /circuits/providers/][%d] circuits_providers_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
