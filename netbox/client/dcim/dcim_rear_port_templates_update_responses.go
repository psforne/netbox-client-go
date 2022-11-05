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

// DcimRearPortTemplatesUpdateReader is a Reader for the DcimRearPortTemplatesUpdate structure.
type DcimRearPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesUpdateOK creates a DcimRearPortTemplatesUpdateOK with default headers values
func NewDcimRearPortTemplatesUpdateOK() *DcimRearPortTemplatesUpdateOK {
	return &DcimRearPortTemplatesUpdateOK{}
}

/*
DcimRearPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesUpdateOK dcim rear port templates update o k
*/
type DcimRearPortTemplatesUpdateOK struct {
	Payload *models.RearPortTemplate
}

// IsSuccess returns true when this dcim rear port templates update o k response has a 2xx status code
func (o *DcimRearPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear port templates update o k response has a 3xx status code
func (o *DcimRearPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear port templates update o k response has a 4xx status code
func (o *DcimRearPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear port templates update o k response has a 5xx status code
func (o *DcimRearPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear port templates update o k response a status code equal to that given
func (o *DcimRearPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRearPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesUpdateDefault creates a DcimRearPortTemplatesUpdateDefault with default headers values
func NewDcimRearPortTemplatesUpdateDefault(code int) *DcimRearPortTemplatesUpdateDefault {
	return &DcimRearPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimRearPortTemplatesUpdateDefault dcim rear port templates update default
*/
type DcimRearPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear port templates update default response
func (o *DcimRearPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rear port templates update default response has a 2xx status code
func (o *DcimRearPortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear port templates update default response has a 3xx status code
func (o *DcimRearPortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear port templates update default response has a 4xx status code
func (o *DcimRearPortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear port templates update default response has a 5xx status code
func (o *DcimRearPortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear port templates update default response a status code equal to that given
func (o *DcimRearPortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRearPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}