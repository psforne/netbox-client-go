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

// DcimPlatformsDeleteReader is a Reader for the DcimPlatformsDelete structure.
type DcimPlatformsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPlatformsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsDeleteNoContent creates a DcimPlatformsDeleteNoContent with default headers values
func NewDcimPlatformsDeleteNoContent() *DcimPlatformsDeleteNoContent {
	return &DcimPlatformsDeleteNoContent{}
}

/*
DcimPlatformsDeleteNoContent describes a response with status code 204, with default header values.

DcimPlatformsDeleteNoContent dcim platforms delete no content
*/
type DcimPlatformsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim platforms delete no content response has a 2xx status code
func (o *DcimPlatformsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms delete no content response has a 3xx status code
func (o *DcimPlatformsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms delete no content response has a 4xx status code
func (o *DcimPlatformsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms delete no content response has a 5xx status code
func (o *DcimPlatformsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms delete no content response a status code equal to that given
func (o *DcimPlatformsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPlatformsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcimPlatformsDeleteNoContent ", 204)
}

func (o *DcimPlatformsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcimPlatformsDeleteNoContent ", 204)
}

func (o *DcimPlatformsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPlatformsDeleteDefault creates a DcimPlatformsDeleteDefault with default headers values
func NewDcimPlatformsDeleteDefault(code int) *DcimPlatformsDeleteDefault {
	return &DcimPlatformsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPlatformsDeleteDefault describes a response with status code -1, with default header values.

DcimPlatformsDeleteDefault dcim platforms delete default
*/
type DcimPlatformsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms delete default response
func (o *DcimPlatformsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim platforms delete default response has a 2xx status code
func (o *DcimPlatformsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim platforms delete default response has a 3xx status code
func (o *DcimPlatformsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim platforms delete default response has a 4xx status code
func (o *DcimPlatformsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim platforms delete default response has a 5xx status code
func (o *DcimPlatformsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim platforms delete default response a status code equal to that given
func (o *DcimPlatformsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPlatformsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcim_platforms_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcim_platforms_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
