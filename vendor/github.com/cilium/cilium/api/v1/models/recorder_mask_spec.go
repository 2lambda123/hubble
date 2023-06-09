// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecorderMaskSpec Configuration of a recorder mask
//
// swagger:model RecorderMaskSpec
type RecorderMaskSpec struct {

	// Layer 4 destination port mask
	DstPortMask string `json:"dst-port-mask,omitempty"`

	// Layer 3 destination IP mask
	DstPrefixMask string `json:"dst-prefix-mask,omitempty"`

	// Priority of this mask
	Priority int64 `json:"priority,omitempty"`

	// Layer 4 protocol mask
	ProtocolMask string `json:"protocol-mask,omitempty"`

	// Layer 4 source port mask
	SrcPortMask string `json:"src-port-mask,omitempty"`

	// Layer 3 source IP mask
	SrcPrefixMask string `json:"src-prefix-mask,omitempty"`

	// Number of users of this mask
	Users int64 `json:"users,omitempty"`
}

// Validate validates this recorder mask spec
func (m *RecorderMaskSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this recorder mask spec based on context it is used
func (m *RecorderMaskSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecorderMaskSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecorderMaskSpec) UnmarshalBinary(b []byte) error {
	var res RecorderMaskSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
