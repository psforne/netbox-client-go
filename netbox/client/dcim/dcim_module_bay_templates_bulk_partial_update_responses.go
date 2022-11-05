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

// DcimModuleBayTemplatesBulkPartialUpdateReader is a Reader for the DcimModuleBayTemplatesBulkPartialUpdate structure.
type DcimModuleBayTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBayTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesBulkPartialUpdateOK creates a DcimModuleBayTemplatesBulkPartialUpdateOK with default headers values
func NewDcimModuleBayTemplatesBulkPartialUpdateOK() *DcimModuleBayTemplatesBulkPartialUpdateOK {
	return &DcimModuleBayTemplatesBulkPartialUpdateOK{}
}

/*
DcimModuleBayTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimModuleBayTemplatesBulkPartialUpdateOK dcim module bay templates bulk partial update o k
*/
type DcimModuleBayTemplatesBulkPartialUpdateOK struct {
	Payload []*models.ModuleBayTemplate
}

// IsSuccess returns true when this dcim module bay templates bulk partial update o k response has a 2xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates bulk partial update o k response has a 3xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates bulk partial update o k response has a 4xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates bulk partial update o k response has a 5xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates bulk partial update o k response a status code equal to that given
func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/][%d] dcimModuleBayTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/][%d] dcimModuleBayTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) GetPayload() []*models.ModuleBayTemplate {
	return o.Payload
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBayTemplatesBulkPartialUpdateDefault creates a DcimModuleBayTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimModuleBayTemplatesBulkPartialUpdateDefault(code int) *DcimModuleBayTemplatesBulkPartialUpdateDefault {
	return &DcimModuleBayTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesBulkPartialUpdateDefault dcim module bay templates bulk partial update default
*/
type DcimModuleBayTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates bulk partial update default response
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bay templates bulk partial update default response has a 2xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates bulk partial update default response has a 3xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates bulk partial update default response has a 4xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates bulk partial update default response has a 5xx status code
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates bulk partial update default response a status code equal to that given
func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/][%d] dcim_module-bay-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/][%d] dcim_module-bay-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
