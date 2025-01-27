// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthInfo auth info
//
// swagger:model AuthInfo
type AuthInfo struct {

	// fee
	Fee *AuthInfoFee `json:"fee,omitempty"`

	// signer infos
	SignerInfos []*SignerInfo `json:"signer_infos"`

	// tip
	Tip *AuthInfoTip `json:"tip,omitempty"`
}

// Validate validates this auth info
func (m *AuthInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignerInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTip(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfo) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(m.Fee) { // not required
		return nil
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *AuthInfo) validateSignerInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.SignerInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.SignerInfos); i++ {
		if swag.IsZero(m.SignerInfos[i]) { // not required
			continue
		}

		if m.SignerInfos[i] != nil {
			if err := m.SignerInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthInfo) validateTip(formats strfmt.Registry) error {
	if swag.IsZero(m.Tip) { // not required
		return nil
	}

	if m.Tip != nil {
		if err := m.Tip.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auth info based on the context it is used
func (m *AuthInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignerInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTip(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfo) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if m.Fee != nil {

		if swag.IsZero(m.Fee) { // not required
			return nil
		}

		if err := m.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *AuthInfo) contextValidateSignerInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SignerInfos); i++ {

		if m.SignerInfos[i] != nil {

			if swag.IsZero(m.SignerInfos[i]) { // not required
				return nil
			}

			if err := m.SignerInfos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthInfo) contextValidateTip(ctx context.Context, formats strfmt.Registry) error {

	if m.Tip != nil {

		if swag.IsZero(m.Tip) { // not required
			return nil
		}

		if err := m.Tip.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfo) UnmarshalBinary(b []byte) error {
	var res AuthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuthInfoFee auth info fee
//
// swagger:model AuthInfoFee
type AuthInfoFee struct {

	// amount
	Amount []*AuthInfoFeeAmountItems0 `json:"amount"`

	// gas limit
	GasLimit string `json:"gas_limit,omitempty"`

	// granter
	Granter string `json:"granter,omitempty"`

	// payer
	Payer string `json:"payer,omitempty"`
}

// Validate validates this auth info fee
func (m *AuthInfoFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfoFee) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	for i := 0; i < len(m.Amount); i++ {
		if swag.IsZero(m.Amount[i]) { // not required
			continue
		}

		if m.Amount[i] != nil {
			if err := m.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this auth info fee based on the context it is used
func (m *AuthInfoFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfoFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amount); i++ {

		if m.Amount[i] != nil {

			if swag.IsZero(m.Amount[i]) { // not required
				return nil
			}

			if err := m.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfoFee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfoFee) UnmarshalBinary(b []byte) error {
	var res AuthInfoFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuthInfoFeeAmountItems0 auth info fee amount items0
//
// swagger:model AuthInfoFeeAmountItems0
type AuthInfoFeeAmountItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this auth info fee amount items0
func (m *AuthInfoFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth info fee amount items0 based on context it is used
func (m *AuthInfoFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfoFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfoFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res AuthInfoFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuthInfoTip auth info tip
//
// swagger:model AuthInfoTip
type AuthInfoTip struct {

	// amount
	Amount []*AuthInfoTipAmountItems0 `json:"amount"`

	// tipper
	Tipper string `json:"tipper,omitempty"`
}

// Validate validates this auth info tip
func (m *AuthInfoTip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfoTip) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	for i := 0; i < len(m.Amount); i++ {
		if swag.IsZero(m.Amount[i]) { // not required
			continue
		}

		if m.Amount[i] != nil {
			if err := m.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tip" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tip" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this auth info tip based on the context it is used
func (m *AuthInfoTip) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfoTip) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amount); i++ {

		if m.Amount[i] != nil {

			if swag.IsZero(m.Amount[i]) { // not required
				return nil
			}

			if err := m.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tip" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tip" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfoTip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfoTip) UnmarshalBinary(b []byte) error {
	var res AuthInfoTip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuthInfoTipAmountItems0 auth info tip amount items0
//
// swagger:model AuthInfoTipAmountItems0
type AuthInfoTipAmountItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this auth info tip amount items0
func (m *AuthInfoTipAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth info tip amount items0 based on context it is used
func (m *AuthInfoTipAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfoTipAmountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfoTipAmountItems0) UnmarshalBinary(b []byte) error {
	var res AuthInfoTipAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
