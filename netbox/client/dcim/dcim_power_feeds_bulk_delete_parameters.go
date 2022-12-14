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

// NewDcimPowerFeedsBulkDeleteParams creates a new DcimPowerFeedsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsBulkDeleteParams() *DcimPowerFeedsBulkDeleteParams {
	return &DcimPowerFeedsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsBulkDeleteParamsWithTimeout creates a new DcimPowerFeedsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsBulkDeleteParams {
	return &DcimPowerFeedsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsBulkDeleteParamsWithContext creates a new DcimPowerFeedsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerFeedsBulkDeleteParams {
	return &DcimPowerFeedsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsBulkDeleteParamsWithHTTPClient creates a new DcimPowerFeedsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsBulkDeleteParams {
	return &DcimPowerFeedsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimPowerFeedsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim power feeds bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimPowerFeedsBulkDeleteParams struct {

	// Data.
	Data []*models.DelObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsBulkDeleteParams) WithDefaults() *DcimPowerFeedsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerFeedsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) WithData(data []*models.DelObject) *DcimPowerFeedsBulkDeleteParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power feeds bulk delete params
func (o *DcimPowerFeedsBulkDeleteParams) SetData(data []*models.DelObject) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
