// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLinksBulkDeleteReader is a Reader for the WirelessWirelessLinksBulkDelete structure.
type WirelessWirelessLinksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLinksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksBulkDeleteNoContent creates a WirelessWirelessLinksBulkDeleteNoContent with default headers values
func NewWirelessWirelessLinksBulkDeleteNoContent() *WirelessWirelessLinksBulkDeleteNoContent {
	return &WirelessWirelessLinksBulkDeleteNoContent{}
}

/*
WirelessWirelessLinksBulkDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLinksBulkDeleteNoContent wireless wireless links bulk delete no content
*/
type WirelessWirelessLinksBulkDeleteNoContent struct {
}

// IsSuccess returns true when this wireless wireless links bulk delete no content response has a 2xx status code
func (o *WirelessWirelessLinksBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links bulk delete no content response has a 3xx status code
func (o *WirelessWirelessLinksBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links bulk delete no content response has a 4xx status code
func (o *WirelessWirelessLinksBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links bulk delete no content response has a 5xx status code
func (o *WirelessWirelessLinksBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links bulk delete no content response a status code equal to that given
func (o *WirelessWirelessLinksBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *WirelessWirelessLinksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/][%d] wirelessWirelessLinksBulkDeleteNoContent ", 204)
}

func (o *WirelessWirelessLinksBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/][%d] wirelessWirelessLinksBulkDeleteNoContent ", 204)
}

func (o *WirelessWirelessLinksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWirelessWirelessLinksBulkDeleteDefault creates a WirelessWirelessLinksBulkDeleteDefault with default headers values
func NewWirelessWirelessLinksBulkDeleteDefault(code int) *WirelessWirelessLinksBulkDeleteDefault {
	return &WirelessWirelessLinksBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksBulkDeleteDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksBulkDeleteDefault wireless wireless links bulk delete default
*/
type WirelessWirelessLinksBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless links bulk delete default response
func (o *WirelessWirelessLinksBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless links bulk delete default response has a 2xx status code
func (o *WirelessWirelessLinksBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links bulk delete default response has a 3xx status code
func (o *WirelessWirelessLinksBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links bulk delete default response has a 4xx status code
func (o *WirelessWirelessLinksBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links bulk delete default response has a 5xx status code
func (o *WirelessWirelessLinksBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links bulk delete default response a status code equal to that given
func (o *WirelessWirelessLinksBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLinksBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/][%d] wireless_wireless-links_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/][%d] wireless_wireless-links_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
