// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// VirtualizationClusterGroupsUpdateReader is a Reader for the VirtualizationClusterGroupsUpdate structure.
type VirtualizationClusterGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsUpdateOK creates a VirtualizationClusterGroupsUpdateOK with default headers values
func NewVirtualizationClusterGroupsUpdateOK() *VirtualizationClusterGroupsUpdateOK {
	return &VirtualizationClusterGroupsUpdateOK{}
}

/*
VirtualizationClusterGroupsUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsUpdateOK virtualization cluster groups update o k
*/
type VirtualizationClusterGroupsUpdateOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups update o k response has a 2xx status code
func (o *VirtualizationClusterGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups update o k response has a 3xx status code
func (o *VirtualizationClusterGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups update o k response has a 4xx status code
func (o *VirtualizationClusterGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups update o k response has a 5xx status code
func (o *VirtualizationClusterGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups update o k response a status code equal to that given
func (o *VirtualizationClusterGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClusterGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsUpdateDefault creates a VirtualizationClusterGroupsUpdateDefault with default headers values
func NewVirtualizationClusterGroupsUpdateDefault(code int) *VirtualizationClusterGroupsUpdateDefault {
	return &VirtualizationClusterGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterGroupsUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterGroupsUpdateDefault virtualization cluster groups update default
*/
type VirtualizationClusterGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster groups update default response
func (o *VirtualizationClusterGroupsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization cluster groups update default response has a 2xx status code
func (o *VirtualizationClusterGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster groups update default response has a 3xx status code
func (o *VirtualizationClusterGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster groups update default response has a 4xx status code
func (o *VirtualizationClusterGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster groups update default response has a 5xx status code
func (o *VirtualizationClusterGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster groups update default response a status code equal to that given
func (o *VirtualizationClusterGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClusterGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
