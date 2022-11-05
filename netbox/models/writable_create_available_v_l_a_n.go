// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCreateAvailableVLAN writable create available v l a n
//
// swagger:model WritableCreateAvailableVLAN
type WritableCreateAvailableVLAN struct {

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Role
	Role *int64 `json:"role,omitempty"`

	// Site
	Site *int64 `json:"site,omitempty"`

	// Status
	// Enum: [active reserved deprecated]
	Status string `json:"status,omitempty"`

	// tags
	// Unique: true
	Tags []int64 `json:"tags,omitempty"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`
}

// Validate validates this writable create available v l a n
func (m *WritableCreateAvailableVLAN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCreateAvailableVLAN) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableCreateAvailableVLAN) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

var writableCreateAvailableVLANTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCreateAvailableVLANTypeStatusPropEnum = append(writableCreateAvailableVLANTypeStatusPropEnum, v)
	}
}

const (

	// WritableCreateAvailableVLANStatusActive captures enum value "active"
	WritableCreateAvailableVLANStatusActive string = "active"

	// WritableCreateAvailableVLANStatusReserved captures enum value "reserved"
	WritableCreateAvailableVLANStatusReserved string = "reserved"

	// WritableCreateAvailableVLANStatusDeprecated captures enum value "deprecated"
	WritableCreateAvailableVLANStatusDeprecated string = "deprecated"
)

// prop value enum
func (m *WritableCreateAvailableVLAN) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCreateAvailableVLANTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCreateAvailableVLAN) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCreateAvailableVLAN) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this writable create available v l a n based on context it is used
func (m *WritableCreateAvailableVLAN) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WritableCreateAvailableVLAN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCreateAvailableVLAN) UnmarshalBinary(b []byte) error {
	var res WritableCreateAvailableVLAN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
