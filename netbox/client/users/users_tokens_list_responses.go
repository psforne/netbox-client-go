// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// UsersTokensListReader is a Reader for the UsersTokensList structure.
type UsersTokensListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensListOK creates a UsersTokensListOK with default headers values
func NewUsersTokensListOK() *UsersTokensListOK {
	return &UsersTokensListOK{}
}

/*
UsersTokensListOK describes a response with status code 200, with default header values.

UsersTokensListOK users tokens list o k
*/
type UsersTokensListOK struct {
	Payload *UsersTokensListOKBody
}

// IsSuccess returns true when this users tokens list o k response has a 2xx status code
func (o *UsersTokensListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens list o k response has a 3xx status code
func (o *UsersTokensListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens list o k response has a 4xx status code
func (o *UsersTokensListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens list o k response has a 5xx status code
func (o *UsersTokensListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens list o k response a status code equal to that given
func (o *UsersTokensListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersTokensListOK) Error() string {
	return fmt.Sprintf("[GET /users/tokens/][%d] usersTokensListOK  %+v", 200, o.Payload)
}

func (o *UsersTokensListOK) String() string {
	return fmt.Sprintf("[GET /users/tokens/][%d] usersTokensListOK  %+v", 200, o.Payload)
}

func (o *UsersTokensListOK) GetPayload() *UsersTokensListOKBody {
	return o.Payload
}

func (o *UsersTokensListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersTokensListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensListDefault creates a UsersTokensListDefault with default headers values
func NewUsersTokensListDefault(code int) *UsersTokensListDefault {
	return &UsersTokensListDefault{
		_statusCode: code,
	}
}

/*
UsersTokensListDefault describes a response with status code -1, with default header values.

UsersTokensListDefault users tokens list default
*/
type UsersTokensListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens list default response
func (o *UsersTokensListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users tokens list default response has a 2xx status code
func (o *UsersTokensListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens list default response has a 3xx status code
func (o *UsersTokensListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens list default response has a 4xx status code
func (o *UsersTokensListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens list default response has a 5xx status code
func (o *UsersTokensListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens list default response a status code equal to that given
func (o *UsersTokensListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersTokensListDefault) Error() string {
	return fmt.Sprintf("[GET /users/tokens/][%d] users_tokens_list default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensListDefault) String() string {
	return fmt.Sprintf("[GET /users/tokens/][%d] users_tokens_list default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UsersTokensListOKBody users tokens list o k body
swagger:model UsersTokensListOKBody
*/
type UsersTokensListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Token `json:"results"`
}

// Validate validates this users tokens list o k body
func (o *UsersTokensListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersTokensListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("usersTokensListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *UsersTokensListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("usersTokensListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UsersTokensListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("usersTokensListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UsersTokensListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("usersTokensListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersTokensListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersTokensListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this users tokens list o k body based on the context it is used
func (o *UsersTokensListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersTokensListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersTokensListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersTokensListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersTokensListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersTokensListOKBody) UnmarshalBinary(b []byte) error {
	var res UsersTokensListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}