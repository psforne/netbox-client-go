// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageAttachment image attachment
//
// swagger:model ImageAttachment
type ImageAttachment struct {

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Image
	// Read Only: true
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// Image height
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	ImageHeight *int64 `json:"image_height"`

	// Image width
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	ImageWidth *int64 `json:"image_width"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Max Length: 50
	Name string `json:"name,omitempty"`

	// Object id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	ObjectID *int64 `json:"object_id"`

	// Parent
	// Read Only: true
	Parent map[string]*string `json:"parent,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this image attachment
func (m *ImageAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAttachment) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImageHeight(formats strfmt.Registry) error {

	if err := validate.Required("image_height", "body", m.ImageHeight); err != nil {
		return err
	}

	if err := validate.MinimumInt("image_height", "body", *m.ImageHeight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("image_height", "body", *m.ImageHeight, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImageWidth(formats strfmt.Registry) error {

	if err := validate.Required("image_width", "body", m.ImageWidth); err != nil {
		return err
	}

	if err := validate.MinimumInt("image_width", "body", *m.ImageWidth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("image_width", "body", *m.ImageWidth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("object_id", "body", *m.ObjectID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("object_id", "body", *m.ObjectID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this image attachment based on the context it is used
func (m *ImageAttachment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAttachment) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "image", "body", strfmt.URI(m.Image)); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ImageAttachment) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAttachment) UnmarshalBinary(b []byte) error {
	var res ImageAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
