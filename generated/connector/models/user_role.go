package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UserRole user role
// swagger:model UserRole
type UserRole string

const (
	// UserRoleOwner captures enum value "owner"
	UserRoleOwner UserRole = "owner"
	// UserRoleMember captures enum value "member"
	UserRoleMember UserRole = "member"
	// UserRoleAdmin captures enum value "admin"
	UserRoleAdmin UserRole = "admin"
)

// for schema
var userRoleEnum []interface{}

func init() {
	var res []UserRole
	if err := json.Unmarshal([]byte(`["owner","member","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRoleEnum = append(userRoleEnum, v)
	}
}

func (m UserRole) validateUserRoleEnum(path, location string, value UserRole) error {
	if err := validate.Enum(path, location, value, userRoleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this user role
func (m UserRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}