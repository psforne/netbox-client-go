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

// DcimPowerOutletsReadReader is a Reader for the DcimPowerOutletsRead structure.
type DcimPowerOutletsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsReadOK creates a DcimPowerOutletsReadOK with default headers values
func NewDcimPowerOutletsReadOK() *DcimPowerOutletsReadOK {
	return &DcimPowerOutletsReadOK{}
}

/*
DcimPowerOutletsReadOK describes a response with status code 200, with default header values.

DcimPowerOutletsReadOK dcim power outlets read o k
*/
type DcimPowerOutletsReadOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets read o k response has a 2xx status code
func (o *DcimPowerOutletsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets read o k response has a 3xx status code
func (o *DcimPowerOutletsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets read o k response has a 4xx status code
func (o *DcimPowerOutletsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets read o k response has a 5xx status code
func (o *DcimPowerOutletsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets read o k response a status code equal to that given
func (o *DcimPowerOutletsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerOutletsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcimPowerOutletsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcimPowerOutletsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsReadOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsReadDefault creates a DcimPowerOutletsReadDefault with default headers values
func NewDcimPowerOutletsReadDefault(code int) *DcimPowerOutletsReadDefault {
	return &DcimPowerOutletsReadDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsReadDefault describes a response with status code -1, with default header values.

DcimPowerOutletsReadDefault dcim power outlets read default
*/
type DcimPowerOutletsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlets read default response
func (o *DcimPowerOutletsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlets read default response has a 2xx status code
func (o *DcimPowerOutletsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets read default response has a 3xx status code
func (o *DcimPowerOutletsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets read default response has a 4xx status code
func (o *DcimPowerOutletsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets read default response has a 5xx status code
func (o *DcimPowerOutletsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets read default response a status code equal to that given
func (o *DcimPowerOutletsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcim_power-outlets_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcim_power-outlets_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}