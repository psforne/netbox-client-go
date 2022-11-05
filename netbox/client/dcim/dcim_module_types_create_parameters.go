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

// NewDcimModuleTypesCreateParams creates a new DcimModuleTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleTypesCreateParams() *DcimModuleTypesCreateParams {
	return &DcimModuleTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleTypesCreateParamsWithTimeout creates a new DcimModuleTypesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleTypesCreateParamsWithTimeout(timeout time.Duration) *DcimModuleTypesCreateParams {
	return &DcimModuleTypesCreateParams{
		timeout: timeout,
	}
}

// NewDcimModuleTypesCreateParamsWithContext creates a new DcimModuleTypesCreateParams object
// with the ability to set a context for a request.
func NewDcimModuleTypesCreateParamsWithContext(ctx context.Context) *DcimModuleTypesCreateParams {
	return &DcimModuleTypesCreateParams{
		Context: ctx,
	}
}

// NewDcimModuleTypesCreateParamsWithHTTPClient creates a new DcimModuleTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleTypesCreateParamsWithHTTPClient(client *http.Client) *DcimModuleTypesCreateParams {
	return &DcimModuleTypesCreateParams{
		HTTPClient: client,
	}
}

/*
DcimModuleTypesCreateParams contains all the parameters to send to the API endpoint

	for the dcim module types create operation.

	Typically these are written to a http.Request.
*/
type DcimModuleTypesCreateParams struct {

	// Data.
	Data []*models.WritableModuleType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesCreateParams) WithDefaults() *DcimModuleTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module types create params
func (o *DcimModuleTypesCreateParams) WithTimeout(timeout time.Duration) *DcimModuleTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module types create params
func (o *DcimModuleTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module types create params
func (o *DcimModuleTypesCreateParams) WithContext(ctx context.Context) *DcimModuleTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module types create params
func (o *DcimModuleTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module types create params
func (o *DcimModuleTypesCreateParams) WithHTTPClient(client *http.Client) *DcimModuleTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module types create params
func (o *DcimModuleTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module types create params
func (o *DcimModuleTypesCreateParams) WithData(data []*models.WritableModuleType) *DcimModuleTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module types create params
func (o *DcimModuleTypesCreateParams) SetData(data []*models.WritableModuleType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
