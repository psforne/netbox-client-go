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

// NewDcimDeviceTypesBulkDeleteParams creates a new DcimDeviceTypesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesBulkDeleteParams() *DcimDeviceTypesBulkDeleteParams {
	return &DcimDeviceTypesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesBulkDeleteParamsWithTimeout creates a new DcimDeviceTypesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesBulkDeleteParams {
	return &DcimDeviceTypesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesBulkDeleteParamsWithContext creates a new DcimDeviceTypesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesBulkDeleteParamsWithContext(ctx context.Context) *DcimDeviceTypesBulkDeleteParams {
	return &DcimDeviceTypesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesBulkDeleteParamsWithHTTPClient creates a new DcimDeviceTypesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesBulkDeleteParams {
	return &DcimDeviceTypesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimDeviceTypesBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim device types bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimDeviceTypesBulkDeleteParams struct {

	// Data.
	Data []*models.DelObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesBulkDeleteParams) WithDefaults() *DcimDeviceTypesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) WithContext(ctx context.Context) *DcimDeviceTypesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) WithData(data []*models.DelObject) *DcimDeviceTypesBulkDeleteParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device types bulk delete params
func (o *DcimDeviceTypesBulkDeleteParams) SetData(data []*models.DelObject) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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