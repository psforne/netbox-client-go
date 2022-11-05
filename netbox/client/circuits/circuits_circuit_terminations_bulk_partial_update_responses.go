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

// CircuitsCircuitTerminationsBulkPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkPartialUpdate structure.
type CircuitsCircuitTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateOK creates a CircuitsCircuitTerminationsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateOK() *CircuitsCircuitTerminationsBulkPartialUpdateOK {
	return &CircuitsCircuitTerminationsBulkPartialUpdateOK{}
}

/*
CircuitsCircuitTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateOK circuits circuit terminations bulk partial update o k
*/
type CircuitsCircuitTerminationsBulkPartialUpdateOK struct {
	Payload []*models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations bulk partial update o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations bulk partial update o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations bulk partial update o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations bulk partial update o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations bulk partial update o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) GetPayload() []*models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateDefault creates a CircuitsCircuitTerminationsBulkPartialUpdateDefault with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateDefault(code int) *CircuitsCircuitTerminationsBulkPartialUpdateDefault {
	return &CircuitsCircuitTerminationsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateDefault circuits circuit terminations bulk partial update default
*/
type CircuitsCircuitTerminationsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations bulk partial update default response
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuit terminations bulk partial update default response has a 2xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations bulk partial update default response has a 3xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations bulk partial update default response has a 4xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations bulk partial update default response has a 5xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations bulk partial update default response a status code equal to that given
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuits_circuit-terminations_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuits_circuit-terminations_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
