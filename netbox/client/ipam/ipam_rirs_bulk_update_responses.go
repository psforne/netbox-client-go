// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// IpamRirsBulkUpdateReader is a Reader for the IpamRirsBulkUpdate structure.
type IpamRirsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsBulkUpdateOK creates a IpamRirsBulkUpdateOK with default headers values
func NewIpamRirsBulkUpdateOK() *IpamRirsBulkUpdateOK {
	return &IpamRirsBulkUpdateOK{}
}

/*
IpamRirsBulkUpdateOK describes a response with status code 200, with default header values.

IpamRirsBulkUpdateOK ipam rirs bulk update o k
*/
type IpamRirsBulkUpdateOK struct {
	Payload []*models.RIR
}

// IsSuccess returns true when this ipam rirs bulk update o k response has a 2xx status code
func (o *IpamRirsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs bulk update o k response has a 3xx status code
func (o *IpamRirsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs bulk update o k response has a 4xx status code
func (o *IpamRirsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs bulk update o k response has a 5xx status code
func (o *IpamRirsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs bulk update o k response a status code equal to that given
func (o *IpamRirsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRirsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/rirs/][%d] ipamRirsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/rirs/][%d] ipamRirsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsBulkUpdateOK) GetPayload() []*models.RIR {
	return o.Payload
}

func (o *IpamRirsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRirsBulkUpdateDefault creates a IpamRirsBulkUpdateDefault with default headers values
func NewIpamRirsBulkUpdateDefault(code int) *IpamRirsBulkUpdateDefault {
	return &IpamRirsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRirsBulkUpdateDefault describes a response with status code -1, with default header values.

IpamRirsBulkUpdateDefault ipam rirs bulk update default
*/
type IpamRirsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam rirs bulk update default response
func (o *IpamRirsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam rirs bulk update default response has a 2xx status code
func (o *IpamRirsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs bulk update default response has a 3xx status code
func (o *IpamRirsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs bulk update default response has a 4xx status code
func (o *IpamRirsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs bulk update default response has a 5xx status code
func (o *IpamRirsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs bulk update default response a status code equal to that given
func (o *IpamRirsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRirsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/rirs/][%d] ipam_rirs_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/rirs/][%d] ipam_rirs_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
