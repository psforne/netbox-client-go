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

// DcimRegionsBulkDeleteReader is a Reader for the DcimRegionsBulkDelete structure.
type DcimRegionsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRegionsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsBulkDeleteNoContent creates a DcimRegionsBulkDeleteNoContent with default headers values
func NewDcimRegionsBulkDeleteNoContent() *DcimRegionsBulkDeleteNoContent {
	return &DcimRegionsBulkDeleteNoContent{}
}

/*
DcimRegionsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRegionsBulkDeleteNoContent dcim regions bulk delete no content
*/
type DcimRegionsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim regions bulk delete no content response has a 2xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions bulk delete no content response has a 3xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk delete no content response has a 4xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions bulk delete no content response has a 5xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk delete no content response a status code equal to that given
func (o *DcimRegionsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRegionsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteNoContent ", 204)
}

func (o *DcimRegionsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteNoContent ", 204)
}

func (o *DcimRegionsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRegionsBulkDeleteDefault creates a DcimRegionsBulkDeleteDefault with default headers values
func NewDcimRegionsBulkDeleteDefault(code int) *DcimRegionsBulkDeleteDefault {
	return &DcimRegionsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimRegionsBulkDeleteDefault dcim regions bulk delete default
*/
type DcimRegionsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions bulk delete default response
func (o *DcimRegionsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim regions bulk delete default response has a 2xx status code
func (o *DcimRegionsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions bulk delete default response has a 3xx status code
func (o *DcimRegionsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions bulk delete default response has a 4xx status code
func (o *DcimRegionsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions bulk delete default response has a 5xx status code
func (o *DcimRegionsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions bulk delete default response a status code equal to that given
func (o *DcimRegionsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRegionsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcim_regions_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcim_regions_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
