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

// DcimPowerPortsTraceReader is a Reader for the DcimPowerPortsTrace structure.
type DcimPowerPortsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsTraceOK creates a DcimPowerPortsTraceOK with default headers values
func NewDcimPowerPortsTraceOK() *DcimPowerPortsTraceOK {
	return &DcimPowerPortsTraceOK{}
}

/*
DcimPowerPortsTraceOK describes a response with status code 200, with default header values.

DcimPowerPortsTraceOK dcim power ports trace o k
*/
type DcimPowerPortsTraceOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports trace o k response has a 2xx status code
func (o *DcimPowerPortsTraceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports trace o k response has a 3xx status code
func (o *DcimPowerPortsTraceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports trace o k response has a 4xx status code
func (o *DcimPowerPortsTraceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports trace o k response has a 5xx status code
func (o *DcimPowerPortsTraceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports trace o k response a status code equal to that given
func (o *DcimPowerPortsTraceOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPortsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcimPowerPortsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsTraceOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcimPowerPortsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsTraceOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsTraceDefault creates a DcimPowerPortsTraceDefault with default headers values
func NewDcimPowerPortsTraceDefault(code int) *DcimPowerPortsTraceDefault {
	return &DcimPowerPortsTraceDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsTraceDefault describes a response with status code -1, with default header values.

DcimPowerPortsTraceDefault dcim power ports trace default
*/
type DcimPowerPortsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power ports trace default response
func (o *DcimPowerPortsTraceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power ports trace default response has a 2xx status code
func (o *DcimPowerPortsTraceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports trace default response has a 3xx status code
func (o *DcimPowerPortsTraceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports trace default response has a 4xx status code
func (o *DcimPowerPortsTraceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports trace default response has a 5xx status code
func (o *DcimPowerPortsTraceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports trace default response a status code equal to that given
func (o *DcimPowerPortsTraceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPortsTraceDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcim_power-ports_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsTraceDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcim_power-ports_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
