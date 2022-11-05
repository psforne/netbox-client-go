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

// DcimDeviceRolesBulkDeleteReader is a Reader for the DcimDeviceRolesBulkDelete structure.
type DcimDeviceRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesBulkDeleteNoContent creates a DcimDeviceRolesBulkDeleteNoContent with default headers values
func NewDcimDeviceRolesBulkDeleteNoContent() *DcimDeviceRolesBulkDeleteNoContent {
	return &DcimDeviceRolesBulkDeleteNoContent{}
}

/*
DcimDeviceRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceRolesBulkDeleteNoContent dcim device roles bulk delete no content
*/
type DcimDeviceRolesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device roles bulk delete no content response has a 2xx status code
func (o *DcimDeviceRolesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles bulk delete no content response has a 3xx status code
func (o *DcimDeviceRolesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles bulk delete no content response has a 4xx status code
func (o *DcimDeviceRolesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles bulk delete no content response has a 5xx status code
func (o *DcimDeviceRolesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles bulk delete no content response a status code equal to that given
func (o *DcimDeviceRolesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimDeviceRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/][%d] dcimDeviceRolesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/][%d] dcimDeviceRolesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceRolesBulkDeleteDefault creates a DcimDeviceRolesBulkDeleteDefault with default headers values
func NewDcimDeviceRolesBulkDeleteDefault(code int) *DcimDeviceRolesBulkDeleteDefault {
	return &DcimDeviceRolesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimDeviceRolesBulkDeleteDefault dcim device roles bulk delete default
*/
type DcimDeviceRolesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles bulk delete default response
func (o *DcimDeviceRolesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device roles bulk delete default response has a 2xx status code
func (o *DcimDeviceRolesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles bulk delete default response has a 3xx status code
func (o *DcimDeviceRolesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles bulk delete default response has a 4xx status code
func (o *DcimDeviceRolesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles bulk delete default response has a 5xx status code
func (o *DcimDeviceRolesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles bulk delete default response a status code equal to that given
func (o *DcimDeviceRolesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceRolesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/][%d] dcim_device-roles_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/][%d] dcim_device-roles_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
