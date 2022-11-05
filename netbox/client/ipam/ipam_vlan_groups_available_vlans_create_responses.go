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

// IpamVlanGroupsAvailableVlansCreateReader is a Reader for the IpamVlanGroupsAvailableVlansCreate structure.
type IpamVlanGroupsAvailableVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsAvailableVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlanGroupsAvailableVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsAvailableVlansCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsAvailableVlansCreateCreated creates a IpamVlanGroupsAvailableVlansCreateCreated with default headers values
func NewIpamVlanGroupsAvailableVlansCreateCreated() *IpamVlanGroupsAvailableVlansCreateCreated {
	return &IpamVlanGroupsAvailableVlansCreateCreated{}
}

/*
IpamVlanGroupsAvailableVlansCreateCreated describes a response with status code 201, with default header values.

IpamVlanGroupsAvailableVlansCreateCreated ipam vlan groups available vlans create created
*/
type IpamVlanGroupsAvailableVlansCreateCreated struct {
	Payload []*models.VLAN
}

// IsSuccess returns true when this ipam vlan groups available vlans create created response has a 2xx status code
func (o *IpamVlanGroupsAvailableVlansCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups available vlans create created response has a 3xx status code
func (o *IpamVlanGroupsAvailableVlansCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups available vlans create created response has a 4xx status code
func (o *IpamVlanGroupsAvailableVlansCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups available vlans create created response has a 5xx status code
func (o *IpamVlanGroupsAvailableVlansCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups available vlans create created response a status code equal to that given
func (o *IpamVlanGroupsAvailableVlansCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamVlanGroupsAvailableVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/{id}/available-vlans/][%d] ipamVlanGroupsAvailableVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsAvailableVlansCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/{id}/available-vlans/][%d] ipamVlanGroupsAvailableVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsAvailableVlansCreateCreated) GetPayload() []*models.VLAN {
	return o.Payload
}

func (o *IpamVlanGroupsAvailableVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsAvailableVlansCreateDefault creates a IpamVlanGroupsAvailableVlansCreateDefault with default headers values
func NewIpamVlanGroupsAvailableVlansCreateDefault(code int) *IpamVlanGroupsAvailableVlansCreateDefault {
	return &IpamVlanGroupsAvailableVlansCreateDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsAvailableVlansCreateDefault describes a response with status code -1, with default header values.

IpamVlanGroupsAvailableVlansCreateDefault ipam vlan groups available vlans create default
*/
type IpamVlanGroupsAvailableVlansCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups available vlans create default response
func (o *IpamVlanGroupsAvailableVlansCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups available vlans create default response has a 2xx status code
func (o *IpamVlanGroupsAvailableVlansCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups available vlans create default response has a 3xx status code
func (o *IpamVlanGroupsAvailableVlansCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups available vlans create default response has a 4xx status code
func (o *IpamVlanGroupsAvailableVlansCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups available vlans create default response has a 5xx status code
func (o *IpamVlanGroupsAvailableVlansCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups available vlans create default response a status code equal to that given
func (o *IpamVlanGroupsAvailableVlansCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsAvailableVlansCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/{id}/available-vlans/][%d] ipam_vlan-groups_available-vlans_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsAvailableVlansCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/{id}/available-vlans/][%d] ipam_vlan-groups_available-vlans_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsAvailableVlansCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsAvailableVlansCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
