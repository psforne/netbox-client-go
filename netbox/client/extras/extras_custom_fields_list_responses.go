// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// ExtrasCustomFieldsListReader is a Reader for the ExtrasCustomFieldsList structure.
type ExtrasCustomFieldsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsListOK creates a ExtrasCustomFieldsListOK with default headers values
func NewExtrasCustomFieldsListOK() *ExtrasCustomFieldsListOK {
	return &ExtrasCustomFieldsListOK{}
}

/*
ExtrasCustomFieldsListOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsListOK extras custom fields list o k
*/
type ExtrasCustomFieldsListOK struct {
	Payload *ExtrasCustomFieldsListOKBody
}

// IsSuccess returns true when this extras custom fields list o k response has a 2xx status code
func (o *ExtrasCustomFieldsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields list o k response has a 3xx status code
func (o *ExtrasCustomFieldsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields list o k response has a 4xx status code
func (o *ExtrasCustomFieldsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields list o k response has a 5xx status code
func (o *ExtrasCustomFieldsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields list o k response a status code equal to that given
func (o *ExtrasCustomFieldsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomFieldsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-fields/][%d] extrasCustomFieldsListOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsListOK) String() string {
	return fmt.Sprintf("[GET /extras/custom-fields/][%d] extrasCustomFieldsListOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsListOK) GetPayload() *ExtrasCustomFieldsListOKBody {
	return o.Payload
}

func (o *ExtrasCustomFieldsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasCustomFieldsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsListDefault creates a ExtrasCustomFieldsListDefault with default headers values
func NewExtrasCustomFieldsListDefault(code int) *ExtrasCustomFieldsListDefault {
	return &ExtrasCustomFieldsListDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsListDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldsListDefault extras custom fields list default
*/
type ExtrasCustomFieldsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom fields list default response
func (o *ExtrasCustomFieldsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom fields list default response has a 2xx status code
func (o *ExtrasCustomFieldsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields list default response has a 3xx status code
func (o *ExtrasCustomFieldsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields list default response has a 4xx status code
func (o *ExtrasCustomFieldsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields list default response has a 5xx status code
func (o *ExtrasCustomFieldsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields list default response a status code equal to that given
func (o *ExtrasCustomFieldsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldsListDefault) Error() string {
	return fmt.Sprintf("[GET /extras/custom-fields/][%d] extras_custom-fields_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsListDefault) String() string {
	return fmt.Sprintf("[GET /extras/custom-fields/][%d] extras_custom-fields_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasCustomFieldsListOKBody extras custom fields list o k body
swagger:model ExtrasCustomFieldsListOKBody
*/
type ExtrasCustomFieldsListOKBody struct {

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
	Results []*models.CustomField `json:"results"`
}

// Validate validates this extras custom fields list o k body
func (o *ExtrasCustomFieldsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExtrasCustomFieldsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasCustomFieldsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomFieldsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasCustomFieldsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomFieldsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasCustomFieldsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomFieldsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasCustomFieldsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasCustomFieldsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasCustomFieldsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras custom fields list o k body based on the context it is used
func (o *ExtrasCustomFieldsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasCustomFieldsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasCustomFieldsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasCustomFieldsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasCustomFieldsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasCustomFieldsListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasCustomFieldsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
