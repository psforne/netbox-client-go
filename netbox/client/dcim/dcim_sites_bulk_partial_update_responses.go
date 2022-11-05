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

// DcimSitesBulkPartialUpdateReader is a Reader for the DcimSitesBulkPartialUpdate structure.
type DcimSitesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesBulkPartialUpdateOK creates a DcimSitesBulkPartialUpdateOK with default headers values
func NewDcimSitesBulkPartialUpdateOK() *DcimSitesBulkPartialUpdateOK {
	return &DcimSitesBulkPartialUpdateOK{}
}

/*
DcimSitesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkPartialUpdateOK dcim sites bulk partial update o k
*/
type DcimSitesBulkPartialUpdateOK struct {
	Payload []*models.Site
}

// IsSuccess returns true when this dcim sites bulk partial update o k response has a 2xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites bulk partial update o k response has a 3xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites bulk partial update o k response has a 4xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites bulk partial update o k response has a 5xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites bulk partial update o k response a status code equal to that given
func (o *DcimSitesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimSitesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSitesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSitesBulkPartialUpdateOK) GetPayload() []*models.Site {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesBulkPartialUpdateDefault creates a DcimSitesBulkPartialUpdateDefault with default headers values
func NewDcimSitesBulkPartialUpdateDefault(code int) *DcimSitesBulkPartialUpdateDefault {
	return &DcimSitesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimSitesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimSitesBulkPartialUpdateDefault dcim sites bulk partial update default
*/
type DcimSitesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites bulk partial update default response
func (o *DcimSitesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim sites bulk partial update default response has a 2xx status code
func (o *DcimSitesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim sites bulk partial update default response has a 3xx status code
func (o *DcimSitesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim sites bulk partial update default response has a 4xx status code
func (o *DcimSitesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim sites bulk partial update default response has a 5xx status code
func (o *DcimSitesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim sites bulk partial update default response a status code equal to that given
func (o *DcimSitesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSitesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcim_sites_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcim_sites_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}