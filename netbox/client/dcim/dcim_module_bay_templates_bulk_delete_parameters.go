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

// NewDcimModuleBayTemplatesBulkDeleteParams creates a new DcimModuleBayTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBayTemplatesBulkDeleteParams() *DcimModuleBayTemplatesBulkDeleteParams {
	return &DcimModuleBayTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBayTemplatesBulkDeleteParamsWithTimeout creates a new DcimModuleBayTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBayTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimModuleBayTemplatesBulkDeleteParams {
	return &DcimModuleBayTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimModuleBayTemplatesBulkDeleteParamsWithContext creates a new DcimModuleBayTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimModuleBayTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimModuleBayTemplatesBulkDeleteParams {
	return &DcimModuleBayTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimModuleBayTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimModuleBayTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBayTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimModuleBayTemplatesBulkDeleteParams {
	return &DcimModuleBayTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimModuleBayTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim module bay templates bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimModuleBayTemplatesBulkDeleteParams struct {

	// Data.
	Data []*models.DelObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module bay templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesBulkDeleteParams) WithDefaults() *DcimModuleBayTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bay templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimModuleBayTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimModuleBayTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimModuleBayTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) WithData(data []*models.DelObject) *DcimModuleBayTemplatesBulkDeleteParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module bay templates bulk delete params
func (o *DcimModuleBayTemplatesBulkDeleteParams) SetData(data []*models.DelObject) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBayTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
