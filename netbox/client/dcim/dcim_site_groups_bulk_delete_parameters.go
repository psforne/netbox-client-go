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

// NewDcimSiteGroupsBulkDeleteParams creates a new DcimSiteGroupsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSiteGroupsBulkDeleteParams() *DcimSiteGroupsBulkDeleteParams {
	return &DcimSiteGroupsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSiteGroupsBulkDeleteParamsWithTimeout creates a new DcimSiteGroupsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimSiteGroupsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimSiteGroupsBulkDeleteParams {
	return &DcimSiteGroupsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimSiteGroupsBulkDeleteParamsWithContext creates a new DcimSiteGroupsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimSiteGroupsBulkDeleteParamsWithContext(ctx context.Context) *DcimSiteGroupsBulkDeleteParams {
	return &DcimSiteGroupsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimSiteGroupsBulkDeleteParamsWithHTTPClient creates a new DcimSiteGroupsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSiteGroupsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimSiteGroupsBulkDeleteParams {
	return &DcimSiteGroupsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimSiteGroupsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim site groups bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimSiteGroupsBulkDeleteParams struct {

	// Data.
	Data []*models.DelObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim site groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsBulkDeleteParams) WithDefaults() *DcimSiteGroupsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim site groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimSiteGroupsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) WithContext(ctx context.Context) *DcimSiteGroupsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimSiteGroupsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) WithData(data []*models.DelObject) *DcimSiteGroupsBulkDeleteParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim site groups bulk delete params
func (o *DcimSiteGroupsBulkDeleteParams) SetData(data []*models.DelObject) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSiteGroupsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
