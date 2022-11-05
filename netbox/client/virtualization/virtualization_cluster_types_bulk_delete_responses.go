// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterTypesBulkDeleteReader is a Reader for the VirtualizationClusterTypesBulkDelete structure.
type VirtualizationClusterTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesBulkDeleteNoContent creates a VirtualizationClusterTypesBulkDeleteNoContent with default headers values
func NewVirtualizationClusterTypesBulkDeleteNoContent() *VirtualizationClusterTypesBulkDeleteNoContent {
	return &VirtualizationClusterTypesBulkDeleteNoContent{}
}

/*
VirtualizationClusterTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterTypesBulkDeleteNoContent virtualization cluster types bulk delete no content
*/
type VirtualizationClusterTypesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization cluster types bulk delete no content response has a 2xx status code
func (o *VirtualizationClusterTypesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster types bulk delete no content response has a 3xx status code
func (o *VirtualizationClusterTypesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster types bulk delete no content response has a 4xx status code
func (o *VirtualizationClusterTypesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster types bulk delete no content response has a 5xx status code
func (o *VirtualizationClusterTypesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster types bulk delete no content response a status code equal to that given
func (o *VirtualizationClusterTypesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationClusterTypesBulkDeleteDefault creates a VirtualizationClusterTypesBulkDeleteDefault with default headers values
func NewVirtualizationClusterTypesBulkDeleteDefault(code int) *VirtualizationClusterTypesBulkDeleteDefault {
	return &VirtualizationClusterTypesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterTypesBulkDeleteDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesBulkDeleteDefault virtualization cluster types bulk delete default
*/
type VirtualizationClusterTypesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types bulk delete default response
func (o *VirtualizationClusterTypesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization cluster types bulk delete default response has a 2xx status code
func (o *VirtualizationClusterTypesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster types bulk delete default response has a 3xx status code
func (o *VirtualizationClusterTypesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster types bulk delete default response has a 4xx status code
func (o *VirtualizationClusterTypesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster types bulk delete default response has a 5xx status code
func (o *VirtualizationClusterTypesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster types bulk delete default response a status code equal to that given
func (o *VirtualizationClusterTypesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClusterTypesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualization_cluster-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualization_cluster-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
