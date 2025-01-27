// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BlockID block ID
//
// swagger:model BlockID
type BlockID struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// part set header
	PartSetHeader *BlockIDPartSetHeader `json:"part_set_header,omitempty"`
}

// Validate validates this block ID
func (m *BlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartSetHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockID) validatePartSetHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.PartSetHeader) { // not required
		return nil
	}

	if m.PartSetHeader != nil {
		if err := m.PartSetHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("part_set_header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block ID based on the context it is used
func (m *BlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartSetHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockID) contextValidatePartSetHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.PartSetHeader != nil {

		if swag.IsZero(m.PartSetHeader) { // not required
			return nil
		}

		if err := m.PartSetHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("part_set_header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockID) UnmarshalBinary(b []byte) error {
	var res BlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockIDPartSetHeader block ID part set header
//
// swagger:model BlockIDPartSetHeader
type BlockIDPartSetHeader struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this block ID part set header
func (m *BlockIDPartSetHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block ID part set header based on context it is used
func (m *BlockIDPartSetHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockIDPartSetHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockIDPartSetHeader) UnmarshalBinary(b []byte) error {
	var res BlockIDPartSetHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
