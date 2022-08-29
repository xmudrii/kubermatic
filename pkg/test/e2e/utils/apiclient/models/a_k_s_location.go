// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AKSLocation AKSLocation is the object representing Azure Location.
//
// swagger:model AKSLocation
type AKSLocation struct {

	// The location name.
	Name string `json:"name,omitempty"`

	// READ-ONLY; The category of the region.
	RegionCategory string `json:"regionCategory,omitempty"`
}

// Validate validates this a k s location
func (m *AKSLocation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a k s location based on context it is used
func (m *AKSLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AKSLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AKSLocation) UnmarshalBinary(b []byte) error {
	var res AKSLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}