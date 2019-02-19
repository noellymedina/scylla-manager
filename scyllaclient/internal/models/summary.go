// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Summary summary
//
// A compaction summary object
// swagger:model summary
type Summary struct {

	// The column family name
	Cf string `json:"cf,omitempty"`

	// The number of units completed
	Completed interface{} `json:"completed,omitempty"`

	// The UUID
	ID string `json:"id,omitempty"`

	// The keyspace name
	Ks string `json:"ks,omitempty"`

	// The task compaction type
	TaskType string `json:"task_type,omitempty"`

	// The total number of units
	Total interface{} `json:"total,omitempty"`

	// The units being used
	Unit string `json:"unit,omitempty"`
}

// Validate validates this summary
func (m *Summary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Summary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Summary) UnmarshalBinary(b []byte) error {
	var res Summary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
