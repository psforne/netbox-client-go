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
	"github.com/go-openapi/swag"
)

// NewTenancyContactAssignmentsListParams creates a new TenancyContactAssignmentsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsListParams() *TenancyContactAssignmentsListParams {
	return &TenancyContactAssignmentsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsListParamsWithTimeout creates a new TenancyContactAssignmentsListParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsListParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsListParams {
	return &TenancyContactAssignmentsListParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsListParamsWithContext creates a new TenancyContactAssignmentsListParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsListParamsWithContext(ctx context.Context) *TenancyContactAssignmentsListParams {
	return &TenancyContactAssignmentsListParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsListParamsWithHTTPClient creates a new TenancyContactAssignmentsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsListParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsListParams {
	return &TenancyContactAssignmentsListParams{
		HTTPClient: client,
	}
}

/*
TenancyContactAssignmentsListParams contains all the parameters to send to the API endpoint

	for the tenancy contact assignments list operation.

	Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsListParams struct {

	// ContactID.
	ContactID *string

	// ContactIDn.
	ContactIDn *string

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// ContentTypeID.
	ContentTypeID *string

	// ContentTypeIDn.
	ContentTypeIDn *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

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

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// ObjectID.
	ObjectID *string

	// ObjectIDGt.
	ObjectIDGt *string

	// ObjectIDGte.
	ObjectIDGte *string

	// ObjectIDLt.
	ObjectIDLt *string

	// ObjectIDLte.
	ObjectIDLte *string

	// ObjectIDn.
	ObjectIDn *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Priority.
	Priority *string

	// Priorityn.
	Priorityn *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsListParams) WithDefaults() *TenancyContactAssignmentsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContext(ctx context.Context) *TenancyContactAssignmentsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContactID(contactID *string) *TenancyContactAssignmentsListParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContactID(contactID *string) {
	o.ContactID = contactID
}

// WithContactIDn adds the contactIDn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContactIDn(contactIDn *string) *TenancyContactAssignmentsListParams {
	o.SetContactIDn(contactIDn)
	return o
}

// SetContactIDn adds the contactIdN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContactIDn(contactIDn *string) {
	o.ContactIDn = contactIDn
}

// WithContentType adds the contentType to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContentType(contentType *string) *TenancyContactAssignmentsListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContentTypen(contentTypen *string) *TenancyContactAssignmentsListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithContentTypeID adds the contentTypeID to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContentTypeID(contentTypeID *string) *TenancyContactAssignmentsListParams {
	o.SetContentTypeID(contentTypeID)
	return o
}

// SetContentTypeID adds the contentTypeId to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContentTypeID(contentTypeID *string) {
	o.ContentTypeID = contentTypeID
}

// WithContentTypeIDn adds the contentTypeIDn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithContentTypeIDn(contentTypeIDn *string) *TenancyContactAssignmentsListParams {
	o.SetContentTypeIDn(contentTypeIDn)
	return o
}

// SetContentTypeIDn adds the contentTypeIdN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetContentTypeIDn(contentTypeIDn *string) {
	o.ContentTypeIDn = contentTypeIDn
}

// WithCreated adds the created to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithCreated(created *string) *TenancyContactAssignmentsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithCreatedGte(createdGte *string) *TenancyContactAssignmentsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithCreatedLte(createdLte *string) *TenancyContactAssignmentsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithID(id *string) *TenancyContactAssignmentsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithIDGt(iDGt *string) *TenancyContactAssignmentsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithIDGte(iDGte *string) *TenancyContactAssignmentsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithIDLt(iDLt *string) *TenancyContactAssignmentsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithIDLte(iDLte *string) *TenancyContactAssignmentsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithIDn(iDn *string) *TenancyContactAssignmentsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithLastUpdated(lastUpdated *string) *TenancyContactAssignmentsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *TenancyContactAssignmentsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *TenancyContactAssignmentsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithLimit(limit *int64) *TenancyContactAssignmentsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithObjectID adds the objectID to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectID(objectID *string) *TenancyContactAssignmentsListParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectID(objectID *string) {
	o.ObjectID = objectID
}

// WithObjectIDGt adds the objectIDGt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectIDGt(objectIDGt *string) *TenancyContactAssignmentsListParams {
	o.SetObjectIDGt(objectIDGt)
	return o
}

// SetObjectIDGt adds the objectIdGt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectIDGt(objectIDGt *string) {
	o.ObjectIDGt = objectIDGt
}

// WithObjectIDGte adds the objectIDGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectIDGte(objectIDGte *string) *TenancyContactAssignmentsListParams {
	o.SetObjectIDGte(objectIDGte)
	return o
}

// SetObjectIDGte adds the objectIdGte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectIDGte(objectIDGte *string) {
	o.ObjectIDGte = objectIDGte
}

// WithObjectIDLt adds the objectIDLt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectIDLt(objectIDLt *string) *TenancyContactAssignmentsListParams {
	o.SetObjectIDLt(objectIDLt)
	return o
}

// SetObjectIDLt adds the objectIdLt to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectIDLt(objectIDLt *string) {
	o.ObjectIDLt = objectIDLt
}

// WithObjectIDLte adds the objectIDLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectIDLte(objectIDLte *string) *TenancyContactAssignmentsListParams {
	o.SetObjectIDLte(objectIDLte)
	return o
}

// SetObjectIDLte adds the objectIdLte to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectIDLte(objectIDLte *string) {
	o.ObjectIDLte = objectIDLte
}

// WithObjectIDn adds the objectIDn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithObjectIDn(objectIDn *string) *TenancyContactAssignmentsListParams {
	o.SetObjectIDn(objectIDn)
	return o
}

// SetObjectIDn adds the objectIdN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetObjectIDn(objectIDn *string) {
	o.ObjectIDn = objectIDn
}

// WithOffset adds the offset to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithOffset(offset *int64) *TenancyContactAssignmentsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPriority adds the priority to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithPriority(priority *string) *TenancyContactAssignmentsListParams {
	o.SetPriority(priority)
	return o
}

// SetPriority adds the priority to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetPriority(priority *string) {
	o.Priority = priority
}

// WithPriorityn adds the priorityn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithPriorityn(priorityn *string) *TenancyContactAssignmentsListParams {
	o.SetPriorityn(priorityn)
	return o
}

// SetPriorityn adds the priorityN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetPriorityn(priorityn *string) {
	o.Priorityn = priorityn
}

// WithRole adds the role to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithRole(role *string) *TenancyContactAssignmentsListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithRolen(rolen *string) *TenancyContactAssignmentsListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithRoleID(roleID *string) *TenancyContactAssignmentsListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) WithRoleIDn(roleIDn *string) *TenancyContactAssignmentsListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the tenancy contact assignments list params
func (o *TenancyContactAssignmentsListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContactID != nil {

		// query param contact_id
		var qrContactID string

		if o.ContactID != nil {
			qrContactID = *o.ContactID
		}
		qContactID := qrContactID
		if qContactID != "" {

			if err := r.SetQueryParam("contact_id", qContactID); err != nil {
				return err
			}
		}
	}

	if o.ContactIDn != nil {

		// query param contact_id__n
		var qrContactIDn string

		if o.ContactIDn != nil {
			qrContactIDn = *o.ContactIDn
		}
		qContactIDn := qrContactIDn
		if qContactIDn != "" {

			if err := r.SetQueryParam("contact_id__n", qContactIDn); err != nil {
				return err
			}
		}
	}

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

	if o.ContentTypeID != nil {

		// query param content_type_id
		var qrContentTypeID string

		if o.ContentTypeID != nil {
			qrContentTypeID = *o.ContentTypeID
		}
		qContentTypeID := qrContentTypeID
		if qContentTypeID != "" {

			if err := r.SetQueryParam("content_type_id", qContentTypeID); err != nil {
				return err
			}
		}
	}

	if o.ContentTypeIDn != nil {

		// query param content_type_id__n
		var qrContentTypeIDn string

		if o.ContentTypeIDn != nil {
			qrContentTypeIDn = *o.ContentTypeIDn
		}
		qContentTypeIDn := qrContentTypeIDn
		if qContentTypeIDn != "" {

			if err := r.SetQueryParam("content_type_id__n", qContentTypeIDn); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
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

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.ObjectID != nil {

		// query param object_id
		var qrObjectID string

		if o.ObjectID != nil {
			qrObjectID = *o.ObjectID
		}
		qObjectID := qrObjectID
		if qObjectID != "" {

			if err := r.SetQueryParam("object_id", qObjectID); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDGt != nil {

		// query param object_id__gt
		var qrObjectIDGt string

		if o.ObjectIDGt != nil {
			qrObjectIDGt = *o.ObjectIDGt
		}
		qObjectIDGt := qrObjectIDGt
		if qObjectIDGt != "" {

			if err := r.SetQueryParam("object_id__gt", qObjectIDGt); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDGte != nil {

		// query param object_id__gte
		var qrObjectIDGte string

		if o.ObjectIDGte != nil {
			qrObjectIDGte = *o.ObjectIDGte
		}
		qObjectIDGte := qrObjectIDGte
		if qObjectIDGte != "" {

			if err := r.SetQueryParam("object_id__gte", qObjectIDGte); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDLt != nil {

		// query param object_id__lt
		var qrObjectIDLt string

		if o.ObjectIDLt != nil {
			qrObjectIDLt = *o.ObjectIDLt
		}
		qObjectIDLt := qrObjectIDLt
		if qObjectIDLt != "" {

			if err := r.SetQueryParam("object_id__lt", qObjectIDLt); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDLte != nil {

		// query param object_id__lte
		var qrObjectIDLte string

		if o.ObjectIDLte != nil {
			qrObjectIDLte = *o.ObjectIDLte
		}
		qObjectIDLte := qrObjectIDLte
		if qObjectIDLte != "" {

			if err := r.SetQueryParam("object_id__lte", qObjectIDLte); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDn != nil {

		// query param object_id__n
		var qrObjectIDn string

		if o.ObjectIDn != nil {
			qrObjectIDn = *o.ObjectIDn
		}
		qObjectIDn := qrObjectIDn
		if qObjectIDn != "" {

			if err := r.SetQueryParam("object_id__n", qObjectIDn); err != nil {
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

	if o.Priority != nil {

		// query param priority
		var qrPriority string

		if o.Priority != nil {
			qrPriority = *o.Priority
		}
		qPriority := qrPriority
		if qPriority != "" {

			if err := r.SetQueryParam("priority", qPriority); err != nil {
				return err
			}
		}
	}

	if o.Priorityn != nil {

		// query param priority__n
		var qrPriorityn string

		if o.Priorityn != nil {
			qrPriorityn = *o.Priorityn
		}
		qPriorityn := qrPriorityn
		if qPriorityn != "" {

			if err := r.SetQueryParam("priority__n", qPriorityn); err != nil {
				return err
			}
		}
	}

	if o.Role != nil {

		// query param role
		var qrRole string

		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {

			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}
	}

	if o.Rolen != nil {

		// query param role__n
		var qrRolen string

		if o.Rolen != nil {
			qrRolen = *o.Rolen
		}
		qRolen := qrRolen
		if qRolen != "" {

			if err := r.SetQueryParam("role__n", qRolen); err != nil {
				return err
			}
		}
	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string

		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {

			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}
	}

	if o.RoleIDn != nil {

		// query param role_id__n
		var qrRoleIDn string

		if o.RoleIDn != nil {
			qrRoleIDn = *o.RoleIDn
		}
		qRoleIDn := qrRoleIDn
		if qRoleIDn != "" {

			if err := r.SetQueryParam("role_id__n", qRoleIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
