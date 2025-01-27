// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GasInfo gas info
//
// swagger:model GasInfo
type GasInfo struct {

	// gas used
	GasUsed string `json:"gas_used,omitempty"`

	// gas wanted
	GasWanted string `json:"gas_wanted,omitempty"`
}

// Validate validates this gas info
func (m *GasInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gas info based on context it is used
func (m *GasInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GasInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GasInfo) UnmarshalBinary(b []byte) error {
	var res GasInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
