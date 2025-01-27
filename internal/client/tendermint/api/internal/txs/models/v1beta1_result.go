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

// V1beta1Result v1beta1 result
//
// swagger:model v1beta1.Result
type V1beta1Result struct {

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// events
	Events []*V1beta1ResultEventsItems0 `json:"events"`

	// log
	Log string `json:"log,omitempty"`

	// msg responses
	MsgResponses []*V1beta1ResultMsgResponsesItems0 `json:"msg_responses"`
}

// Validate validates this v1beta1 result
func (m *V1beta1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsgResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Result) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1beta1Result) validateMsgResponses(formats strfmt.Registry) error {
	if swag.IsZero(m.MsgResponses) { // not required
		return nil
	}

	for i := 0; i < len(m.MsgResponses); i++ {
		if swag.IsZero(m.MsgResponses[i]) { // not required
			continue
		}

		if m.MsgResponses[i] != nil {
			if err := m.MsgResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("msg_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("msg_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1beta1 result based on the context it is used
func (m *V1beta1Result) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsgResponses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Result) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if swag.IsZero(m.Events[i]) { // not required
				return nil
			}

			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1beta1Result) contextValidateMsgResponses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MsgResponses); i++ {

		if m.MsgResponses[i] != nil {

			if swag.IsZero(m.MsgResponses[i]) { // not required
				return nil
			}

			if err := m.MsgResponses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("msg_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("msg_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Result) UnmarshalBinary(b []byte) error {
	var res V1beta1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1beta1ResultEventsItems0 v1beta1 result events items0
//
// swagger:model V1beta1ResultEventsItems0
type V1beta1ResultEventsItems0 struct {

	// attributes
	Attributes []*V1beta1ResultEventsItems0AttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1beta1 result events items0
func (m *V1beta1ResultEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ResultEventsItems0) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1beta1 result events items0 based on the context it is used
func (m *V1beta1ResultEventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ResultEventsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {

			if swag.IsZero(m.Attributes[i]) { // not required
				return nil
			}

			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ResultEventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ResultEventsItems0) UnmarshalBinary(b []byte) error {
	var res V1beta1ResultEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1beta1ResultEventsItems0AttributesItems0 v1beta1 result events items0 attributes items0
//
// swagger:model V1beta1ResultEventsItems0AttributesItems0
type V1beta1ResultEventsItems0AttributesItems0 struct {

	// index
	Index bool `json:"index,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1beta1 result events items0 attributes items0
func (m *V1beta1ResultEventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1beta1 result events items0 attributes items0 based on context it is used
func (m *V1beta1ResultEventsItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ResultEventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ResultEventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res V1beta1ResultEventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1beta1ResultMsgResponsesItems0 v1beta1 result msg responses items0
//
// swagger:model V1beta1ResultMsgResponsesItems0
type V1beta1ResultMsgResponsesItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// v1beta1 result msg responses items0
	V1beta1ResultMsgResponsesItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *V1beta1ResultMsgResponsesItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv V1beta1ResultMsgResponsesItems0

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
		m.V1beta1ResultMsgResponsesItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m V1beta1ResultMsgResponsesItems0) MarshalJSON() ([]byte, error) {
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

	if len(m.V1beta1ResultMsgResponsesItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.V1beta1ResultMsgResponsesItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this v1beta1 result msg responses items0
func (m *V1beta1ResultMsgResponsesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1beta1 result msg responses items0 based on context it is used
func (m *V1beta1ResultMsgResponsesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ResultMsgResponsesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ResultMsgResponsesItems0) UnmarshalBinary(b []byte) error {
	var res V1beta1ResultMsgResponsesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
