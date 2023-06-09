// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Identity Security identity
//
// swagger:model Identity
type Identity struct {

	// Unique identifier
	ID int64 `json:"id,omitempty"`

	// Labels describing the identity
	Labels Labels `json:"labels,omitempty"`

	// SHA256 of labels
	LabelsSHA256 string `json:"labelsSHA256,omitempty"`
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identity) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

// ContextValidate validate this identity based on the context it is used
func (m *Identity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identity) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
