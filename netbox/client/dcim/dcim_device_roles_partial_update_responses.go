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

// DcimDeviceRolesPartialUpdateReader is a Reader for the DcimDeviceRolesPartialUpdate structure.
type DcimDeviceRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesPartialUpdateOK creates a DcimDeviceRolesPartialUpdateOK with default headers values
func NewDcimDeviceRolesPartialUpdateOK() *DcimDeviceRolesPartialUpdateOK {
	return &DcimDeviceRolesPartialUpdateOK{}
}

/*
DcimDeviceRolesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesPartialUpdateOK dcim device roles partial update o k
*/
type DcimDeviceRolesPartialUpdateOK struct {
	Payload *models.DeviceRole
}

// IsSuccess returns true when this dcim device roles partial update o k response has a 2xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles partial update o k response has a 3xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles partial update o k response has a 4xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles partial update o k response has a 5xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles partial update o k response a status code equal to that given
func (o *DcimDeviceRolesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesPartialUpdateDefault creates a DcimDeviceRolesPartialUpdateDefault with default headers values
func NewDcimDeviceRolesPartialUpdateDefault(code int) *DcimDeviceRolesPartialUpdateDefault {
	return &DcimDeviceRolesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceRolesPartialUpdateDefault dcim device roles partial update default
*/
type DcimDeviceRolesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles partial update default response
func (o *DcimDeviceRolesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device roles partial update default response has a 2xx status code
func (o *DcimDeviceRolesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles partial update default response has a 3xx status code
func (o *DcimDeviceRolesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles partial update default response has a 4xx status code
func (o *DcimDeviceRolesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles partial update default response has a 5xx status code
func (o *DcimDeviceRolesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles partial update default response a status code equal to that given
func (o *DcimDeviceRolesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceRolesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcim_device-roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcim_device-roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
