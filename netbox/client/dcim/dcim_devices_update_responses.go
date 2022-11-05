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

// DcimDevicesUpdateReader is a Reader for the DcimDevicesUpdate structure.
type DcimDevicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesUpdateOK creates a DcimDevicesUpdateOK with default headers values
func NewDcimDevicesUpdateOK() *DcimDevicesUpdateOK {
	return &DcimDevicesUpdateOK{}
}

/*
DcimDevicesUpdateOK describes a response with status code 200, with default header values.

DcimDevicesUpdateOK dcim devices update o k
*/
type DcimDevicesUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

// IsSuccess returns true when this dcim devices update o k response has a 2xx status code
func (o *DcimDevicesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices update o k response has a 3xx status code
func (o *DcimDevicesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices update o k response has a 4xx status code
func (o *DcimDevicesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices update o k response has a 5xx status code
func (o *DcimDevicesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices update o k response a status code equal to that given
func (o *DcimDevicesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDevicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcimDevicesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcimDevicesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesUpdateDefault creates a DcimDevicesUpdateDefault with default headers values
func NewDcimDevicesUpdateDefault(code int) *DcimDevicesUpdateDefault {
	return &DcimDevicesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesUpdateDefault describes a response with status code -1, with default header values.

DcimDevicesUpdateDefault dcim devices update default
*/
type DcimDevicesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices update default response
func (o *DcimDevicesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim devices update default response has a 2xx status code
func (o *DcimDevicesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices update default response has a 3xx status code
func (o *DcimDevicesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices update default response has a 4xx status code
func (o *DcimDevicesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices update default response has a 5xx status code
func (o *DcimDevicesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices update default response a status code equal to that given
func (o *DcimDevicesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDevicesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcim_devices_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcim_devices_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}