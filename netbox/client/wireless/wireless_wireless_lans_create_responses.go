// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// WirelessWirelessLansCreateReader is a Reader for the WirelessWirelessLansCreate structure.
type WirelessWirelessLansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWirelessWirelessLansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansCreateCreated creates a WirelessWirelessLansCreateCreated with default headers values
func NewWirelessWirelessLansCreateCreated() *WirelessWirelessLansCreateCreated {
	return &WirelessWirelessLansCreateCreated{}
}

/*
WirelessWirelessLansCreateCreated describes a response with status code 201, with default header values.

WirelessWirelessLansCreateCreated wireless wireless lans create created
*/
type WirelessWirelessLansCreateCreated struct {
	Payload []*models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans create created response has a 2xx status code
func (o *WirelessWirelessLansCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans create created response has a 3xx status code
func (o *WirelessWirelessLansCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans create created response has a 4xx status code
func (o *WirelessWirelessLansCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans create created response has a 5xx status code
func (o *WirelessWirelessLansCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans create created response a status code equal to that given
func (o *WirelessWirelessLansCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WirelessWirelessLansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-lans/][%d] wirelessWirelessLansCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLansCreateCreated) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-lans/][%d] wirelessWirelessLansCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLansCreateCreated) GetPayload() []*models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLansCreateDefault creates a WirelessWirelessLansCreateDefault with default headers values
func NewWirelessWirelessLansCreateDefault(code int) *WirelessWirelessLansCreateDefault {
	return &WirelessWirelessLansCreateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLansCreateDefault describes a response with status code -1, with default header values.

WirelessWirelessLansCreateDefault wireless wireless lans create default
*/
type WirelessWirelessLansCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lans create default response
func (o *WirelessWirelessLansCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless lans create default response has a 2xx status code
func (o *WirelessWirelessLansCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lans create default response has a 3xx status code
func (o *WirelessWirelessLansCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lans create default response has a 4xx status code
func (o *WirelessWirelessLansCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lans create default response has a 5xx status code
func (o *WirelessWirelessLansCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lans create default response a status code equal to that given
func (o *WirelessWirelessLansCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLansCreateDefault) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-lans/][%d] wireless_wireless-lans_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansCreateDefault) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-lans/][%d] wireless_wireless-lans_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}