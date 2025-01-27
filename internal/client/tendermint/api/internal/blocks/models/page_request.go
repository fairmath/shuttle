// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PageRequest page request
//
// swagger:model PageRequest
type PageRequest struct {

	// count total
	CountTotal bool `json:"count_total,omitempty"`

	// key
	// Format: byte
	Key strfmt.Base64 `json:"key,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// offset
	Offset string `json:"offset,omitempty"`

	// reverse
	Reverse bool `json:"reverse,omitempty"`
}

// Validate validates this page request
func (m *PageRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this page request based on context it is used
func (m *PageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageRequest) UnmarshalBinary(b []byte) error {
	var res PageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
