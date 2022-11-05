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

// DcimDeviceRolesReadReader is a Reader for the DcimDeviceRolesRead structure.
type DcimDeviceRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesReadOK creates a DcimDeviceRolesReadOK with default headers values
func NewDcimDeviceRolesReadOK() *DcimDeviceRolesReadOK {
	return &DcimDeviceRolesReadOK{}
}

/*
DcimDeviceRolesReadOK describes a response with status code 200, with default header values.

DcimDeviceRolesReadOK dcim device roles read o k
*/
type DcimDeviceRolesReadOK struct {
	Payload *models.DeviceRole
}

// IsSuccess returns true when this dcim device roles read o k response has a 2xx status code
func (o *DcimDeviceRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles read o k response has a 3xx status code
func (o *DcimDeviceRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles read o k response has a 4xx status code
func (o *DcimDeviceRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles read o k response has a 5xx status code
func (o *DcimDeviceRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles read o k response a status code equal to that given
func (o *DcimDeviceRolesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcimDeviceRolesReadOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcimDeviceRolesReadOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesReadOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesReadDefault creates a DcimDeviceRolesReadDefault with default headers values
func NewDcimDeviceRolesReadDefault(code int) *DcimDeviceRolesReadDefault {
	return &DcimDeviceRolesReadDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesReadDefault describes a response with status code -1, with default header values.

DcimDeviceRolesReadDefault dcim device roles read default
*/
type DcimDeviceRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles read default response
func (o *DcimDeviceRolesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device roles read default response has a 2xx status code
func (o *DcimDeviceRolesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles read default response has a 3xx status code
func (o *DcimDeviceRolesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles read default response has a 4xx status code
func (o *DcimDeviceRolesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles read default response has a 5xx status code
func (o *DcimDeviceRolesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles read default response a status code equal to that given
func (o *DcimDeviceRolesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcim_device-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcim_device-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
