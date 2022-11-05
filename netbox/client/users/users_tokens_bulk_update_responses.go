// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// UsersTokensBulkUpdateReader is a Reader for the UsersTokensBulkUpdate structure.
type UsersTokensBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensBulkUpdateOK creates a UsersTokensBulkUpdateOK with default headers values
func NewUsersTokensBulkUpdateOK() *UsersTokensBulkUpdateOK {
	return &UsersTokensBulkUpdateOK{}
}

/*
UsersTokensBulkUpdateOK describes a response with status code 200, with default header values.

UsersTokensBulkUpdateOK users tokens bulk update o k
*/
type UsersTokensBulkUpdateOK struct {
	Payload []*models.Token
}

// IsSuccess returns true when this users tokens bulk update o k response has a 2xx status code
func (o *UsersTokensBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens bulk update o k response has a 3xx status code
func (o *UsersTokensBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens bulk update o k response has a 4xx status code
func (o *UsersTokensBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens bulk update o k response has a 5xx status code
func (o *UsersTokensBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens bulk update o k response a status code equal to that given
func (o *UsersTokensBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersTokensBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] usersTokensBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] usersTokensBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensBulkUpdateOK) GetPayload() []*models.Token {
	return o.Payload
}

func (o *UsersTokensBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensBulkUpdateDefault creates a UsersTokensBulkUpdateDefault with default headers values
func NewUsersTokensBulkUpdateDefault(code int) *UsersTokensBulkUpdateDefault {
	return &UsersTokensBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersTokensBulkUpdateDefault describes a response with status code -1, with default header values.

UsersTokensBulkUpdateDefault users tokens bulk update default
*/
type UsersTokensBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens bulk update default response
func (o *UsersTokensBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users tokens bulk update default response has a 2xx status code
func (o *UsersTokensBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens bulk update default response has a 3xx status code
func (o *UsersTokensBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens bulk update default response has a 4xx status code
func (o *UsersTokensBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens bulk update default response has a 5xx status code
func (o *UsersTokensBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens bulk update default response a status code equal to that given
func (o *UsersTokensBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersTokensBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] users_tokens_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] users_tokens_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
