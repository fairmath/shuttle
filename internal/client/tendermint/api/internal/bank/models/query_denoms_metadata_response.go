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

// QueryDenomsMetadataResponse query denoms metadata response
//
// swagger:model QueryDenomsMetadataResponse
type QueryDenomsMetadataResponse struct {

	// metadatas
	Metadatas []*QueryDenomsMetadataResponseMetadatasItems0 `json:"metadatas"`

	// pagination
	Pagination *QueryDenomsMetadataResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this query denoms metadata response
func (m *QueryDenomsMetadataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadatas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryDenomsMetadataResponse) validateMetadatas(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadatas) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadatas); i++ {
		if swag.IsZero(m.Metadatas[i]) { // not required
			continue
		}

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadatas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryDenomsMetadataResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this query denoms metadata response based on the context it is used
func (m *QueryDenomsMetadataResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadatas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryDenomsMetadataResponse) contextValidateMetadatas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadatas); i++ {

		if m.Metadatas[i] != nil {

			if swag.IsZero(m.Metadatas[i]) { // not required
				return nil
			}

			if err := m.Metadatas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadatas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryDenomsMetadataResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *QueryDenomsMetadataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryDenomsMetadataResponse) UnmarshalBinary(b []byte) error {
	var res QueryDenomsMetadataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryDenomsMetadataResponseMetadatasItems0 query denoms metadata response metadatas items0
//
// swagger:model QueryDenomsMetadataResponseMetadatasItems0
type QueryDenomsMetadataResponseMetadatasItems0 struct {

	// base
	Base string `json:"base,omitempty"`

	// denom units
	DenomUnits []*QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0 `json:"denom_units"`

	// description
	Description string `json:"description,omitempty"`

	// display
	Display string `json:"display,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`

	// uri hash
	URIHash string `json:"uri_hash,omitempty"`
}

// Validate validates this query denoms metadata response metadatas items0
func (m *QueryDenomsMetadataResponseMetadatasItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDenomUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryDenomsMetadataResponseMetadatasItems0) validateDenomUnits(formats strfmt.Registry) error {
	if swag.IsZero(m.DenomUnits) { // not required
		return nil
	}

	for i := 0; i < len(m.DenomUnits); i++ {
		if swag.IsZero(m.DenomUnits[i]) { // not required
			continue
		}

		if m.DenomUnits[i] != nil {
			if err := m.DenomUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query denoms metadata response metadatas items0 based on the context it is used
func (m *QueryDenomsMetadataResponseMetadatasItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDenomUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryDenomsMetadataResponseMetadatasItems0) contextValidateDenomUnits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DenomUnits); i++ {

		if m.DenomUnits[i] != nil {

			if swag.IsZero(m.DenomUnits[i]) { // not required
				return nil
			}

			if err := m.DenomUnits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryDenomsMetadataResponseMetadatasItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryDenomsMetadataResponseMetadatasItems0) UnmarshalBinary(b []byte) error {
	var res QueryDenomsMetadataResponseMetadatasItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0 query denoms metadata response metadatas items0 denom units items0
//
// swagger:model QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0
type QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0 struct {

	// aliases
	Aliases []string `json:"aliases"`

	// denom
	Denom string `json:"denom,omitempty"`

	// exponent
	Exponent int64 `json:"exponent,omitempty"`
}

// Validate validates this query denoms metadata response metadatas items0 denom units items0
func (m *QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query denoms metadata response metadatas items0 denom units items0 based on context it is used
func (m *QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0) UnmarshalBinary(b []byte) error {
	var res QueryDenomsMetadataResponseMetadatasItems0DenomUnitsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryDenomsMetadataResponsePagination query denoms metadata response pagination
//
// swagger:model QueryDenomsMetadataResponsePagination
type QueryDenomsMetadataResponsePagination struct {

	// next key
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this query denoms metadata response pagination
func (m *QueryDenomsMetadataResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query denoms metadata response pagination based on context it is used
func (m *QueryDenomsMetadataResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryDenomsMetadataResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryDenomsMetadataResponsePagination) UnmarshalBinary(b []byte) error {
	var res QueryDenomsMetadataResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
