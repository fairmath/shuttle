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

// BroadcastTxResponse broadcast tx response
//
// swagger:model BroadcastTxResponse
type BroadcastTxResponse struct {

	// tx response
	TxResponse *BroadcastTxResponseTxResponse `json:"tx_response,omitempty"`
}

// Validate validates this broadcast tx response
func (m *BroadcastTxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTxResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponse) validateTxResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.TxResponse) { // not required
		return nil
	}

	if m.TxResponse != nil {
		if err := m.TxResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx_response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx_response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast tx response based on the context it is used
func (m *BroadcastTxResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTxResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponse) contextValidateTxResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.TxResponse != nil {

		if swag.IsZero(m.TxResponse) { // not required
			return nil
		}

		if err := m.TxResponse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx_response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx_response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastTxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponse) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponse broadcast tx response tx response
//
// swagger:model BroadcastTxResponseTxResponse
type BroadcastTxResponseTxResponse struct {

	// code
	Code int64 `json:"code,omitempty"`

	// codespace
	Codespace string `json:"codespace,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// events
	Events []*BroadcastTxResponseTxResponseEventsItems0 `json:"events"`

	// gas used
	GasUsed string `json:"gas_used,omitempty"`

	// gas wanted
	GasWanted string `json:"gas_wanted,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// logs
	Logs []*BroadcastTxResponseTxResponseLogsItems0 `json:"logs"`

	// raw log
	RawLog string `json:"raw_log,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// tx
	Tx *BroadcastTxResponseTxResponseTx `json:"tx,omitempty"`

	// txhash
	Txhash string `json:"txhash,omitempty"`
}

// Validate validates this broadcast tx response tx response
func (m *BroadcastTxResponseTxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponse) validateEvents(formats strfmt.Registry) error {
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
					return ve.ValidateName("tx_response" + "." + "events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tx_response" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastTxResponseTxResponse) validateLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tx_response" + "." + "logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tx_response" + "." + "logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastTxResponseTxResponse) validateTx(formats strfmt.Registry) error {
	if swag.IsZero(m.Tx) { // not required
		return nil
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx_response" + "." + "tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx_response" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast tx response tx response based on the context it is used
func (m *BroadcastTxResponseTxResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponse) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if swag.IsZero(m.Events[i]) { // not required
				return nil
			}

			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tx_response" + "." + "events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tx_response" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastTxResponseTxResponse) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logs); i++ {

		if m.Logs[i] != nil {

			if swag.IsZero(m.Logs[i]) { // not required
				return nil
			}

			if err := m.Logs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tx_response" + "." + "logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tx_response" + "." + "logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastTxResponseTxResponse) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if m.Tx != nil {

		if swag.IsZero(m.Tx) { // not required
			return nil
		}

		if err := m.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx_response" + "." + "tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx_response" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponse) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseEventsItems0 broadcast tx response tx response events items0
//
// swagger:model BroadcastTxResponseTxResponseEventsItems0
type BroadcastTxResponseTxResponseEventsItems0 struct {

	// attributes
	Attributes []*BroadcastTxResponseTxResponseEventsItems0AttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this broadcast tx response tx response events items0
func (m *BroadcastTxResponseTxResponseEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseEventsItems0) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this broadcast tx response tx response events items0 based on the context it is used
func (m *BroadcastTxResponseTxResponseEventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseEventsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BroadcastTxResponseTxResponseEventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseEventsItems0) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseEventsItems0AttributesItems0 broadcast tx response tx response events items0 attributes items0
//
// swagger:model BroadcastTxResponseTxResponseEventsItems0AttributesItems0
type BroadcastTxResponseTxResponseEventsItems0AttributesItems0 struct {

	// index
	Index bool `json:"index,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this broadcast tx response tx response events items0 attributes items0
func (m *BroadcastTxResponseTxResponseEventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast tx response tx response events items0 attributes items0 based on context it is used
func (m *BroadcastTxResponseTxResponseEventsItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseEventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseEventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseEventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseLogsItems0 broadcast tx response tx response logs items0
//
// swagger:model BroadcastTxResponseTxResponseLogsItems0
type BroadcastTxResponseTxResponseLogsItems0 struct {

	// events
	Events []*BroadcastTxResponseTxResponseLogsItems0EventsItems0 `json:"events"`

	// log
	Log string `json:"log,omitempty"`

	// msg index
	MsgIndex int64 `json:"msg_index,omitempty"`
}

// Validate validates this broadcast tx response tx response logs items0
func (m *BroadcastTxResponseTxResponseLogsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseLogsItems0) validateEvents(formats strfmt.Registry) error {
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

// ContextValidate validate this broadcast tx response tx response logs items0 based on the context it is used
func (m *BroadcastTxResponseTxResponseLogsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseLogsItems0) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseLogsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseLogsItems0) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseLogsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseLogsItems0EventsItems0 broadcast tx response tx response logs items0 events items0
//
// swagger:model BroadcastTxResponseTxResponseLogsItems0EventsItems0
type BroadcastTxResponseTxResponseLogsItems0EventsItems0 struct {

	// attributes
	Attributes []*BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this broadcast tx response tx response logs items0 events items0
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this broadcast tx response tx response logs items0 events items0 based on the context it is used
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseLogsItems0EventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0 broadcast tx response tx response logs items0 events items0 attributes items0
//
// swagger:model BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0
type BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this broadcast tx response tx response logs items0 events items0 attributes items0
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast tx response tx response logs items0 events items0 attributes items0 based on context it is used
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseLogsItems0EventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastTxResponseTxResponseTx broadcast tx response tx response tx
//
// swagger:model BroadcastTxResponseTxResponseTx
type BroadcastTxResponseTxResponseTx struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// broadcast tx response tx response tx
	BroadcastTxResponseTxResponseTx map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *BroadcastTxResponseTxResponseTx) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv BroadcastTxResponseTxResponseTx

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
		m.BroadcastTxResponseTxResponseTx = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m BroadcastTxResponseTxResponseTx) MarshalJSON() ([]byte, error) {
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

	if len(m.BroadcastTxResponseTxResponseTx) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.BroadcastTxResponseTxResponseTx)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this broadcast tx response tx response tx
func (m *BroadcastTxResponseTxResponseTx) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast tx response tx response tx based on context it is used
func (m *BroadcastTxResponseTxResponseTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastTxResponseTxResponseTx) UnmarshalBinary(b []byte) error {
	var res BroadcastTxResponseTxResponseTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
