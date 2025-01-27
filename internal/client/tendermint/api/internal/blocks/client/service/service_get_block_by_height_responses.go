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

	"github.com/fairmath/shuttle/internal/client/tendermint/api/internal/blocks/models"
)

// ServiceGetBlockByHeightReader is a Reader for the ServiceGetBlockByHeight structure.
type ServiceGetBlockByHeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceGetBlockByHeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceGetBlockByHeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewServiceGetBlockByHeightNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewServiceGetBlockByHeightDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceGetBlockByHeightOK creates a ServiceGetBlockByHeightOK with default headers values
func NewServiceGetBlockByHeightOK() *ServiceGetBlockByHeightOK {
	return &ServiceGetBlockByHeightOK{}
}

/*
ServiceGetBlockByHeightOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServiceGetBlockByHeightOK struct {
	Payload *models.CosmosBlock
}

// IsSuccess returns true when this service get block by height o k response has a 2xx status code
func (o *ServiceGetBlockByHeightOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service get block by height o k response has a 3xx status code
func (o *ServiceGetBlockByHeightOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service get block by height o k response has a 4xx status code
func (o *ServiceGetBlockByHeightOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service get block by height o k response has a 5xx status code
func (o *ServiceGetBlockByHeightOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service get block by height o k response a status code equal to that given
func (o *ServiceGetBlockByHeightOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service get block by height o k response
func (o *ServiceGetBlockByHeightOK) Code() int {
	return 200
}

func (o *ServiceGetBlockByHeightOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] serviceGetBlockByHeightOK %s", 200, payload)
}

func (o *ServiceGetBlockByHeightOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] serviceGetBlockByHeightOK %s", 200, payload)
}

func (o *ServiceGetBlockByHeightOK) GetPayload() *models.CosmosBlock {
	return o.Payload
}

func (o *ServiceGetBlockByHeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CosmosBlock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceGetBlockByHeightNotFound creates a ServiceGetBlockByHeightNotFound with default headers values
func NewServiceGetBlockByHeightNotFound() *ServiceGetBlockByHeightNotFound {
	return &ServiceGetBlockByHeightNotFound{}
}

/*
ServiceGetBlockByHeightNotFound describes a response with status code 404, with default header values.

Block not found
*/
type ServiceGetBlockByHeightNotFound struct {
}

// IsSuccess returns true when this service get block by height not found response has a 2xx status code
func (o *ServiceGetBlockByHeightNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service get block by height not found response has a 3xx status code
func (o *ServiceGetBlockByHeightNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service get block by height not found response has a 4xx status code
func (o *ServiceGetBlockByHeightNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service get block by height not found response has a 5xx status code
func (o *ServiceGetBlockByHeightNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service get block by height not found response a status code equal to that given
func (o *ServiceGetBlockByHeightNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service get block by height not found response
func (o *ServiceGetBlockByHeightNotFound) Code() int {
	return 404
}

func (o *ServiceGetBlockByHeightNotFound) Error() string {
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] serviceGetBlockByHeightNotFound", 404)
}

func (o *ServiceGetBlockByHeightNotFound) String() string {
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] serviceGetBlockByHeightNotFound", 404)
}

func (o *ServiceGetBlockByHeightNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceGetBlockByHeightDefault creates a ServiceGetBlockByHeightDefault with default headers values
func NewServiceGetBlockByHeightDefault(code int) *ServiceGetBlockByHeightDefault {
	return &ServiceGetBlockByHeightDefault{
		_statusCode: code,
	}
}

/*
ServiceGetBlockByHeightDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServiceGetBlockByHeightDefault struct {
	_statusCode int

	Payload *ServiceGetBlockByHeightDefaultBody
}

// IsSuccess returns true when this service get block by height default response has a 2xx status code
func (o *ServiceGetBlockByHeightDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service get block by height default response has a 3xx status code
func (o *ServiceGetBlockByHeightDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service get block by height default response has a 4xx status code
func (o *ServiceGetBlockByHeightDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service get block by height default response has a 5xx status code
func (o *ServiceGetBlockByHeightDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service get block by height default response a status code equal to that given
func (o *ServiceGetBlockByHeightDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service get block by height default response
func (o *ServiceGetBlockByHeightDefault) Code() int {
	return o._statusCode
}

func (o *ServiceGetBlockByHeightDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] Service_GetBlockByHeight default %s", o._statusCode, payload)
}

func (o *ServiceGetBlockByHeightDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cosmos/base/tendermint/v1beta1/blocks/{height}][%d] Service_GetBlockByHeight default %s", o._statusCode, payload)
}

func (o *ServiceGetBlockByHeightDefault) GetPayload() *ServiceGetBlockByHeightDefaultBody {
	return o.Payload
}

func (o *ServiceGetBlockByHeightDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ServiceGetBlockByHeightDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServiceGetBlockByHeightDefaultBody service get block by height default body
swagger:model ServiceGetBlockByHeightDefaultBody
*/
type ServiceGetBlockByHeightDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ServiceGetBlockByHeightDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this service get block by height default body
func (o *ServiceGetBlockByHeightDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServiceGetBlockByHeightDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Service_GetBlockByHeight default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Service_GetBlockByHeight default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service get block by height default body based on the context it is used
func (o *ServiceGetBlockByHeightDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServiceGetBlockByHeightDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Service_GetBlockByHeight default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Service_GetBlockByHeight default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ServiceGetBlockByHeightDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceGetBlockByHeightDefaultBody) UnmarshalBinary(b []byte) error {
	var res ServiceGetBlockByHeightDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ServiceGetBlockByHeightDefaultBodyDetailsItems0 service get block by height default body details items0
swagger:model ServiceGetBlockByHeightDefaultBodyDetailsItems0
*/
type ServiceGetBlockByHeightDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`

	// service get block by height default body details items0
	ServiceGetBlockByHeightDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ServiceGetBlockByHeightDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ServiceGetBlockByHeightDefaultBodyDetailsItems0

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
		o.ServiceGetBlockByHeightDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ServiceGetBlockByHeightDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.ServiceGetBlockByHeightDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ServiceGetBlockByHeightDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this service get block by height default body details items0
func (o *ServiceGetBlockByHeightDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service get block by height default body details items0 based on context it is used
func (o *ServiceGetBlockByHeightDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceGetBlockByHeightDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceGetBlockByHeightDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceGetBlockByHeightDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
