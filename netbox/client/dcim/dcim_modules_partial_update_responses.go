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

// DcimModulesPartialUpdateReader is a Reader for the DcimModulesPartialUpdate structure.
type DcimModulesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModulesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModulesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModulesPartialUpdateOK creates a DcimModulesPartialUpdateOK with default headers values
func NewDcimModulesPartialUpdateOK() *DcimModulesPartialUpdateOK {
	return &DcimModulesPartialUpdateOK{}
}

/*
DcimModulesPartialUpdateOK describes a response with status code 200, with default header values.

DcimModulesPartialUpdateOK dcim modules partial update o k
*/
type DcimModulesPartialUpdateOK struct {
	Payload *models.Module
}

// IsSuccess returns true when this dcim modules partial update o k response has a 2xx status code
func (o *DcimModulesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim modules partial update o k response has a 3xx status code
func (o *DcimModulesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules partial update o k response has a 4xx status code
func (o *DcimModulesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim modules partial update o k response has a 5xx status code
func (o *DcimModulesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules partial update o k response a status code equal to that given
func (o *DcimModulesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModulesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/modules/{id}/][%d] dcimModulesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModulesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/modules/{id}/][%d] dcimModulesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModulesPartialUpdateOK) GetPayload() *models.Module {
	return o.Payload
}

func (o *DcimModulesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Module)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModulesPartialUpdateDefault creates a DcimModulesPartialUpdateDefault with default headers values
func NewDcimModulesPartialUpdateDefault(code int) *DcimModulesPartialUpdateDefault {
	return &DcimModulesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModulesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModulesPartialUpdateDefault dcim modules partial update default
*/
type DcimModulesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim modules partial update default response
func (o *DcimModulesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim modules partial update default response has a 2xx status code
func (o *DcimModulesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim modules partial update default response has a 3xx status code
func (o *DcimModulesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim modules partial update default response has a 4xx status code
func (o *DcimModulesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim modules partial update default response has a 5xx status code
func (o *DcimModulesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim modules partial update default response a status code equal to that given
func (o *DcimModulesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModulesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/modules/{id}/][%d] dcim_modules_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/modules/{id}/][%d] dcim_modules_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
