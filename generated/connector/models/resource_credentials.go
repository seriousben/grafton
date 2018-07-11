package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceCredentials resource credentials
// swagger:model ResourceCredentials
type ResourceCredentials struct {

	// created on
	// Required: true
	CreatedOn *strfmt.DateTime `json:"created_on"`

	// Map of configuration variable aliases to original names
	CustomNames map[string]string `json:"custom_names,omitempty"`

	// Map of configuration variable names to values, names
	// must IEEE 1003.1 - 2001 Standard (checked in code).
	//
	// Required: true
	Keys map[string]string `json:"keys"`
}

// Validate validates this resource credentials
func (m *ResourceCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceCredentials) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	return nil
}

func (m *ResourceCredentials) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	return nil
}
