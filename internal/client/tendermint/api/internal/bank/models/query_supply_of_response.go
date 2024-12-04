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

// QuerySupplyOfResponse query supply of response
//
// swagger:model QuerySupplyOfResponse
type QuerySupplyOfResponse struct {

	// amount
	Amount *QuerySupplyOfResponseAmount `json:"amount,omitempty"`
}

// Validate validates this query supply of response
func (m *QuerySupplyOfResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuerySupplyOfResponse) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this query supply of response based on the context it is used
func (m *QuerySupplyOfResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuerySupplyOfResponse) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {

		if swag.IsZero(m.Amount) { // not required
			return nil
		}

		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuerySupplyOfResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuerySupplyOfResponse) UnmarshalBinary(b []byte) error {
	var res QuerySupplyOfResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuerySupplyOfResponseAmount query supply of response amount
//
// swagger:model QuerySupplyOfResponseAmount
type QuerySupplyOfResponseAmount struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this query supply of response amount
func (m *QuerySupplyOfResponseAmount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query supply of response amount based on context it is used
func (m *QuerySupplyOfResponseAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuerySupplyOfResponseAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuerySupplyOfResponseAmount) UnmarshalBinary(b []byte) error {
	var res QuerySupplyOfResponseAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
