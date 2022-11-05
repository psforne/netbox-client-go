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
)

// NewExtrasExportTemplatesListParams creates a new ExtrasExportTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesListParams() *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesListParamsWithTimeout creates a new ExtrasExportTemplatesListParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesListParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesListParamsWithContext creates a new ExtrasExportTemplatesListParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesListParamsWithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesListParamsWithHTTPClient creates a new ExtrasExportTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesListParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		HTTPClient: client,
	}
}

/*
ExtrasExportTemplatesListParams contains all the parameters to send to the API endpoint

	for the extras export templates list operation.

	Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesListParams struct {

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesListParams) WithDefaults() *ExtrasExportTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContentType(contentType *string) *ExtrasExportTemplatesListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContentTypen(contentTypen *string) *ExtrasExportTemplatesListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithDescription adds the description to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescription(description *string) *ExtrasExportTemplatesListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionIc(descriptionIc *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionIe(descriptionIe *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionIew(descriptionIew *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionIsw(descriptionIsw *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionn(descriptionn *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionNic(descriptionNic *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionNie(descriptionNie *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionNiew(descriptionNiew *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithDescriptionNisw(descriptionNisw *string) *ExtrasExportTemplatesListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithID(id *string) *ExtrasExportTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDGt(iDGt *string) *ExtrasExportTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDGte(iDGte *string) *ExtrasExportTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDLt(iDLt *string) *ExtrasExportTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDLte(iDLte *string) *ExtrasExportTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDn(iDn *string) *ExtrasExportTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithLimit(limit *int64) *ExtrasExportTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithName(name *string) *ExtrasExportTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIc(nameIc *string) *ExtrasExportTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIe(nameIe *string) *ExtrasExportTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIew(nameIew *string) *ExtrasExportTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIsw(nameIsw *string) *ExtrasExportTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNamen(namen *string) *ExtrasExportTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNic(nameNic *string) *ExtrasExportTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNie(nameNie *string) *ExtrasExportTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNiew(nameNiew *string) *ExtrasExportTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNisw(nameNisw *string) *ExtrasExportTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOffset(offset *int64) *ExtrasExportTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithQ(q *string) *ExtrasExportTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string

		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {

			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}
	}

	if o.ContentTypen != nil {

		// query param content_type__n
		var qrContentTypen string

		if o.ContentTypen != nil {
			qrContentTypen = *o.ContentTypen
		}
		qContentTypen := qrContentTypen
		if qContentTypen != "" {

			if err := r.SetQueryParam("content_type__n", qContentTypen); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
