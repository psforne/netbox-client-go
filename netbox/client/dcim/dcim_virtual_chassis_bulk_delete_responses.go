// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimVirtualChassisBulkDeleteReader is a Reader for the DcimVirtualChassisBulkDelete structure.
type DcimVirtualChassisBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimVirtualChassisBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisBulkDeleteNoContent creates a DcimVirtualChassisBulkDeleteNoContent with default headers values
func NewDcimVirtualChassisBulkDeleteNoContent() *DcimVirtualChassisBulkDeleteNoContent {
	return &DcimVirtualChassisBulkDeleteNoContent{}
}

/*
DcimVirtualChassisBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimVirtualChassisBulkDeleteNoContent dcim virtual chassis bulk delete no content
*/
type DcimVirtualChassisBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim virtual chassis bulk delete no content response has a 2xx status code
func (o *DcimVirtualChassisBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis bulk delete no content response has a 3xx status code
func (o *DcimVirtualChassisBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis bulk delete no content response has a 4xx status code
func (o *DcimVirtualChassisBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis bulk delete no content response has a 5xx status code
func (o *DcimVirtualChassisBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis bulk delete no content response a status code equal to that given
func (o *DcimVirtualChassisBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimVirtualChassisBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkDeleteNoContent ", 204)
}

func (o *DcimVirtualChassisBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkDeleteNoContent ", 204)
}

func (o *DcimVirtualChassisBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimVirtualChassisBulkDeleteDefault creates a DcimVirtualChassisBulkDeleteDefault with default headers values
func NewDcimVirtualChassisBulkDeleteDefault(code int) *DcimVirtualChassisBulkDeleteDefault {
	return &DcimVirtualChassisBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualChassisBulkDeleteDefault describes a response with status code -1, with default header values.

DcimVirtualChassisBulkDeleteDefault dcim virtual chassis bulk delete default
*/
type DcimVirtualChassisBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis bulk delete default response
func (o *DcimVirtualChassisBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim virtual chassis bulk delete default response has a 2xx status code
func (o *DcimVirtualChassisBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual chassis bulk delete default response has a 3xx status code
func (o *DcimVirtualChassisBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual chassis bulk delete default response has a 4xx status code
func (o *DcimVirtualChassisBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual chassis bulk delete default response has a 5xx status code
func (o *DcimVirtualChassisBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual chassis bulk delete default response a status code equal to that given
func (o *DcimVirtualChassisBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimVirtualChassisBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcim_virtual-chassis_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcim_virtual-chassis_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}