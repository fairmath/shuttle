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

// QueryTotalSupplyReader is a Reader for the QueryTotalSupply structure.
type QueryTotalSupplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryTotalSupplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryTotalSupplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryTotalSupplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryTotalSupplyOK creates a QueryTotalSupplyOK with default headers values
func NewQueryTotalSupplyOK() *QueryTotalSupplyOK {
	return &QueryTotalSupplyOK{}
}

/*
QueryTotalSupplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryTotalSupplyOK struct {
	Payload *QueryTotalSupplyOKBody
}

// IsSuccess returns true when this query total supply o k response has a 2xx status code
func (o *QueryTotalSupplyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query total supply o k response has a 3xx status code
func (o *QueryTotalSupplyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query total supply o k response has a 4xx status code
func (o *QueryTotalSupplyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query total supply o k response has a 5xx status code
func (o *QueryTotalSupplyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query total supply o k response a status code equal to that given
func (o *QueryTotalSupplyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query total supply o k response
func (o *QueryTotalSupplyOK) Code() int {
	return 200
}

func (o *QueryTotalSupplyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply][%d] queryTotalSupplyOK %s", 200, payload)
}

func (o *QueryTotalSupplyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply][%d] queryTotalSupplyOK %s", 200, payload)
}

func (o *QueryTotalSupplyOK) GetPayload() *QueryTotalSupplyOKBody {
	return o.Payload
}

func (o *QueryTotalSupplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryTotalSupplyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryTotalSupplyDefault creates a QueryTotalSupplyDefault with default headers values
func NewQueryTotalSupplyDefault(code int) *QueryTotalSupplyDefault {
	return &QueryTotalSupplyDefault{
		_statusCode: code,
	}
}

/*
QueryTotalSupplyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type QueryTotalSupplyDefault struct {
	_statusCode int

	Payload *QueryTotalSupplyDefaultBody
}

// IsSuccess returns true when this query total supply default response has a 2xx status code
func (o *QueryTotalSupplyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query total supply default response has a 3xx status code
func (o *QueryTotalSupplyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query total supply default response has a 4xx status code
func (o *QueryTotalSupplyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query total supply default response has a 5xx status code
func (o *QueryTotalSupplyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query total supply default response a status code equal to that given
func (o *QueryTotalSupplyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query total supply default response
func (o *QueryTotalSupplyDefault) Code() int {
	return o._statusCode
}

func (o *QueryTotalSupplyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply][%d] Query_TotalSupply default %s", o._statusCode, payload)
}

func (o *QueryTotalSupplyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply][%d] Query_TotalSupply default %s", o._statusCode, payload)
}

func (o *QueryTotalSupplyDefault) GetPayload() *QueryTotalSupplyDefaultBody {
	return o.Payload
}

func (o *QueryTotalSupplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryTotalSupplyDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
QueryTotalSupplyDefaultBody query total supply default body
swagger:model QueryTotalSupplyDefaultBody
*/
type QueryTotalSupplyDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*QueryTotalSupplyDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this query total supply default body
func (o *QueryTotalSupplyDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryTotalSupplyDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Query_TotalSupply default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Query_TotalSupply default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query total supply default body based on the context it is used
func (o *QueryTotalSupplyDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryTotalSupplyDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Query_TotalSupply default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Query_TotalSupply default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryTotalSupplyDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryTotalSupplyDefaultBody) UnmarshalBinary(b []byte) error {
	var res QueryTotalSupplyDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryTotalSupplyDefaultBodyDetailsItems0 query total supply default body details items0
swagger:model QueryTotalSupplyDefaultBodyDetailsItems0
*/
type QueryTotalSupplyDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// query total supply default body details items0
	QueryTotalSupplyDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *QueryTotalSupplyDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv QueryTotalSupplyDefaultBodyDetailsItems0

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
		o.QueryTotalSupplyDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o QueryTotalSupplyDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.QueryTotalSupplyDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.QueryTotalSupplyDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this query total supply default body details items0
func (o *QueryTotalSupplyDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query total supply default body details items0 based on context it is used
func (o *QueryTotalSupplyDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryTotalSupplyDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryTotalSupplyDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res QueryTotalSupplyDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryTotalSupplyOKBody query total supply o k body
swagger:model QueryTotalSupplyOKBody
*/
type QueryTotalSupplyOKBody struct {

	// pagination
	Pagination *QueryTotalSupplyOKBodyPagination `json:"pagination,omitempty"`

	// supply
	Supply []*QueryTotalSupplyOKBodySupplyItems0 `json:"supply"`
}

// Validate validates this query total supply o k body
func (o *QueryTotalSupplyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryTotalSupplyOKBody) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryTotalSupplyOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryTotalSupplyOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

func (o *QueryTotalSupplyOKBody) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(o.Supply) { // not required
		return nil
	}

	for i := 0; i < len(o.Supply); i++ {
		if swag.IsZero(o.Supply[i]) { // not required
			continue
		}

		if o.Supply[i] != nil {
			if err := o.Supply[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryTotalSupplyOK" + "." + "supply" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queryTotalSupplyOK" + "." + "supply" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query total supply o k body based on the context it is used
func (o *QueryTotalSupplyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSupply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryTotalSupplyOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if swag.IsZero(o.Pagination) { // not required
			return nil
		}

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryTotalSupplyOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryTotalSupplyOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

func (o *QueryTotalSupplyOKBody) contextValidateSupply(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Supply); i++ {

		if o.Supply[i] != nil {

			if swag.IsZero(o.Supply[i]) { // not required
				return nil
			}

			if err := o.Supply[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryTotalSupplyOK" + "." + "supply" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queryTotalSupplyOK" + "." + "supply" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryTotalSupplyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryTotalSupplyOKBody) UnmarshalBinary(b []byte) error {
	var res QueryTotalSupplyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryTotalSupplyOKBodyPagination query total supply o k body pagination
swagger:model QueryTotalSupplyOKBodyPagination
*/
type QueryTotalSupplyOKBodyPagination struct {

	// next key
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this query total supply o k body pagination
func (o *QueryTotalSupplyOKBodyPagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query total supply o k body pagination based on context it is used
func (o *QueryTotalSupplyOKBodyPagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryTotalSupplyOKBodyPagination) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryTotalSupplyOKBodyPagination) UnmarshalBinary(b []byte) error {
	var res QueryTotalSupplyOKBodyPagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
QueryTotalSupplyOKBodySupplyItems0 query total supply o k body supply items0
swagger:model QueryTotalSupplyOKBodySupplyItems0
*/
type QueryTotalSupplyOKBodySupplyItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this query total supply o k body supply items0
func (o *QueryTotalSupplyOKBodySupplyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query total supply o k body supply items0 based on context it is used
func (o *QueryTotalSupplyOKBodySupplyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryTotalSupplyOKBodySupplyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryTotalSupplyOKBodySupplyItems0) UnmarshalBinary(b []byte) error {
	var res QueryTotalSupplyOKBodySupplyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
