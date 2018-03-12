// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFileIDOKBody delete file Id o k body
// swagger:model deleteFileIdOKBody
type DeleteFileIDOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this delete file Id o k body
func (m *DeleteFileIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFileIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFileIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteFileIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}