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

// DcimSitesBulkDeleteReader is a Reader for the DcimSitesBulkDelete structure.
type DcimSitesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSitesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesBulkDeleteNoContent creates a DcimSitesBulkDeleteNoContent with default headers values
func NewDcimSitesBulkDeleteNoContent() *DcimSitesBulkDeleteNoContent {
	return &DcimSitesBulkDeleteNoContent{}
}

/*
DcimSitesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimSitesBulkDeleteNoContent dcim sites bulk delete no content
*/
type DcimSitesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim sites bulk delete no content response has a 2xx status code
func (o *DcimSitesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites bulk delete no content response has a 3xx status code
func (o *DcimSitesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites bulk delete no content response has a 4xx status code
func (o *DcimSitesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites bulk delete no content response has a 5xx status code
func (o *DcimSitesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites bulk delete no content response a status code equal to that given
func (o *DcimSitesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimSitesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteNoContent ", 204)
}

func (o *DcimSitesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteNoContent ", 204)
}

func (o *DcimSitesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimSitesBulkDeleteDefault creates a DcimSitesBulkDeleteDefault with default headers values
func NewDcimSitesBulkDeleteDefault(code int) *DcimSitesBulkDeleteDefault {
	return &DcimSitesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimSitesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimSitesBulkDeleteDefault dcim sites bulk delete default
*/
type DcimSitesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites bulk delete default response
func (o *DcimSitesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim sites bulk delete default response has a 2xx status code
func (o *DcimSitesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim sites bulk delete default response has a 3xx status code
func (o *DcimSitesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim sites bulk delete default response has a 4xx status code
func (o *DcimSitesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim sites bulk delete default response has a 5xx status code
func (o *DcimSitesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim sites bulk delete default response a status code equal to that given
func (o *DcimSitesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSitesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcim_sites_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcim_sites_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
