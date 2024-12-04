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
	"github.com/go-openapi/validate"
)

// ModeInfoSingle mode info single
//
// swagger:model ModeInfo.Single
type ModeInfoSingle struct {

	// mode
	// Enum: ["SIGN_MODE_UNSPECIFIED","SIGN_MODE_DIRECT","SIGN_MODE_TEXTUAL","SIGN_MODE_DIRECT_AUX","SIGN_MODE_LEGACY_AMINO_JSON","SIGN_MODE_EIP_191"]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this mode info single
func (m *ModeInfoSingle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modeInfoSingleTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SIGN_MODE_UNSPECIFIED","SIGN_MODE_DIRECT","SIGN_MODE_TEXTUAL","SIGN_MODE_DIRECT_AUX","SIGN_MODE_LEGACY_AMINO_JSON","SIGN_MODE_EIP_191"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modeInfoSingleTypeModePropEnum = append(modeInfoSingleTypeModePropEnum, v)
	}
}

const (

	// ModeInfoSingleModeSIGNMODEUNSPECIFIED captures enum value "SIGN_MODE_UNSPECIFIED"
	ModeInfoSingleModeSIGNMODEUNSPECIFIED string = "SIGN_MODE_UNSPECIFIED"

	// ModeInfoSingleModeSIGNMODEDIRECT captures enum value "SIGN_MODE_DIRECT"
	ModeInfoSingleModeSIGNMODEDIRECT string = "SIGN_MODE_DIRECT"

	// ModeInfoSingleModeSIGNMODETEXTUAL captures enum value "SIGN_MODE_TEXTUAL"
	ModeInfoSingleModeSIGNMODETEXTUAL string = "SIGN_MODE_TEXTUAL"

	// ModeInfoSingleModeSIGNMODEDIRECTAUX captures enum value "SIGN_MODE_DIRECT_AUX"
	ModeInfoSingleModeSIGNMODEDIRECTAUX string = "SIGN_MODE_DIRECT_AUX"

	// ModeInfoSingleModeSIGNMODELEGACYAMINOJSON captures enum value "SIGN_MODE_LEGACY_AMINO_JSON"
	ModeInfoSingleModeSIGNMODELEGACYAMINOJSON string = "SIGN_MODE_LEGACY_AMINO_JSON"

	// ModeInfoSingleModeSIGNMODEEIP191 captures enum value "SIGN_MODE_EIP_191"
	ModeInfoSingleModeSIGNMODEEIP191 string = "SIGN_MODE_EIP_191"
)

// prop value enum
func (m *ModeInfoSingle) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modeInfoSingleTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModeInfoSingle) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mode info single based on context it is used
func (m *ModeInfoSingle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModeInfoSingle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModeInfoSingle) UnmarshalBinary(b []byte) error {
	var res ModeInfoSingle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
