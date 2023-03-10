// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemData item data
//
// swagger:model ItemData
type ItemData struct {

	// box code
	BoxCode string `json:"boxCode,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// props
	Props map[string]string `json:"props,omitempty"`

	// server ID
	ServerID string `json:"serverID,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this item data
func (m *ItemData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item data based on context it is used
func (m *ItemData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemData) UnmarshalBinary(b []byte) error {
	var res ItemData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
