// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersPermissionsBulkPartialUpdateParams creates a new UsersPermissionsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsBulkPartialUpdateParams() *UsersPermissionsBulkPartialUpdateParams {
	return &UsersPermissionsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsBulkPartialUpdateParamsWithTimeout creates a new UsersPermissionsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersPermissionsBulkPartialUpdateParams {
	return &UsersPermissionsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsBulkPartialUpdateParamsWithContext creates a new UsersPermissionsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewUsersPermissionsBulkPartialUpdateParamsWithContext(ctx context.Context) *UsersPermissionsBulkPartialUpdateParams {
	return &UsersPermissionsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewUsersPermissionsBulkPartialUpdateParamsWithHTTPClient creates a new UsersPermissionsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersPermissionsBulkPartialUpdateParams {
	return &UsersPermissionsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
UsersPermissionsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the users permissions bulk partial update operation.

	Typically these are written to a http.Request.
*/
type UsersPermissionsBulkPartialUpdateParams struct {

	// Data.
	Data []*models.WritableObjectPermission

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsBulkPartialUpdateParams) WithDefaults() *UsersPermissionsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersPermissionsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) WithContext(ctx context.Context) *UsersPermissionsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersPermissionsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) WithData(data []*models.WritableObjectPermission) *UsersPermissionsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users permissions bulk partial update params
func (o *UsersPermissionsBulkPartialUpdateParams) SetData(data []*models.WritableObjectPermission) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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