// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewWirelessWirelessLinksBulkPartialUpdateParams creates a new WirelessWirelessLinksBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLinksBulkPartialUpdateParams() *WirelessWirelessLinksBulkPartialUpdateParams {
	return &WirelessWirelessLinksBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLinksBulkPartialUpdateParamsWithTimeout creates a new WirelessWirelessLinksBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLinksBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *WirelessWirelessLinksBulkPartialUpdateParams {
	return &WirelessWirelessLinksBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLinksBulkPartialUpdateParamsWithContext creates a new WirelessWirelessLinksBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLinksBulkPartialUpdateParamsWithContext(ctx context.Context) *WirelessWirelessLinksBulkPartialUpdateParams {
	return &WirelessWirelessLinksBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLinksBulkPartialUpdateParamsWithHTTPClient creates a new WirelessWirelessLinksBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLinksBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *WirelessWirelessLinksBulkPartialUpdateParams {
	return &WirelessWirelessLinksBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLinksBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the wireless wireless links bulk partial update operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLinksBulkPartialUpdateParams struct {

	// Data.
	Data []*models.WritableWirelessLink

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless links bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WithDefaults() *WirelessWirelessLinksBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless links bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLinksBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *WirelessWirelessLinksBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WithContext(ctx context.Context) *WirelessWirelessLinksBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *WirelessWirelessLinksBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WithData(data []*models.WritableWirelessLink) *WirelessWirelessLinksBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the wireless wireless links bulk partial update params
func (o *WirelessWirelessLinksBulkPartialUpdateParams) SetData(data []*models.WritableWirelessLink) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLinksBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
