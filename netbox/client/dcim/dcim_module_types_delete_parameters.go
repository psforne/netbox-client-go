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
	"github.com/go-openapi/swag"
)

// NewDcimModuleTypesDeleteParams creates a new DcimModuleTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleTypesDeleteParams() *DcimModuleTypesDeleteParams {
	return &DcimModuleTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleTypesDeleteParamsWithTimeout creates a new DcimModuleTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimModuleTypesDeleteParamsWithTimeout(timeout time.Duration) *DcimModuleTypesDeleteParams {
	return &DcimModuleTypesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimModuleTypesDeleteParamsWithContext creates a new DcimModuleTypesDeleteParams object
// with the ability to set a context for a request.
func NewDcimModuleTypesDeleteParamsWithContext(ctx context.Context) *DcimModuleTypesDeleteParams {
	return &DcimModuleTypesDeleteParams{
		Context: ctx,
	}
}

// NewDcimModuleTypesDeleteParamsWithHTTPClient creates a new DcimModuleTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleTypesDeleteParamsWithHTTPClient(client *http.Client) *DcimModuleTypesDeleteParams {
	return &DcimModuleTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimModuleTypesDeleteParams contains all the parameters to send to the API endpoint

	for the dcim module types delete operation.

	Typically these are written to a http.Request.
*/
type DcimModuleTypesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this module type.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesDeleteParams) WithDefaults() *DcimModuleTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) WithTimeout(timeout time.Duration) *DcimModuleTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) WithContext(ctx context.Context) *DcimModuleTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) WithHTTPClient(client *http.Client) *DcimModuleTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) WithID(id int64) *DcimModuleTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim module types delete params
func (o *DcimModuleTypesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
