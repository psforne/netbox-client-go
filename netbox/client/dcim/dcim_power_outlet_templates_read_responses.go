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

// DcimPowerOutletTemplatesReadReader is a Reader for the DcimPowerOutletTemplatesRead structure.
type DcimPowerOutletTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesReadOK creates a DcimPowerOutletTemplatesReadOK with default headers values
func NewDcimPowerOutletTemplatesReadOK() *DcimPowerOutletTemplatesReadOK {
	return &DcimPowerOutletTemplatesReadOK{}
}

/*
DcimPowerOutletTemplatesReadOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesReadOK dcim power outlet templates read o k
*/
type DcimPowerOutletTemplatesReadOK struct {
	Payload *models.PowerOutletTemplate
}

// IsSuccess returns true when this dcim power outlet templates read o k response has a 2xx status code
func (o *DcimPowerOutletTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates read o k response has a 3xx status code
func (o *DcimPowerOutletTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates read o k response has a 4xx status code
func (o *DcimPowerOutletTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates read o k response has a 5xx status code
func (o *DcimPowerOutletTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates read o k response a status code equal to that given
func (o *DcimPowerOutletTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerOutletTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesReadOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesReadDefault creates a DcimPowerOutletTemplatesReadDefault with default headers values
func NewDcimPowerOutletTemplatesReadDefault(code int) *DcimPowerOutletTemplatesReadDefault {
	return &DcimPowerOutletTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesReadDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesReadDefault dcim power outlet templates read default
*/
type DcimPowerOutletTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlet templates read default response
func (o *DcimPowerOutletTemplatesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlet templates read default response has a 2xx status code
func (o *DcimPowerOutletTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates read default response has a 3xx status code
func (o *DcimPowerOutletTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates read default response has a 4xx status code
func (o *DcimPowerOutletTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates read default response has a 5xx status code
func (o *DcimPowerOutletTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates read default response a status code equal to that given
func (o *DcimPowerOutletTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
