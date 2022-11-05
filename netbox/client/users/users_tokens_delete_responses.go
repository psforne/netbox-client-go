// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersTokensDeleteReader is a Reader for the UsersTokensDelete structure.
type UsersTokensDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersTokensDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensDeleteNoContent creates a UsersTokensDeleteNoContent with default headers values
func NewUsersTokensDeleteNoContent() *UsersTokensDeleteNoContent {
	return &UsersTokensDeleteNoContent{}
}

/*
UsersTokensDeleteNoContent describes a response with status code 204, with default header values.

UsersTokensDeleteNoContent users tokens delete no content
*/
type UsersTokensDeleteNoContent struct {
}

// IsSuccess returns true when this users tokens delete no content response has a 2xx status code
func (o *UsersTokensDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens delete no content response has a 3xx status code
func (o *UsersTokensDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens delete no content response has a 4xx status code
func (o *UsersTokensDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens delete no content response has a 5xx status code
func (o *UsersTokensDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens delete no content response a status code equal to that given
func (o *UsersTokensDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UsersTokensDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/tokens/{id}/][%d] usersTokensDeleteNoContent ", 204)
}

func (o *UsersTokensDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/tokens/{id}/][%d] usersTokensDeleteNoContent ", 204)
}

func (o *UsersTokensDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersTokensDeleteDefault creates a UsersTokensDeleteDefault with default headers values
func NewUsersTokensDeleteDefault(code int) *UsersTokensDeleteDefault {
	return &UsersTokensDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersTokensDeleteDefault describes a response with status code -1, with default header values.

UsersTokensDeleteDefault users tokens delete default
*/
type UsersTokensDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens delete default response
func (o *UsersTokensDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users tokens delete default response has a 2xx status code
func (o *UsersTokensDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens delete default response has a 3xx status code
func (o *UsersTokensDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens delete default response has a 4xx status code
func (o *UsersTokensDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens delete default response has a 5xx status code
func (o *UsersTokensDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens delete default response a status code equal to that given
func (o *UsersTokensDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersTokensDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/tokens/{id}/][%d] users_tokens_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /users/tokens/{id}/][%d] users_tokens_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}