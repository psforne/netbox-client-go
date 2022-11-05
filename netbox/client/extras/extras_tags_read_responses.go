// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// ExtrasTagsReadReader is a Reader for the ExtrasTagsRead structure.
type ExtrasTagsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsReadOK creates a ExtrasTagsReadOK with default headers values
func NewExtrasTagsReadOK() *ExtrasTagsReadOK {
	return &ExtrasTagsReadOK{}
}

/*
ExtrasTagsReadOK describes a response with status code 200, with default header values.

ExtrasTagsReadOK extras tags read o k
*/
type ExtrasTagsReadOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this extras tags read o k response has a 2xx status code
func (o *ExtrasTagsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags read o k response has a 3xx status code
func (o *ExtrasTagsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags read o k response has a 4xx status code
func (o *ExtrasTagsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags read o k response has a 5xx status code
func (o *ExtrasTagsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags read o k response a status code equal to that given
func (o *ExtrasTagsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasTagsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/tags/{id}/][%d] extrasTagsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/tags/{id}/][%d] extrasTagsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsReadOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsReadDefault creates a ExtrasTagsReadDefault with default headers values
func NewExtrasTagsReadDefault(code int) *ExtrasTagsReadDefault {
	return &ExtrasTagsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasTagsReadDefault describes a response with status code -1, with default header values.

ExtrasTagsReadDefault extras tags read default
*/
type ExtrasTagsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags read default response
func (o *ExtrasTagsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras tags read default response has a 2xx status code
func (o *ExtrasTagsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras tags read default response has a 3xx status code
func (o *ExtrasTagsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras tags read default response has a 4xx status code
func (o *ExtrasTagsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras tags read default response has a 5xx status code
func (o *ExtrasTagsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras tags read default response a status code equal to that given
func (o *ExtrasTagsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasTagsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/tags/{id}/][%d] extras_tags_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/tags/{id}/][%d] extras_tags_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
