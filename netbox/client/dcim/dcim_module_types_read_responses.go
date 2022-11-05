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

// DcimModuleTypesReadReader is a Reader for the DcimModuleTypesRead structure.
type DcimModuleTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesReadOK creates a DcimModuleTypesReadOK with default headers values
func NewDcimModuleTypesReadOK() *DcimModuleTypesReadOK {
	return &DcimModuleTypesReadOK{}
}

/*
DcimModuleTypesReadOK describes a response with status code 200, with default header values.

DcimModuleTypesReadOK dcim module types read o k
*/
type DcimModuleTypesReadOK struct {
	Payload *models.ModuleType
}

// IsSuccess returns true when this dcim module types read o k response has a 2xx status code
func (o *DcimModuleTypesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types read o k response has a 3xx status code
func (o *DcimModuleTypesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types read o k response has a 4xx status code
func (o *DcimModuleTypesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types read o k response has a 5xx status code
func (o *DcimModuleTypesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types read o k response a status code equal to that given
func (o *DcimModuleTypesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/module-types/{id}/][%d] dcimModuleTypesReadOK  %+v", 200, o.Payload)
}

func (o *DcimModuleTypesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/module-types/{id}/][%d] dcimModuleTypesReadOK  %+v", 200, o.Payload)
}

func (o *DcimModuleTypesReadOK) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesReadDefault creates a DcimModuleTypesReadDefault with default headers values
func NewDcimModuleTypesReadDefault(code int) *DcimModuleTypesReadDefault {
	return &DcimModuleTypesReadDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesReadDefault describes a response with status code -1, with default header values.

DcimModuleTypesReadDefault dcim module types read default
*/
type DcimModuleTypesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types read default response
func (o *DcimModuleTypesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module types read default response has a 2xx status code
func (o *DcimModuleTypesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types read default response has a 3xx status code
func (o *DcimModuleTypesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types read default response has a 4xx status code
func (o *DcimModuleTypesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types read default response has a 5xx status code
func (o *DcimModuleTypesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types read default response a status code equal to that given
func (o *DcimModuleTypesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleTypesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/module-types/{id}/][%d] dcim_module-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/module-types/{id}/][%d] dcim_module-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
