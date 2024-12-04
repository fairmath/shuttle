// Code generated by go-swagger; DO NOT EDIT.

package service

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

// ServiceTxEncodeAminoReader is a Reader for the ServiceTxEncodeAmino structure.
type ServiceTxEncodeAminoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceTxEncodeAminoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceTxEncodeAminoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceTxEncodeAminoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceTxEncodeAminoOK creates a ServiceTxEncodeAminoOK with default headers values
func NewServiceTxEncodeAminoOK() *ServiceTxEncodeAminoOK {
	return &ServiceTxEncodeAminoOK{}
}

/*
ServiceTxEncodeAminoOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServiceTxEncodeAminoOK struct {
	Payload *ServiceTxEncodeAminoOKBody
}

// IsSuccess returns true when this service tx encode amino o k response has a 2xx status code
func (o *ServiceTxEncodeAminoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service tx encode amino o k response has a 3xx status code
func (o *ServiceTxEncodeAminoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service tx encode amino o k response has a 4xx status code
func (o *ServiceTxEncodeAminoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service tx encode amino o k response has a 5xx status code
func (o *ServiceTxEncodeAminoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service tx encode amino o k response a status code equal to that given
func (o *ServiceTxEncodeAminoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service tx encode amino o k response
func (o *ServiceTxEncodeAminoOK) Code() int {
	return 200
}

func (o *ServiceTxEncodeAminoOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cosmos/tx/v1beta1/encode/amino][%d] serviceTxEncodeAminoOK %s", 200, payload)
}

func (o *ServiceTxEncodeAminoOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cosmos/tx/v1beta1/encode/amino][%d] serviceTxEncodeAminoOK %s", 200, payload)
}

func (o *ServiceTxEncodeAminoOK) GetPayload() *ServiceTxEncodeAminoOKBody {
	return o.Payload
}

func (o *ServiceTxEncodeAminoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ServiceTxEncodeAminoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceTxEncodeAminoDefault creates a ServiceTxEncodeAminoDefault with default headers values
func NewServiceTxEncodeAminoDefault(code int) *ServiceTxEncodeAminoDefault {
	return &ServiceTxEncodeAminoDefault{
		_statusCode: code,
	}
}

/*
ServiceTxEncodeAminoDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServiceTxEncodeAminoDefault struct {
	_statusCode int

	Payload *ServiceTxEncodeAminoDefaultBody
}

// IsSuccess returns true when this service tx encode amino default response has a 2xx status code
func (o *ServiceTxEncodeAminoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service tx encode amino default response has a 3xx status code
func (o *ServiceTxEncodeAminoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service tx encode amino default response has a 4xx status code
func (o *ServiceTxEncodeAminoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service tx encode amino default response has a 5xx status code
func (o *ServiceTxEncodeAminoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service tx encode amino default response a status code equal to that given
func (o *ServiceTxEncodeAminoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service tx encode amino default response
func (o *ServiceTxEncodeAminoDefault) Code() int {
	return o._statusCode
}

func (o *ServiceTxEncodeAminoDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cosmos/tx/v1beta1/encode/amino][%d] Service_TxEncodeAmino default %s", o._statusCode, payload)
}

func (o *ServiceTxEncodeAminoDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cosmos/tx/v1beta1/encode/amino][%d] Service_TxEncodeAmino default %s", o._statusCode, payload)
}

func (o *ServiceTxEncodeAminoDefault) GetPayload() *ServiceTxEncodeAminoDefaultBody {
	return o.Payload
}

func (o *ServiceTxEncodeAminoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ServiceTxEncodeAminoDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServiceTxEncodeAminoBody service tx encode amino body
swagger:model ServiceTxEncodeAminoBody
*/
type ServiceTxEncodeAminoBody struct {

	// amino json
	AminoJSON string `json:"amino_json,omitempty"`
}

// Validate validates this service tx encode amino body
func (o *ServiceTxEncodeAminoBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service tx encode amino body based on context it is used
func (o *ServiceTxEncodeAminoBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceTxEncodeAminoBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceTxEncodeAminoBody) UnmarshalBinary(b []byte) error {
	var res ServiceTxEncodeAminoBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ServiceTxEncodeAminoDefaultBody service tx encode amino default body
swagger:model ServiceTxEncodeAminoDefaultBody
*/
type ServiceTxEncodeAminoDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ServiceTxEncodeAminoDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this service tx encode amino default body
func (o *ServiceTxEncodeAminoDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServiceTxEncodeAminoDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Service_TxEncodeAmino default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Service_TxEncodeAmino default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service tx encode amino default body based on the context it is used
func (o *ServiceTxEncodeAminoDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServiceTxEncodeAminoDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Service_TxEncodeAmino default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Service_TxEncodeAmino default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ServiceTxEncodeAminoDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceTxEncodeAminoDefaultBody) UnmarshalBinary(b []byte) error {
	var res ServiceTxEncodeAminoDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ServiceTxEncodeAminoDefaultBodyDetailsItems0 service tx encode amino default body details items0
swagger:model ServiceTxEncodeAminoDefaultBodyDetailsItems0
*/
type ServiceTxEncodeAminoDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// service tx encode amino default body details items0
	ServiceTxEncodeAminoDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ServiceTxEncodeAminoDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ServiceTxEncodeAminoDefaultBodyDetailsItems0

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
		o.ServiceTxEncodeAminoDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ServiceTxEncodeAminoDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.ServiceTxEncodeAminoDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ServiceTxEncodeAminoDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this service tx encode amino default body details items0
func (o *ServiceTxEncodeAminoDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service tx encode amino default body details items0 based on context it is used
func (o *ServiceTxEncodeAminoDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceTxEncodeAminoDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceTxEncodeAminoDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceTxEncodeAminoDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ServiceTxEncodeAminoOKBody service tx encode amino o k body
swagger:model ServiceTxEncodeAminoOKBody
*/
type ServiceTxEncodeAminoOKBody struct {

	// amino binary
	// Format: byte
	AminoBinary strfmt.Base64 `json:"amino_binary,omitempty"`
}

// Validate validates this service tx encode amino o k body
func (o *ServiceTxEncodeAminoOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service tx encode amino o k body based on context it is used
func (o *ServiceTxEncodeAminoOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceTxEncodeAminoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceTxEncodeAminoOKBody) UnmarshalBinary(b []byte) error {
	var res ServiceTxEncodeAminoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
