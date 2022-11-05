// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/psforne/netbox-client-go/netbox/models"
)

// NewDcimDevicesBulkPartialUpdateParams creates a new DcimDevicesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesBulkPartialUpdateParams() *DcimDevicesBulkPartialUpdateParams {
	return &DcimDevicesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesBulkPartialUpdateParamsWithTimeout creates a new DcimDevicesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDevicesBulkPartialUpdateParams {
	return &DcimDevicesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDevicesBulkPartialUpdateParamsWithContext creates a new DcimDevicesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimDevicesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimDevicesBulkPartialUpdateParams {
	return &DcimDevicesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimDevicesBulkPartialUpdateParamsWithHTTPClient creates a new DcimDevicesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDevicesBulkPartialUpdateParams {
	return &DcimDevicesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimDevicesBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim devices bulk partial update operation.

	Typically these are written to a http.Request.
*/
type DcimDevicesBulkPartialUpdateParams struct {

	// Data.
	Data []*models.WritableDeviceWithConfigContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesBulkPartialUpdateParams) WithDefaults() *DcimDevicesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDevicesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimDevicesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDevicesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) WithData(data []*models.WritableDeviceWithConfigContext) *DcimDevicesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim devices bulk partial update params
func (o *DcimDevicesBulkPartialUpdateParams) SetData(data []*models.WritableDeviceWithConfigContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
