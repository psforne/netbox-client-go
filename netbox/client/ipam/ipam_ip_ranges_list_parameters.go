// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIpamIPRangesListParams creates a new IpamIPRangesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPRangesListParams() *IpamIPRangesListParams {
	return &IpamIPRangesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPRangesListParamsWithTimeout creates a new IpamIPRangesListParams object
// with the ability to set a timeout on a request.
func NewIpamIPRangesListParamsWithTimeout(timeout time.Duration) *IpamIPRangesListParams {
	return &IpamIPRangesListParams{
		timeout: timeout,
	}
}

// NewIpamIPRangesListParamsWithContext creates a new IpamIPRangesListParams object
// with the ability to set a context for a request.
func NewIpamIPRangesListParamsWithContext(ctx context.Context) *IpamIPRangesListParams {
	return &IpamIPRangesListParams{
		Context: ctx,
	}
}

// NewIpamIPRangesListParamsWithHTTPClient creates a new IpamIPRangesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPRangesListParamsWithHTTPClient(client *http.Client) *IpamIPRangesListParams {
	return &IpamIPRangesListParams{
		HTTPClient: client,
	}
}

/*
IpamIPRangesListParams contains all the parameters to send to the API endpoint

	for the ipam ip ranges list operation.

	Typically these are written to a http.Request.
*/
type IpamIPRangesListParams struct {

	// Contains.
	Contains *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

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

	// Family.
	Family *float64

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

	// Status.
	Status *string

	// Statusn.
	Statusn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// Tenantn.
	Tenantn *string

	// TenantGroup.
	TenantGroup *string

	// TenantGroupn.
	TenantGroupn *string

	// TenantGroupID.
	TenantGroupID *string

	// TenantGroupIDn.
	TenantGroupIDn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	// Vrf.
	Vrf *string

	// Vrfn.
	Vrfn *string

	// VrfID.
	VrfID *string

	// VrfIDn.
	VrfIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip ranges list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesListParams) WithDefaults() *IpamIPRangesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip ranges list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTimeout(timeout time.Duration) *IpamIPRangesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithContext(ctx context.Context) *IpamIPRangesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithHTTPClient(client *http.Client) *IpamIPRangesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContains adds the contains to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithContains(contains *string) *IpamIPRangesListParams {
	o.SetContains(contains)
	return o
}

// SetContains adds the contains to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetContains(contains *string) {
	o.Contains = contains
}

// WithCreated adds the created to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithCreated(created *string) *IpamIPRangesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithCreatedGte(createdGte *string) *IpamIPRangesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithCreatedLte(createdLte *string) *IpamIPRangesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescription(description *string) *IpamIPRangesListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionIc(descriptionIc *string) *IpamIPRangesListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionIe(descriptionIe *string) *IpamIPRangesListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionIew(descriptionIew *string) *IpamIPRangesListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionIsw(descriptionIsw *string) *IpamIPRangesListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionn(descriptionn *string) *IpamIPRangesListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionNic(descriptionNic *string) *IpamIPRangesListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionNie(descriptionNie *string) *IpamIPRangesListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionNiew(descriptionNiew *string) *IpamIPRangesListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithDescriptionNisw(descriptionNisw *string) *IpamIPRangesListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithFamily adds the family to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithFamily(family *float64) *IpamIPRangesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetFamily(family *float64) {
	o.Family = family
}

