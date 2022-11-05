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

// IpamServiceTemplatesUpdateReader is a Reader for the IpamServiceTemplatesUpdate structure.
type IpamServiceTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServiceTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServiceTemplatesUpdateOK creates a IpamServiceTemplatesUpdateOK with default headers values
func NewIpamServiceTemplatesUpdateOK() *IpamServiceTemplatesUpdateOK {
	return &IpamServiceTemplatesUpdateOK{}
}

/*
IpamServiceTemplatesUpdateOK describes a response with status code 200, with default header values.

IpamServiceTemplatesUpdateOK ipam service templates update o k
*/
type IpamServiceTemplatesUpdateOK struct {
	Payload *models.ServiceTemplate
}

// IsSuccess returns true when this ipam service templates update o k response has a 2xx status code
func (o *IpamServiceTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates update o k response has a 3xx status code
func (o *IpamServiceTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates update o k response has a 4xx status code
func (o *IpamServiceTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates update o k response has a 5xx status code
func (o *IpamServiceTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates update o k response a status code equal to that given
func (o *IpamServiceTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamServiceTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/service-templates/{id}/][%d] ipamServiceTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/service-templates/{id}/][%d] ipamServiceTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesUpdateOK) GetPayload() *models.ServiceTemplate {
	return o.Payload
}

func (o *IpamServiceTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServiceTemplatesUpdateDefault creates a IpamServiceTemplatesUpdateDefault with default headers values
func NewIpamServiceTemplatesUpdateDefault(code int) *IpamServiceTemplatesUpdateDefault {
	return &IpamServiceTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServiceTemplatesUpdateDefault describes a response with status code -1, with default header values.

IpamServiceTemplatesUpdateDefault ipam service templates update default
*/
type IpamServiceTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam service templates update default response
func (o *IpamServiceTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam service templates update default response has a 2xx status code
func (o *IpamServiceTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam service templates update default response has a 3xx status code
func (o *IpamServiceTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam service templates update default response has a 4xx status code
func (o *IpamServiceTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam service templates update default response has a 5xx status code
func (o *IpamServiceTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam service templates update default response a status code equal to that given
func (o *IpamServiceTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamServiceTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/service-templates/{id}/][%d] ipam_service-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/service-templates/{id}/][%d] ipam_service-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServiceTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
