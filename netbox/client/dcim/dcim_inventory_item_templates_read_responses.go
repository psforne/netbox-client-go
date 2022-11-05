// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// DcimInventoryItemTemplatesReadReader is a Reader for the DcimInventoryItemTemplatesRead structure.
type DcimInventoryItemTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesReadOK creates a DcimInventoryItemTemplatesReadOK with default headers values
func NewDcimInventoryItemTemplatesReadOK() *DcimInventoryItemTemplatesReadOK {
	return &DcimInventoryItemTemplatesReadOK{}
}

/*
DcimInventoryItemTemplatesReadOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesReadOK dcim inventory item templates read o k
*/
type DcimInventoryItemTemplatesReadOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates read o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates read o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates read o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates read o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates read o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesReadOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesReadDefault creates a DcimInventoryItemTemplatesReadDefault with default headers values
func NewDcimInventoryItemTemplatesReadDefault(code int) *DcimInventoryItemTemplatesReadDefault {
	return &DcimInventoryItemTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesReadDefault describes a response with status code -1, with default header values.

DcimInventoryItemTemplatesReadDefault dcim inventory item templates read default
*/
type DcimInventoryItemTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item templates read default response
func (o *DcimInventoryItemTemplatesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory item templates read default response has a 2xx status code
func (o *DcimInventoryItemTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates read default response has a 3xx status code
func (o *DcimInventoryItemTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates read default response has a 4xx status code
func (o *DcimInventoryItemTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates read default response has a 5xx status code
func (o *DcimInventoryItemTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates read default response a status code equal to that given
func (o *DcimInventoryItemTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