// WithID adds the id to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithID(id *string) *IpamIPRangesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithIDGt(iDGt *string) *IpamIPRangesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithIDGte(iDGte *string) *IpamIPRangesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithIDLt(iDLt *string) *IpamIPRangesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithIDLte(iDLte *string) *IpamIPRangesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithIDn(iDn *string) *IpamIPRangesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithLastUpdated(lastUpdated *string) *IpamIPRangesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamIPRangesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamIPRangesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithLimit(limit *int64) *IpamIPRangesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithOffset(offset *int64) *IpamIPRangesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithQ(q *string) *IpamIPRangesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithRole(role *string) *IpamIPRangesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithRolen(rolen *string) *IpamIPRangesListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithRoleID(roleID *string) *IpamIPRangesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithRoleIDn(roleIDn *string) *IpamIPRangesListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithStatus adds the status to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithStatus(status *string) *IpamIPRangesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithStatusn(statusn *string) *IpamIPRangesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTag(tag *string) *IpamIPRangesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTagn(tagn *string) *IpamIPRangesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenant(tenant *string) *IpamIPRangesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantn(tenantn *string) *IpamIPRangesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantGroup(tenantGroup *string) *IpamIPRangesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantGroupn(tenantGroupn *string) *IpamIPRangesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantGroupID(tenantGroupID *string) *IpamIPRangesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamIPRangesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantID(tenantID *string) *IpamIPRangesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithTenantIDn(tenantIDn *string) *IpamIPRangesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVrf adds the vrf to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithVrf(vrf *string) *IpamIPRangesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfn adds the vrfn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithVrfn(vrfn *string) *IpamIPRangesListParams {
	o.SetVrfn(vrfn)
	return o
}

// SetVrfn adds the vrfN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetVrfn(vrfn *string) {
	o.Vrfn = vrfn
}

// WithVrfID adds the vrfID to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithVrfID(vrfID *string) *IpamIPRangesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetVrfID(vrfID *string) {
	o.VrfID = vrfID
}

// WithVrfIDn adds the vrfIDn to the ipam ip ranges list params
func (o *IpamIPRangesListParams) WithVrfIDn(vrfIDn *string) *IpamIPRangesListParams {
	o.SetVrfIDn(vrfIDn)
	return o
}

// SetVrfIDn adds the vrfIdN to the ipam ip ranges list params
func (o *IpamIPRangesListParams) SetVrfIDn(vrfIDn *string) {
	o.VrfIDn = vrfIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPRangesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contains != nil {

		// query param contains
		var qrContains string

		if o.Contains != nil {
			qrContains = *o.Contains
		}
		qContains := qrContains
		if qContains != "" {

			if err := r.SetQueryParam("contains", qContains); err != nil {
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

	if o.Family != nil {

		// query param family
		var qrFamily float64

		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := swag.FormatFloat64(qrFamily)
		if qFamily != "" {

			if err := r.SetQueryParam("family", qFamily); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Statusn != nil {

		// query param status__n
		var qrStatusn string

		if o.Statusn != nil {
			qrStatusn = *o.Statusn
		}
		qStatusn := qrStatusn
		if qStatusn != "" {

			if err := r.SetQueryParam("status__n", qStatusn); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string

		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {

			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string

		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {

			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string

		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {

			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string

		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {

			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if o.Vrf != nil {

		// query param vrf
		var qrVrf string

		if o.Vrf != nil {
			qrVrf = *o.Vrf
		}
		qVrf := qrVrf
		if qVrf != "" {

			if err := r.SetQueryParam("vrf", qVrf); err != nil {
				return err
			}
		}
	}

	if o.Vrfn != nil {

		// query param vrf__n
		var qrVrfn string

		if o.Vrfn != nil {
			qrVrfn = *o.Vrfn
		}
		qVrfn := qrVrfn
		if qVrfn != "" {

			if err := r.SetQueryParam("vrf__n", qVrfn); err != nil {
				return err
			}
		}
	}

	if o.VrfID != nil {

		// query param vrf_id
		var qrVrfID string

		if o.VrfID != nil {
			qrVrfID = *o.VrfID
		}
		qVrfID := qrVrfID
		if qVrfID != "" {

			if err := r.SetQueryParam("vrf_id", qVrfID); err != nil {
				return err
			}
		}
	}

	if o.VrfIDn != nil {

		// query param vrf_id__n
		var qrVrfIDn string

		if o.VrfIDn != nil {
			qrVrfIDn = *o.VrfIDn
		}
		qVrfIDn := qrVrfIDn
		if qVrfIDn != "" {

			if err := r.SetQueryParam("vrf_id__n", qVrfIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
