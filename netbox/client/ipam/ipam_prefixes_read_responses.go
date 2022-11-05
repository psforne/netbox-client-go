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

// IpamPrefixesReadReader is a Reader for the IpamPrefixesRead structure.
type IpamPrefixesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesReadOK creates a IpamPrefixesReadOK with default headers values
func NewIpamPrefixesReadOK() *IpamPrefixesReadOK {
	return &IpamPrefixesReadOK{}
}

/*
IpamPrefixesReadOK describes a response with status code 200, with default header values.

IpamPrefixesReadOK ipam prefixes read o k
*/
type IpamPrefixesReadOK struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes read o k response has a 2xx status code
func (o *IpamPrefixesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes read o k response has a 3xx status code
func (o *IpamPrefixesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes read o k response has a 4xx status code
func (o *IpamPrefixesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes read o k response has a 5xx status code
func (o *IpamPrefixesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes read o k response a status code equal to that given
func (o *IpamPrefixesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamPrefixesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipamPrefixesReadOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipamPrefixesReadOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesReadOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesReadDefault creates a IpamPrefixesReadDefault with default headers values
func NewIpamPrefixesReadDefault(code int) *IpamPrefixesReadDefault {
	return &IpamPrefixesReadDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesReadDefault describes a response with status code -1, with default header values.

IpamPrefixesReadDefault ipam prefixes read default
*/
type IpamPrefixesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes read default response
func (o *IpamPrefixesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam prefixes read default response has a 2xx status code
func (o *IpamPrefixesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes read default response has a 3xx status code
func (o *IpamPrefixesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes read default response has a 4xx status code
func (o *IpamPrefixesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes read default response has a 5xx status code
func (o *IpamPrefixesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes read default response a status code equal to that given
func (o *IpamPrefixesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamPrefixesReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipam_prefixes_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipam_prefixes_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
