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

// DcimModuleBaysUpdateReader is a Reader for the DcimModuleBaysUpdate structure.
type DcimModuleBaysUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysUpdateOK creates a DcimModuleBaysUpdateOK with default headers values
func NewDcimModuleBaysUpdateOK() *DcimModuleBaysUpdateOK {
	return &DcimModuleBaysUpdateOK{}
}

/*
DcimModuleBaysUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysUpdateOK dcim module bays update o k
*/
type DcimModuleBaysUpdateOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays update o k response has a 2xx status code
func (o *DcimModuleBaysUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays update o k response has a 3xx status code
func (o *DcimModuleBaysUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays update o k response has a 4xx status code
func (o *DcimModuleBaysUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays update o k response has a 5xx status code
func (o *DcimModuleBaysUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays update o k response a status code equal to that given
func (o *DcimModuleBaysUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBaysUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysUpdateDefault creates a DcimModuleBaysUpdateDefault with default headers values
func NewDcimModuleBaysUpdateDefault(code int) *DcimModuleBaysUpdateDefault {
	return &DcimModuleBaysUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBaysUpdateDefault dcim module bays update default
*/
type DcimModuleBaysUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays update default response
func (o *DcimModuleBaysUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays update default response has a 2xx status code
func (o *DcimModuleBaysUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays update default response has a 3xx status code
func (o *DcimModuleBaysUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays update default response has a 4xx status code
func (o *DcimModuleBaysUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays update default response has a 5xx status code
func (o *DcimModuleBaysUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays update default response a status code equal to that given
func (o *DcimModuleBaysUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcim_module-bays_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcim_module-bays_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
