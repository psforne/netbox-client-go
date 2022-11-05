// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// NewCircuitsCircuitTerminationsCreateParams creates a new CircuitsCircuitTerminationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsCreateParams() *CircuitsCircuitTerminationsCreateParams {
	return &CircuitsCircuitTerminationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsCreateParamsWithTimeout creates a new CircuitsCircuitTerminationsCreateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsCreateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsCreateParams {
	return &CircuitsCircuitTerminationsCreateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsCreateParamsWithContext creates a new CircuitsCircuitTerminationsCreateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsCreateParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsCreateParams {
	return &CircuitsCircuitTerminationsCreateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsCreateParamsWithHTTPClient creates a new CircuitsCircuitTerminationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsCreateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsCreateParams {
	return &CircuitsCircuitTerminationsCreateParams{
		HTTPClient: client,
	}
}

/*
CircuitsCircuitTerminationsCreateParams contains all the parameters to send to the API endpoint

	for the circuits circuit terminations create operation.

	Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsCreateParams struct {

	// Data.
	Data []*models.WritableCircuitTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsCreateParams) WithDefaults() *CircuitsCircuitTerminationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) WithData(data []*models.WritableCircuitTermination) *CircuitsCircuitTerminationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit terminations create params
func (o *CircuitsCircuitTerminationsCreateParams) SetData(data []*models.WritableCircuitTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
