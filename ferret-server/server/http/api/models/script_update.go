// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ScriptUpdate Script Update
//
// The properties that are allowed when updating a Script.
// swagger:model script-update
type ScriptUpdate struct {
	ScriptCommon
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScriptUpdate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ScriptCommon
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ScriptCommon = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScriptUpdate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ScriptCommon)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this script update
func (m *ScriptUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScriptUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptUpdate) UnmarshalBinary(b []byte) error {
	var res ScriptUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
