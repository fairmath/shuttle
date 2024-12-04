// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1Validator v1beta1 validator
//
// swagger:model v1beta1.Validator
type V1beta1Validator struct {

	// address
	Address string `json:"address,omitempty"`

	// proposer priority
	ProposerPriority string `json:"proposer_priority,omitempty"`

	// pub key
	PubKey *V1beta1ValidatorPubKey `json:"pub_key,omitempty"`

	// voting power
	VotingPower string `json:"voting_power,omitempty"`
}

// Validate validates this v1beta1 validator
func (m *V1beta1Validator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Validator) validatePubKey(formats strfmt.Registry) error {
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

// ContextValidate validate this v1beta1 validator based on the context it is used
func (m *V1beta1Validator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Validator) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1beta1Validator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Validator) UnmarshalBinary(b []byte) error {
	var res V1beta1Validator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1beta1ValidatorPubKey v1beta1 validator pub key
//
// swagger:model V1beta1ValidatorPubKey
type V1beta1ValidatorPubKey struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// v1beta1 validator pub key
	V1beta1ValidatorPubKey map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *V1beta1ValidatorPubKey) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv V1beta1ValidatorPubKey

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
		m.V1beta1ValidatorPubKey = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m V1beta1ValidatorPubKey) MarshalJSON() ([]byte, error) {
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

	if len(m.V1beta1ValidatorPubKey) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.V1beta1ValidatorPubKey)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this v1beta1 validator pub key
func (m *V1beta1ValidatorPubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1beta1 validator pub key based on context it is used
func (m *V1beta1ValidatorPubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ValidatorPubKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ValidatorPubKey) UnmarshalBinary(b []byte) error {
	var res V1beta1ValidatorPubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
