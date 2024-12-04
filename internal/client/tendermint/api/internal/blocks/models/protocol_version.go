// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtocolVersion protocol version
//
// swagger:model ProtocolVersion
type ProtocolVersion struct {

	// app
	App string `json:"app,omitempty"`

	// block
	Block string `json:"block,omitempty"`

	// p2p
	P2p string `json:"p2p,omitempty"`
}

// Validate validates this protocol version
func (m *ProtocolVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this protocol version based on context it is used
func (m *ProtocolVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtocolVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtocolVersion) UnmarshalBinary(b []byte) error {
	var res ProtocolVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
