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

// UsersGroupsCreateReader is a Reader for the UsersGroupsCreate structure.
type UsersGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsCreateCreated creates a UsersGroupsCreateCreated with default headers values
func NewUsersGroupsCreateCreated() *UsersGroupsCreateCreated {
	return &UsersGroupsCreateCreated{}
}

/*
UsersGroupsCreateCreated describes a response with status code 201, with default header values.

UsersGroupsCreateCreated users groups create created
*/
type UsersGroupsCreateCreated struct {
	Payload []*models.Group
}

// IsSuccess returns true when this users groups create created response has a 2xx status code
func (o *UsersGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups create created response has a 3xx status code
func (o *UsersGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups create created response has a 4xx status code
func (o *UsersGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups create created response has a 5xx status code
func (o *UsersGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups create created response a status code equal to that given
func (o *UsersGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersGroupsCreateCreated) GetPayload() []*models.Group {
	return o.Payload
}

func (o *UsersGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsCreateDefault creates a UsersGroupsCreateDefault with default headers values
func NewUsersGroupsCreateDefault(code int) *UsersGroupsCreateDefault {
	return &UsersGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
UsersGroupsCreateDefault describes a response with status code -1, with default header values.

UsersGroupsCreateDefault users groups create default
*/
type UsersGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups create default response
func (o *UsersGroupsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users groups create default response has a 2xx status code
func (o *UsersGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users groups create default response has a 3xx status code
func (o *UsersGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users groups create default response has a 4xx status code
func (o *UsersGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users groups create default response has a 5xx status code
func (o *UsersGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users groups create default response a status code equal to that given
func (o *UsersGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /users/groups/][%d] users_groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /users/groups/][%d] users_groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}