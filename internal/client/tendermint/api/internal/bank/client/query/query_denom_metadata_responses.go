// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryDenomMetadataReader is a Reader for the QueryDenomMetadata structure.
type QueryDenomMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDenomMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDenomMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryDenomMetadataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryDenomMetadataOK creates a QueryDenomMetadataOK with default headers values
func NewQueryDenomMetadataOK() *QueryDenomMetadataOK {
	return &QueryDenomMetadataOK{}
}

/*
QueryDenomMetadataOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryDenomMetadataOK struct {
	Payload *QueryDenomMetadataOKBody
}

// IsSuccess returns true when this query denom metadata o k response has a 2xx status code
func (o *QueryDenomMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query denom metadata o k response has a 3xx status code
func (o *QueryDenomMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query denom metadata o k response has a 4xx status code
func (o *QueryDenomMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query denom metadata o k response has a 5xx status code
func (o *QueryDenomMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query denom metadata o k response a status code equal to that given
func (o *QueryDenomMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query denom metadata o k response
func (o *QueryDenomMetadataOK) Code() int {
	return 200
}

func (o *QueryDenomMetadataOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/denoms_metadata/{denom}][%d] queryDenomMetadataOK %s", 200, payload)
}

func (o *QueryDenomMetadataOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/denoms_metadata/{denom}][%d] queryDenomMetadataOK %s", 200, payload)
}

func (o *QueryDenomMetadataOK) GetPayload() *QueryDenomMetadataOKBody {
	return o.Payload
}

func (o *QueryDenomMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryDenomMetadataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDenomMetadataDefault creates a QueryDenomMetadataDefault with default headers values
func NewQueryDenomMetadataDefault(code int) *QueryDenomMetadataDefault {
	return &QueryDenomMetadataDefault{
		_statusCode: code,
	}
}

/*
QueryDenomMetadataDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type QueryDenomMetadataDefault struct {
	_statusCode int

	Payload *QueryDenomMetadataDefaultBody
}

// IsSuccess returns true when this query denom metadata default response has a 2xx status code
func (o *QueryDenomMetadataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query denom metadata default response has a 3xx status code
func (o *QueryDenomMetadataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query denom metadata default response has a 4xx status code
func (o *QueryDenomMetadataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query denom metadata default response has a 5xx status code
func (o *QueryDenomMetadataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query denom metadata default response a status code equal to that given
func (o *QueryDenomMetadataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query denom metadata default response
func (o *QueryDenomMetadataDefault) Code() int {
	return o._statusCode
}

func (o *QueryDenomMetadataDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/denoms_metadata/{denom}][%d] Query_DenomMetadata default %s", o._statusCode, payload)
}

func (o *QueryDenomMetadataDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/denoms_metadata/{denom}][%d] Query_DenomMetadata default %s", o._statusCode, payload)
}

func (o *QueryDenomMetadataDefault) GetPayload() *QueryDenomMetadataDefaultBody {
	return o.Payload
}

func (o *QueryDenomMetadataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryDenomMetadataDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
QueryDenomMetadataDefaultBody query denom metadata default body
swagger:model QueryDenomMetadataDefaultBody
*/
type QueryDenomMetadataDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*QueryDenomMetadataDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this query denom metadata default body
func (o *QueryDenomMetadataDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Query_DenomMetadata default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Query_DenomMetadata default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query denom metadata default body based on the context it is used
func (o *QueryDenomMetadataDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Query_DenomMetadata default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Query_DenomMetadata default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryDenomMetadataDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDenomMetadataDefaultBody) UnmarshalBinary(b []byte) error {
	var res QueryDenomMetadataDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryDenomMetadataDefaultBodyDetailsItems0 query denom metadata default body details items0
swagger:model QueryDenomMetadataDefaultBodyDetailsItems0
*/
type QueryDenomMetadataDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// query denom metadata default body details items0
	QueryDenomMetadataDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *QueryDenomMetadataDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv QueryDenomMetadataDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

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
		o.QueryDenomMetadataDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o QueryDenomMetadataDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.QueryDenomMetadataDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.QueryDenomMetadataDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this query denom metadata default body details items0
func (o *QueryDenomMetadataDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query denom metadata default body details items0 based on context it is used
func (o *QueryDenomMetadataDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryDenomMetadataDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDenomMetadataDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res QueryDenomMetadataDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryDenomMetadataOKBody query denom metadata o k body
swagger:model QueryDenomMetadataOKBody
*/
type QueryDenomMetadataOKBody struct {

	// metadata
	Metadata *QueryDenomMetadataOKBodyMetadata `json:"metadata,omitempty"`
}

// Validate validates this query denom metadata o k body
func (o *QueryDenomMetadataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataOKBody) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(o.Metadata) { // not required
		return nil
	}

	if o.Metadata != nil {
		if err := o.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryDenomMetadataOK" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryDenomMetadataOK" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this query denom metadata o k body based on the context it is used
func (o *QueryDenomMetadataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataOKBody) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if o.Metadata != nil {

		if swag.IsZero(o.Metadata) { // not required
			return nil
		}

		if err := o.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryDenomMetadataOK" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryDenomMetadataOK" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryDenomMetadataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDenomMetadataOKBody) UnmarshalBinary(b []byte) error {
	var res QueryDenomMetadataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryDenomMetadataOKBodyMetadata query denom metadata o k body metadata
swagger:model QueryDenomMetadataOKBodyMetadata
*/
type QueryDenomMetadataOKBodyMetadata struct {

	// base
	Base string `json:"base,omitempty"`

	// denom units
	DenomUnits []*QueryDenomMetadataOKBodyMetadataDenomUnitsItems0 `json:"denom_units"`

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

// Validate validates this query denom metadata o k body metadata
func (o *QueryDenomMetadataOKBodyMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDenomUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataOKBodyMetadata) validateDenomUnits(formats strfmt.Registry) error {
	if swag.IsZero(o.DenomUnits) { // not required
		return nil
	}

	for i := 0; i < len(o.DenomUnits); i++ {
		if swag.IsZero(o.DenomUnits[i]) { // not required
			continue
		}

		if o.DenomUnits[i] != nil {
			if err := o.DenomUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryDenomMetadataOK" + "." + "metadata" + "." + "denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queryDenomMetadataOK" + "." + "metadata" + "." + "denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query denom metadata o k body metadata based on the context it is used
func (o *QueryDenomMetadataOKBodyMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDenomUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDenomMetadataOKBodyMetadata) contextValidateDenomUnits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DenomUnits); i++ {

		if o.DenomUnits[i] != nil {

			if swag.IsZero(o.DenomUnits[i]) { // not required
				return nil
			}

			if err := o.DenomUnits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryDenomMetadataOK" + "." + "metadata" + "." + "denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queryDenomMetadataOK" + "." + "metadata" + "." + "denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryDenomMetadataOKBodyMetadata) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDenomMetadataOKBodyMetadata) UnmarshalBinary(b []byte) error {
	var res QueryDenomMetadataOKBodyMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryDenomMetadataOKBodyMetadataDenomUnitsItems0 query denom metadata o k body metadata denom units items0
swagger:model QueryDenomMetadataOKBodyMetadataDenomUnitsItems0
*/
type QueryDenomMetadataOKBodyMetadataDenomUnitsItems0 struct {

	// aliases
	Aliases []string `json:"aliases"`

	// denom
	Denom string `json:"denom,omitempty"`

	// exponent
	Exponent int64 `json:"exponent,omitempty"`
}

// Validate validates this query denom metadata o k body metadata denom units items0
func (o *QueryDenomMetadataOKBodyMetadataDenomUnitsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query denom metadata o k body metadata denom units items0 based on context it is used
func (o *QueryDenomMetadataOKBodyMetadataDenomUnitsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryDenomMetadataOKBodyMetadataDenomUnitsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDenomMetadataOKBodyMetadataDenomUnitsItems0) UnmarshalBinary(b []byte) error {
	var res QueryDenomMetadataOKBodyMetadataDenomUnitsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
