// Code generated by go-swagger; DO NOT EDIT.

package extras

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

	"github.com/psforne/netbox-client-go/netbox/models"
)

// NewExtrasCustomFieldsPartialUpdateParams creates a new ExtrasCustomFieldsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsPartialUpdateParams() *ExtrasCustomFieldsPartialUpdateParams {
	return &ExtrasCustomFieldsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsPartialUpdateParamsWithTimeout creates a new ExtrasCustomFieldsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsPartialUpdateParams {
	return &ExtrasCustomFieldsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsPartialUpdateParamsWithContext creates a new ExtrasCustomFieldsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasCustomFieldsPartialUpdateParams {
	return &ExtrasCustomFieldsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsPartialUpdateParamsWithHTTPClient creates a new ExtrasCustomFieldsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsPartialUpdateParams {
	return &ExtrasCustomFieldsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasCustomFieldsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the extras custom fields partial update operation.

	Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsPartialUpdateParams struct {

	// Data.
	Data *models.WritableCustomField

	/* ID.

	   A unique integer value identifying this custom field.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsPartialUpdateParams) WithDefaults() *ExtrasCustomFieldsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasCustomFieldsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) WithData(data *models.WritableCustomField) *ExtrasCustomFieldsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) SetData(data *models.WritableCustomField) {
	o.Data = data
}

// WithID adds the id to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) WithID(id int64) *ExtrasCustomFieldsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields partial update params
func (o *ExtrasCustomFieldsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
