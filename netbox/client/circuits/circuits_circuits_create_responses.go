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

// CircuitsCircuitsCreateReader is a Reader for the CircuitsCircuitsCreate structure.
type CircuitsCircuitsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsCreateCreated creates a CircuitsCircuitsCreateCreated with default headers values
func NewCircuitsCircuitsCreateCreated() *CircuitsCircuitsCreateCreated {
	return &CircuitsCircuitsCreateCreated{}
}

/*
CircuitsCircuitsCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitsCreateCreated circuits circuits create created
*/
type CircuitsCircuitsCreateCreated struct {
	Payload *models.Circuit
}

// IsSuccess returns true when this circuits circuits create created response has a 2xx status code
func (o *CircuitsCircuitsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuits create created response has a 3xx status code
func (o *CircuitsCircuitsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuits create created response has a 4xx status code
func (o *CircuitsCircuitsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuits create created response has a 5xx status code
func (o *CircuitsCircuitsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuits create created response a status code equal to that given
func (o *CircuitsCircuitsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CircuitsCircuitsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuitsCircuitsCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsCircuitsCreateCreated) String() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuitsCircuitsCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsCircuitsCreateCreated) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsCreateDefault creates a CircuitsCircuitsCreateDefault with default headers values
func NewCircuitsCircuitsCreateDefault(code int) *CircuitsCircuitsCreateDefault {
	return &CircuitsCircuitsCreateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitsCreateDefault describes a response with status code -1, with default header values.

CircuitsCircuitsCreateDefault circuits circuits create default
*/
type CircuitsCircuitsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuits create default response
func (o *CircuitsCircuitsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuits create default response has a 2xx status code
func (o *CircuitsCircuitsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuits create default response has a 3xx status code
func (o *CircuitsCircuitsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuits create default response has a 4xx status code
func (o *CircuitsCircuitsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuits create default response has a 5xx status code
func (o *CircuitsCircuitsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuits create default response a status code equal to that given
func (o *CircuitsCircuitsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuits_circuits_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitsCreateDefault) String() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuits_circuits_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}