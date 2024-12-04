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

// TxDecodeResponse tx decode response
//
// swagger:model TxDecodeResponse
type TxDecodeResponse struct {

	// tx
	Tx *Tx `json:"tx,omitempty"`
}

// Validate validates this tx decode response
func (m *TxDecodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxDecodeResponse) validateTx(formats strfmt.Registry) error {
	if swag.IsZero(m.Tx) { // not required
		return nil
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tx decode response based on the context it is used
func (m *TxDecodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxDecodeResponse) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if m.Tx != nil {

		if swag.IsZero(m.Tx) { // not required
			return nil
		}

		if err := m.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TxDecodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxDecodeResponse) UnmarshalBinary(b []byte) error {
	var res TxDecodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
