// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PageResponse page response
//
// swagger:model PageResponse
type PageResponse struct {

	// next key
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this page response
func (m *PageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this page response based on context it is used
func (m *PageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageResponse) UnmarshalBinary(b []byte) error {
	var res PageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
