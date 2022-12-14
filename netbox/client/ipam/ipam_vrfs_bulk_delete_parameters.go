// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIpamVrfsBulkDeleteParams creates a new IpamVrfsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsBulkDeleteParams() *IpamVrfsBulkDeleteParams {
	return &IpamVrfsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsBulkDeleteParamsWithTimeout creates a new IpamVrfsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamVrfsBulkDeleteParams {
	return &IpamVrfsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVrfsBulkDeleteParamsWithContext creates a new IpamVrfsBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamVrfsBulkDeleteParamsWithContext(ctx context.Context) *IpamVrfsBulkDeleteParams {
	return &IpamVrfsBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamVrfsBulkDeleteParamsWithHTTPClient creates a new IpamVrfsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamVrfsBulkDeleteParams {
	return &IpamVrfsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
IpamVrfsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the ipam vrfs bulk delete operation.

	Typically these are written to a http.Request.
*/
type IpamVrfsBulkDeleteParams struct {

	// Data.
	Data []*models.DelObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkDeleteParams) WithDefaults() *IpamVrfsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamVrfsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) WithContext(ctx context.Context) *IpamVrfsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamVrfsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) WithData(data []*models.DelObject) *IpamVrfsBulkDeleteParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs bulk delete params
func (o *IpamVrfsBulkDeleteParams) SetData(data []*models.DelObject) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
