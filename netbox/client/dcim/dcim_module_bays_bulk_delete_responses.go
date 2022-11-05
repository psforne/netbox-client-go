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

// DcimModuleBaysBulkDeleteReader is a Reader for the DcimModuleBaysBulkDelete structure.
type DcimModuleBaysBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleBaysBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysBulkDeleteNoContent creates a DcimModuleBaysBulkDeleteNoContent with default headers values
func NewDcimModuleBaysBulkDeleteNoContent() *DcimModuleBaysBulkDeleteNoContent {
	return &DcimModuleBaysBulkDeleteNoContent{}
}

/*
DcimModuleBaysBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleBaysBulkDeleteNoContent dcim module bays bulk delete no content
*/
type DcimModuleBaysBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module bays bulk delete no content response has a 2xx status code
func (o *DcimModuleBaysBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays bulk delete no content response has a 3xx status code
func (o *DcimModuleBaysBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays bulk delete no content response has a 4xx status code
func (o *DcimModuleBaysBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays bulk delete no content response has a 5xx status code
func (o *DcimModuleBaysBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays bulk delete no content response a status code equal to that given
func (o *DcimModuleBaysBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModuleBaysBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bays/][%d] dcimModuleBaysBulkDeleteNoContent ", 204)
}

func (o *DcimModuleBaysBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bays/][%d] dcimModuleBaysBulkDeleteNoContent ", 204)
}

func (o *DcimModuleBaysBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleBaysBulkDeleteDefault creates a DcimModuleBaysBulkDeleteDefault with default headers values
func NewDcimModuleBaysBulkDeleteDefault(code int) *DcimModuleBaysBulkDeleteDefault {
	return &DcimModuleBaysBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysBulkDeleteDefault describes a response with status code -1, with default header values.

DcimModuleBaysBulkDeleteDefault dcim module bays bulk delete default
*/
type DcimModuleBaysBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays bulk delete default response
func (o *DcimModuleBaysBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays bulk delete default response has a 2xx status code
func (o *DcimModuleBaysBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays bulk delete default response has a 3xx status code
func (o *DcimModuleBaysBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays bulk delete default response has a 4xx status code
func (o *DcimModuleBaysBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays bulk delete default response has a 5xx status code
func (o *DcimModuleBaysBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays bulk delete default response a status code equal to that given
func (o *DcimModuleBaysBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bays/][%d] dcim_module-bays_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bays/][%d] dcim_module-bays_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
