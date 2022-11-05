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

// UsersGroupsUpdateReader is a Reader for the UsersGroupsUpdate structure.
type UsersGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsUpdateOK creates a UsersGroupsUpdateOK with default headers values
func NewUsersGroupsUpdateOK() *UsersGroupsUpdateOK {
	return &UsersGroupsUpdateOK{}
}

/*
UsersGroupsUpdateOK describes a response with status code 200, with default header values.

UsersGroupsUpdateOK users groups update o k
*/
type UsersGroupsUpdateOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this users groups update o k response has a 2xx status code
func (o *UsersGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups update o k response has a 3xx status code
func (o *UsersGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups update o k response has a 4xx status code
func (o *UsersGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups update o k response has a 5xx status code
func (o *UsersGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups update o k response a status code equal to that given
func (o *UsersGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsUpdateDefault creates a UsersGroupsUpdateDefault with default headers values
func NewUsersGroupsUpdateDefault(code int) *UsersGroupsUpdateDefault {
	return &UsersGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersGroupsUpdateDefault describes a response with status code -1, with default header values.

UsersGroupsUpdateDefault users groups update default
*/
type UsersGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups update default response
func (o *UsersGroupsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users groups update default response has a 2xx status code
func (o *UsersGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users groups update default response has a 3xx status code
func (o *UsersGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users groups update default response has a 4xx status code
func (o *UsersGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users groups update default response has a 5xx status code
func (o *UsersGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users groups update default response a status code equal to that given
func (o *UsersGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] users_groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] users_groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}