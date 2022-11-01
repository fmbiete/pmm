// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartMySQLExplainJSONActionReader is a Reader for the StartMySQLExplainJSONAction structure.
type StartMySQLExplainJSONActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLExplainJSONActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMySQLExplainJSONActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMySQLExplainJSONActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLExplainJSONActionOK creates a StartMySQLExplainJSONActionOK with default headers values
func NewStartMySQLExplainJSONActionOK() *StartMySQLExplainJSONActionOK {
	return &StartMySQLExplainJSONActionOK{}
}

/*
StartMySQLExplainJSONActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartMySQLExplainJSONActionOK struct {
	Payload *StartMySQLExplainJSONActionOKBody
}

func (o *StartMySQLExplainJSONActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLExplainJSON][%d] startMySqlExplainJsonActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLExplainJSONActionOK) GetPayload() *StartMySQLExplainJSONActionOKBody {
	return o.Payload
}

func (o *StartMySQLExplainJSONActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartMySQLExplainJSONActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLExplainJSONActionDefault creates a StartMySQLExplainJSONActionDefault with default headers values
func NewStartMySQLExplainJSONActionDefault(code int) *StartMySQLExplainJSONActionDefault {
	return &StartMySQLExplainJSONActionDefault{
		_statusCode: code,
	}
}

/*
StartMySQLExplainJSONActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartMySQLExplainJSONActionDefault struct {
	_statusCode int

	Payload *StartMySQLExplainJSONActionDefaultBody
}

// Code gets the status code for the start my SQL explain JSON action default response
func (o *StartMySQLExplainJSONActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLExplainJSONActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLExplainJSON][%d] StartMySQLExplainJSONAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLExplainJSONActionDefault) GetPayload() *StartMySQLExplainJSONActionDefaultBody {
	return o.Payload
}

func (o *StartMySQLExplainJSONActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartMySQLExplainJSONActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StartMySQLExplainJSONActionBody start my SQL explain JSON action body
swagger:model StartMySQLExplainJSONActionBody
*/
type StartMySQLExplainJSONActionBody struct {
	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Deprecated: should not be used, should be removed.
	Query string `json:"query,omitempty"`

	// Database name. Required if it can't be deduced from the query ID.
	Database string `json:"database,omitempty"`

	// Array of placeholder values
	Placeholders []string `json:"placeholders"`

	// Query ID of query. Query ID or query is required.
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this start my SQL explain JSON action body
func (o *StartMySQLExplainJSONActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain JSON action body based on context it is used
func (o *StartMySQLExplainJSONActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartMySQLExplainJSONActionDefaultBody start my SQL explain JSON action default body
swagger:model StartMySQLExplainJSONActionDefaultBody
*/
type StartMySQLExplainJSONActionDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartMySQLExplainJSONActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start my SQL explain JSON action default body
func (o *StartMySQLExplainJSONActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLExplainJSONActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartMySQLExplainJSONAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLExplainJSONAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start my SQL explain JSON action default body based on the context it is used
func (o *StartMySQLExplainJSONActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLExplainJSONActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartMySQLExplainJSONAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLExplainJSONAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartMySQLExplainJSONActionDefaultBodyDetailsItems0 start my SQL explain JSON action default body details items0
swagger:model StartMySQLExplainJSONActionDefaultBodyDetailsItems0
*/
type StartMySQLExplainJSONActionDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this start my SQL explain JSON action default body details items0
func (o *StartMySQLExplainJSONActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain JSON action default body details items0 based on context it is used
func (o *StartMySQLExplainJSONActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartMySQLExplainJSONActionOKBody start my SQL explain JSON action OK body
swagger:model StartMySQLExplainJSONActionOKBody
*/
type StartMySQLExplainJSONActionOKBody struct {
	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL explain JSON action OK body
func (o *StartMySQLExplainJSONActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain JSON action OK body based on context it is used
func (o *StartMySQLExplainJSONActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
