package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CallbackResponse A callback sent from a provider to complete an asynchronous request.
//
// Credentials can only be specified *if* the callback corresponds with a
// credential provisioning request.
//
// swagger:model CallbackResponse
type CallbackResponse struct {

	// credentials
	Credentials map[string]string `json:"credentials,omitempty"`

	// message
	// Required: true
	// Max Length: 256
	// Min Length: 3
	Message *string `json:"message"`

	// state
	// Required: true
	State *string `json:"state"`
}

// Validate validates this callback response
func (m *CallbackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallbackResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", string(*m.Message), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("message", "body", string(*m.Message), 256); err != nil {
		return err
	}

	return nil
}

var callbackResponseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["done","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callbackResponseTypeStatePropEnum = append(callbackResponseTypeStatePropEnum, v)
	}
}

const (
	// CallbackResponseStateDone captures enum value "done"
	CallbackResponseStateDone string = "done"
	// CallbackResponseStateError captures enum value "error"
	CallbackResponseStateError string = "error"
)

// prop value enum
func (m *CallbackResponse) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callbackResponseTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallbackResponse) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}
