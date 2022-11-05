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

// DcimInterfacesPartialUpdateReader is a Reader for the DcimInterfacesPartialUpdate structure.
type DcimInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesPartialUpdateOK creates a DcimInterfacesPartialUpdateOK with default headers values
func NewDcimInterfacesPartialUpdateOK() *DcimInterfacesPartialUpdateOK {
	return &DcimInterfacesPartialUpdateOK{}
}

/*
DcimInterfacesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesPartialUpdateOK dcim interfaces partial update o k
*/
type DcimInterfacesPartialUpdateOK struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces partial update o k response has a 2xx status code
func (o *DcimInterfacesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces partial update o k response has a 3xx status code
func (o *DcimInterfacesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces partial update o k response has a 4xx status code
func (o *DcimInterfacesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces partial update o k response has a 5xx status code
func (o *DcimInterfacesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces partial update o k response a status code equal to that given
func (o *DcimInterfacesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfacesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcimInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcimInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesPartialUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesPartialUpdateDefault creates a DcimInterfacesPartialUpdateDefault with default headers values
func NewDcimInterfacesPartialUpdateDefault(code int) *DcimInterfacesPartialUpdateDefault {
	return &DcimInterfacesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInterfacesPartialUpdateDefault dcim interfaces partial update default
*/
type DcimInterfacesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces partial update default response
func (o *DcimInterfacesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interfaces partial update default response has a 2xx status code
func (o *DcimInterfacesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces partial update default response has a 3xx status code
func (o *DcimInterfacesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces partial update default response has a 4xx status code
func (o *DcimInterfacesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces partial update default response has a 5xx status code
func (o *DcimInterfacesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces partial update default response a status code equal to that given
func (o *DcimInterfacesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfacesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcim_interfaces_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcim_interfaces_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
