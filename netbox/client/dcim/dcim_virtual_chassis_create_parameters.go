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

// NewDcimVirtualChassisCreateParams creates a new DcimVirtualChassisCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisCreateParams() *DcimVirtualChassisCreateParams {
	return &DcimVirtualChassisCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisCreateParamsWithTimeout creates a new DcimVirtualChassisCreateParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisCreateParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisCreateParams {
	return &DcimVirtualChassisCreateParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisCreateParamsWithContext creates a new DcimVirtualChassisCreateParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisCreateParamsWithContext(ctx context.Context) *DcimVirtualChassisCreateParams {
	return &DcimVirtualChassisCreateParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisCreateParamsWithHTTPClient creates a new DcimVirtualChassisCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisCreateParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisCreateParams {
	return &DcimVirtualChassisCreateParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualChassisCreateParams contains all the parameters to send to the API endpoint

	for the dcim virtual chassis create operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualChassisCreateParams struct {

	// Data.
	Data []*models.WritableVirtualChassis

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisCreateParams) WithDefaults() *DcimVirtualChassisCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithContext(ctx context.Context) *DcimVirtualChassisCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithData(data []*models.WritableVirtualChassis) *DcimVirtualChassisCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetData(data []*models.WritableVirtualChassis) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
