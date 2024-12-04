// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CompactBitArray compact bit array
//
// swagger:model CompactBitArray
type CompactBitArray struct {

	// elems
	// Format: byte
	Elems strfmt.Base64 `json:"elems,omitempty"`

	// extra bits stored
	ExtraBitsStored int64 `json:"extra_bits_stored,omitempty"`
}

// Validate validates this compact bit array
func (m *CompactBitArray) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this compact bit array based on context it is used
func (m *CompactBitArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompactBitArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompactBitArray) UnmarshalBinary(b []byte) error {
	var res CompactBitArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
