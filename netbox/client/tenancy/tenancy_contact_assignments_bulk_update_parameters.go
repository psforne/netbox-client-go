// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// NewTenancyContactAssignmentsBulkUpdateParams creates a new TenancyContactAssignmentsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsBulkUpdateParams() *TenancyContactAssignmentsBulkUpdateParams {
	return &TenancyContactAssignmentsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsBulkUpdateParamsWithTimeout creates a new TenancyContactAssignmentsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsBulkUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsBulkUpdateParams {
	return &TenancyContactAssignmentsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsBulkUpdateParamsWithContext creates a new TenancyContactAssignmentsBulkUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsBulkUpdateParamsWithContext(ctx context.Context) *TenancyContactAssignmentsBulkUpdateParams {
	return &TenancyContactAssignmentsBulkUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsBulkUpdateParamsWithHTTPClient creates a new TenancyContactAssignmentsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsBulkUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsBulkUpdateParams {
	return &TenancyContactAssignmentsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactAssignmentsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the tenancy contact assignments bulk update operation.

	Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsBulkUpdateParams struct {

	// Data.
	Data []*models.WritableContactAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsBulkUpdateParams) WithDefaults() *TenancyContactAssignmentsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) WithContext(ctx context.Context) *TenancyContactAssignmentsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) WithData(data []*models.WritableContactAssignment) *TenancyContactAssignmentsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact assignments bulk update params
func (o *TenancyContactAssignmentsBulkUpdateParams) SetData(data []*models.WritableContactAssignment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
