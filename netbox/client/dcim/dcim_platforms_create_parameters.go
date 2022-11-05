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

// NewDcimPlatformsCreateParams creates a new DcimPlatformsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsCreateParams() *DcimPlatformsCreateParams {
	return &DcimPlatformsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsCreateParamsWithTimeout creates a new DcimPlatformsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsCreateParamsWithTimeout(timeout time.Duration) *DcimPlatformsCreateParams {
	return &DcimPlatformsCreateParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsCreateParamsWithContext creates a new DcimPlatformsCreateParams object
// with the ability to set a context for a request.
func NewDcimPlatformsCreateParamsWithContext(ctx context.Context) *DcimPlatformsCreateParams {
	return &DcimPlatformsCreateParams{
		Context: ctx,
	}
}

// NewDcimPlatformsCreateParamsWithHTTPClient creates a new DcimPlatformsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsCreateParamsWithHTTPClient(client *http.Client) *DcimPlatformsCreateParams {
	return &DcimPlatformsCreateParams{
		HTTPClient: client,
	}
}

/*
DcimPlatformsCreateParams contains all the parameters to send to the API endpoint

	for the dcim platforms create operation.

	Typically these are written to a http.Request.
*/
type DcimPlatformsCreateParams struct {

	// Data.
	Data []*models.WritablePlatform

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsCreateParams) WithDefaults() *DcimPlatformsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms create params
func (o *DcimPlatformsCreateParams) WithTimeout(timeout time.Duration) *DcimPlatformsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms create params
func (o *DcimPlatformsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms create params
func (o *DcimPlatformsCreateParams) WithContext(ctx context.Context) *DcimPlatformsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms create params
func (o *DcimPlatformsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms create params
func (o *DcimPlatformsCreateParams) WithHTTPClient(client *http.Client) *DcimPlatformsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms create params
func (o *DcimPlatformsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim platforms create params
func (o *DcimPlatformsCreateParams) WithData(data []*models.WritablePlatform) *DcimPlatformsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim platforms create params
func (o *DcimPlatformsCreateParams) SetData(data []*models.WritablePlatform) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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