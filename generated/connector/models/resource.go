package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	manifold "github.com/manifoldco/go-manifold"
)

// Resource A view of a Resource provisioned through Manifold.
//
// Do not store any of this data, instead query Manifold for the most up to
// date information.
//
// swagger:model Resource
type Resource struct {

	// created at
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	ID manifold.ID `json:"id,omitempty"`

	// name
	Name manifold.Name `json:"name,omitempty"`

	// plan
	Plan manifold.Label `json:"plan,omitempty"`

	// product
	Product manifold.Label `json:"product,omitempty"`

	// region
	Region RegionSlug `json:"region,omitempty"`

	// updated at
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Resource) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *Resource) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if err := m.Plan.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("plan")
		}
		return err
	}

	return nil
}

func (m *Resource) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if err := m.Product.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("product")
		}
		return err
	}

	return nil
}

func (m *Resource) validateRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := m.Region.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("region")
		}
		return err
	}

	return nil
}
