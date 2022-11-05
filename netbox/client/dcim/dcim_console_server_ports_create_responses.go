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

// DcimConsoleServerPortsCreateReader is a Reader for the DcimConsoleServerPortsCreate structure.
type DcimConsoleServerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsoleServerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsCreateCreated creates a DcimConsoleServerPortsCreateCreated with default headers values
func NewDcimConsoleServerPortsCreateCreated() *DcimConsoleServerPortsCreateCreated {
	return &DcimConsoleServerPortsCreateCreated{}
}

/*
DcimConsoleServerPortsCreateCreated describes a response with status code 201, with default header values.

DcimConsoleServerPortsCreateCreated dcim console server ports create created
*/
type DcimConsoleServerPortsCreateCreated struct {
	Payload []*models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports create created response has a 2xx status code
func (o *DcimConsoleServerPortsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports create created response has a 3xx status code
func (o *DcimConsoleServerPortsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports create created response has a 4xx status code
func (o *DcimConsoleServerPortsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports create created response has a 5xx status code
func (o *DcimConsoleServerPortsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports create created response a status code equal to that given
func (o *DcimConsoleServerPortsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimConsoleServerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcimConsoleServerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsoleServerPortsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcimConsoleServerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsoleServerPortsCreateCreated) GetPayload() []*models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsCreateDefault creates a DcimConsoleServerPortsCreateDefault with default headers values
func NewDcimConsoleServerPortsCreateDefault(code int) *DcimConsoleServerPortsCreateDefault {
	return &DcimConsoleServerPortsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsCreateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsCreateDefault dcim console server ports create default
*/
type DcimConsoleServerPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports create default response
func (o *DcimConsoleServerPortsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console server ports create default response has a 2xx status code
func (o *DcimConsoleServerPortsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports create default response has a 3xx status code
func (o *DcimConsoleServerPortsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports create default response has a 4xx status code
func (o *DcimConsoleServerPortsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports create default response has a 5xx status code
func (o *DcimConsoleServerPortsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports create default response a status code equal to that given
func (o *DcimConsoleServerPortsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsoleServerPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcim_console-server-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcim_console-server-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
