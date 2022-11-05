// Code generated by go-swagger; DO NOT EDIT.

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// TenancyContactAssignmentsBulkUpdateReader is a Reader for the TenancyContactAssignmentsBulkUpdate structure.
type TenancyContactAssignmentsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsBulkUpdateOK creates a TenancyContactAssignmentsBulkUpdateOK with default headers values
func NewTenancyContactAssignmentsBulkUpdateOK() *TenancyContactAssignmentsBulkUpdateOK {
	return &TenancyContactAssignmentsBulkUpdateOK{}
}

/*
TenancyContactAssignmentsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsBulkUpdateOK tenancy contact assignments bulk update o k
*/
type TenancyContactAssignmentsBulkUpdateOK struct {
	Payload []*models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments bulk update o k response has a 2xx status code
func (o *TenancyContactAssignmentsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments bulk update o k response has a 3xx status code
func (o *TenancyContactAssignmentsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments bulk update o k response has a 4xx status code
func (o *TenancyContactAssignmentsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments bulk update o k response has a 5xx status code
func (o *TenancyContactAssignmentsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments bulk update o k response a status code equal to that given
func (o *TenancyContactAssignmentsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactAssignmentsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/][%d] tenancyContactAssignmentsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactAssignmentsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/][%d] tenancyContactAssignmentsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactAssignmentsBulkUpdateOK) GetPayload() []*models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsBulkUpdateDefault creates a TenancyContactAssignmentsBulkUpdateDefault with default headers values
func NewTenancyContactAssignmentsBulkUpdateDefault(code int) *TenancyContactAssignmentsBulkUpdateDefault {
	return &TenancyContactAssignmentsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsBulkUpdateDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsBulkUpdateDefault tenancy contact assignments bulk update default
*/
type TenancyContactAssignmentsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact assignments bulk update default response
func (o *TenancyContactAssignmentsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact assignments bulk update default response has a 2xx status code
func (o *TenancyContactAssignmentsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments bulk update default response has a 3xx status code
func (o *TenancyContactAssignmentsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments bulk update default response has a 4xx status code
func (o *TenancyContactAssignmentsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments bulk update default response has a 5xx status code
func (o *TenancyContactAssignmentsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments bulk update default response a status code equal to that given
func (o *TenancyContactAssignmentsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactAssignmentsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/][%d] tenancy_contact-assignments_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactAssignmentsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/][%d] tenancy_contact-assignments_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactAssignmentsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
