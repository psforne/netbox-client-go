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

// ExtrasImageAttachmentsUpdateReader is a Reader for the ExtrasImageAttachmentsUpdate structure.
type ExtrasImageAttachmentsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsUpdateOK creates a ExtrasImageAttachmentsUpdateOK with default headers values
func NewExtrasImageAttachmentsUpdateOK() *ExtrasImageAttachmentsUpdateOK {
	return &ExtrasImageAttachmentsUpdateOK{}
}

/*
ExtrasImageAttachmentsUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsUpdateOK extras image attachments update o k
*/
type ExtrasImageAttachmentsUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasImageAttachmentsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extrasImageAttachmentsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extrasImageAttachmentsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsUpdateDefault creates a ExtrasImageAttachmentsUpdateDefault with default headers values
func NewExtrasImageAttachmentsUpdateDefault(code int) *ExtrasImageAttachmentsUpdateDefault {
	return &ExtrasImageAttachmentsUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsUpdateDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsUpdateDefault extras image attachments update default
*/
type ExtrasImageAttachmentsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments update default response
func (o *ExtrasImageAttachmentsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras image attachments update default response has a 2xx status code
func (o *ExtrasImageAttachmentsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments update default response has a 3xx status code
func (o *ExtrasImageAttachmentsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments update default response has a 4xx status code
func (o *ExtrasImageAttachmentsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments update default response has a 5xx status code
func (o *ExtrasImageAttachmentsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments update default response a status code equal to that given
func (o *ExtrasImageAttachmentsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasImageAttachmentsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extras_image-attachments_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/{id}/][%d] extras_image-attachments_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
