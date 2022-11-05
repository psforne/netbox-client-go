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

// DcimDeviceBayTemplatesBulkUpdateReader is a Reader for the DcimDeviceBayTemplatesBulkUpdate structure.
type DcimDeviceBayTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateOK creates a DcimDeviceBayTemplatesBulkUpdateOK with default headers values
func NewDcimDeviceBayTemplatesBulkUpdateOK() *DcimDeviceBayTemplatesBulkUpdateOK {
	return &DcimDeviceBayTemplatesBulkUpdateOK{}
}

/*
DcimDeviceBayTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesBulkUpdateOK dcim device bay templates bulk update o k
*/
type DcimDeviceBayTemplatesBulkUpdateOK struct {
	Payload []*models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates bulk update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates bulk update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates bulk update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates bulk update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates bulk update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) GetPayload() []*models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesBulkUpdateDefault creates a DcimDeviceBayTemplatesBulkUpdateDefault with default headers values
func NewDcimDeviceBayTemplatesBulkUpdateDefault(code int) *DcimDeviceBayTemplatesBulkUpdateDefault {
	return &DcimDeviceBayTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesBulkUpdateDefault dcim device bay templates bulk update default
*/
type DcimDeviceBayTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates bulk update default response
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device bay templates bulk update default response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates bulk update default response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates bulk update default response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates bulk update default response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates bulk update default response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcim_device-bay-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcim_device-bay-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
