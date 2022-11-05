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

// DcimSitesCreateReader is a Reader for the DcimSitesCreate structure.
type DcimSitesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimSitesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesCreateCreated creates a DcimSitesCreateCreated with default headers values
func NewDcimSitesCreateCreated() *DcimSitesCreateCreated {
	return &DcimSitesCreateCreated{}
}

/*
DcimSitesCreateCreated describes a response with status code 201, with default header values.

DcimSitesCreateCreated dcim sites create created
*/
type DcimSitesCreateCreated struct {
	Payload []*models.Site
}

// IsSuccess returns true when this dcim sites create created response has a 2xx status code
func (o *DcimSitesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites create created response has a 3xx status code
func (o *DcimSitesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites create created response has a 4xx status code
func (o *DcimSitesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites create created response has a 5xx status code
func (o *DcimSitesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites create created response a status code equal to that given
func (o *DcimSitesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimSitesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcimSitesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimSitesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcimSitesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimSitesCreateCreated) GetPayload() []*models.Site {
	return o.Payload
}

func (o *DcimSitesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesCreateDefault creates a DcimSitesCreateDefault with default headers values
func NewDcimSitesCreateDefault(code int) *DcimSitesCreateDefault {
	return &DcimSitesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimSitesCreateDefault describes a response with status code -1, with default header values.

DcimSitesCreateDefault dcim sites create default
*/
type DcimSitesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites create default response
func (o *DcimSitesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim sites create default response has a 2xx status code
func (o *DcimSitesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim sites create default response has a 3xx status code
func (o *DcimSitesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim sites create default response has a 4xx status code
func (o *DcimSitesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim sites create default response has a 5xx status code
func (o *DcimSitesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim sites create default response a status code equal to that given
func (o *DcimSitesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSitesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcim_sites_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcim_sites_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
