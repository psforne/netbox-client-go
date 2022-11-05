// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// DcimFrontPortsCreateReader is a Reader for the DcimFrontPortsCreate structure.
type DcimFrontPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimFrontPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsCreateCreated creates a DcimFrontPortsCreateCreated with default headers values
func NewDcimFrontPortsCreateCreated() *DcimFrontPortsCreateCreated {
	return &DcimFrontPortsCreateCreated{}
}

/*
DcimFrontPortsCreateCreated describes a response with status code 201, with default header values.

DcimFrontPortsCreateCreated dcim front ports create created
*/
type DcimFrontPortsCreateCreated struct {
	Payload []*models.FrontPort
}

// IsSuccess returns true when this dcim front ports create created response has a 2xx status code
func (o *DcimFrontPortsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports create created response has a 3xx status code
func (o *DcimFrontPortsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports create created response has a 4xx status code
func (o *DcimFrontPortsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports create created response has a 5xx status code
func (o *DcimFrontPortsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports create created response a status code equal to that given
func (o *DcimFrontPortsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimFrontPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/front-ports/][%d] dcimFrontPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimFrontPortsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/front-ports/][%d] dcimFrontPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimFrontPortsCreateCreated) GetPayload() []*models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsCreateDefault creates a DcimFrontPortsCreateDefault with default headers values
func NewDcimFrontPortsCreateDefault(code int) *DcimFrontPortsCreateDefault {
	return &DcimFrontPortsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortsCreateDefault describes a response with status code -1, with default header values.

DcimFrontPortsCreateDefault dcim front ports create default
*/
type DcimFrontPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports create default response
func (o *DcimFrontPortsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim front ports create default response has a 2xx status code
func (o *DcimFrontPortsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front ports create default response has a 3xx status code
func (o *DcimFrontPortsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front ports create default response has a 4xx status code
func (o *DcimFrontPortsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front ports create default response has a 5xx status code
func (o *DcimFrontPortsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front ports create default response a status code equal to that given
func (o *DcimFrontPortsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimFrontPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/front-ports/][%d] dcim_front-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/front-ports/][%d] dcim_front-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
