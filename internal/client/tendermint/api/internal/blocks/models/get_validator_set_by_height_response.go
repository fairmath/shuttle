// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetValidatorSetByHeightResponse get validator set by height response
//
// swagger:model GetValidatorSetByHeightResponse
type GetValidatorSetByHeightResponse struct {

	// block height
	BlockHeight string `json:"block_height,omitempty"`

	// pagination
	Pagination *GetValidatorSetByHeightResponsePagination `json:"pagination,omitempty"`

	// validators
	Validators []*GetValidatorSetByHeightResponseValidatorsItems0 `json:"validators"`
}

// Validate validates this get validator set by height response
func (m *GetValidatorSetByHeightResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorSetByHeightResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *GetValidatorSetByHeightResponse) validateValidators(formats strfmt.Registry) error {
	if swag.IsZero(m.Validators) { // not required
		return nil
	}

	for i := 0; i < len(m.Validators); i++ {
		if swag.IsZero(m.Validators[i]) { // not required
			continue
		}

		if m.Validators[i] != nil {
			if err := m.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get validator set by height response based on the context it is used
func (m *GetValidatorSetByHeightResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorSetByHeightResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *GetValidatorSetByHeightResponse) contextValidateValidators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Validators); i++ {

		if m.Validators[i] != nil {

			if swag.IsZero(m.Validators[i]) { // not required
				return nil
			}

			if err := m.Validators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponse) UnmarshalBinary(b []byte) error {
	var res GetValidatorSetByHeightResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetValidatorSetByHeightResponsePagination get validator set by height response pagination
//
// swagger:model GetValidatorSetByHeightResponsePagination
type GetValidatorSetByHeightResponsePagination struct {

	// next key
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this get validator set by height response pagination
func (m *GetValidatorSetByHeightResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get validator set by height response pagination based on context it is used
func (m *GetValidatorSetByHeightResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponsePagination) UnmarshalBinary(b []byte) error {
	var res GetValidatorSetByHeightResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetValidatorSetByHeightResponseValidatorsItems0 get validator set by height response validators items0
//
// swagger:model GetValidatorSetByHeightResponseValidatorsItems0
type GetValidatorSetByHeightResponseValidatorsItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// proposer priority
	ProposerPriority string `json:"proposer_priority,omitempty"`

	// pub key
	PubKey *GetValidatorSetByHeightResponseValidatorsItems0PubKey `json:"pub_key,omitempty"`

	// voting power
	VotingPower string `json:"voting_power,omitempty"`
}

// Validate validates this get validator set by height response validators items0
func (m *GetValidatorSetByHeightResponseValidatorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorSetByHeightResponseValidatorsItems0) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PubKey) { // not required
		return nil
	}

	if m.PubKey != nil {
		if err := m.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get validator set by height response validators items0 based on the context it is used
func (m *GetValidatorSetByHeightResponseValidatorsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorSetByHeightResponseValidatorsItems0) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PubKey != nil {

		if swag.IsZero(m.PubKey) { // not required
			return nil
		}

		if err := m.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponseValidatorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponseValidatorsItems0) UnmarshalBinary(b []byte) error {
	var res GetValidatorSetByHeightResponseValidatorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetValidatorSetByHeightResponseValidatorsItems0PubKey get validator set by height response validators items0 pub key
//
// swagger:model GetValidatorSetByHeightResponseValidatorsItems0PubKey
type GetValidatorSetByHeightResponseValidatorsItems0PubKey struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// get validator set by height response validators items0 pub key
	GetValidatorSetByHeightResponseValidatorsItems0PubKey map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *GetValidatorSetByHeightResponseValidatorsItems0PubKey) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GetValidatorSetByHeightResponseValidatorsItems0PubKey

	rcv.AtType = stage1.AtType
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.GetValidatorSetByHeightResponseValidatorsItems0PubKey = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m GetValidatorSetByHeightResponseValidatorsItems0PubKey) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = m.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.GetValidatorSetByHeightResponseValidatorsItems0PubKey) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.GetValidatorSetByHeightResponseValidatorsItems0PubKey)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this get validator set by height response validators items0 pub key
func (m *GetValidatorSetByHeightResponseValidatorsItems0PubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get validator set by height response validators items0 pub key based on context it is used
func (m *GetValidatorSetByHeightResponseValidatorsItems0PubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponseValidatorsItems0PubKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorSetByHeightResponseValidatorsItems0PubKey) UnmarshalBinary(b []byte) error {
	var res GetValidatorSetByHeightResponseValidatorsItems0PubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
