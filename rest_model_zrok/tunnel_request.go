// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TunnelRequest tunnel request
//
// swagger:model tunnelRequest
type TunnelRequest struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// ziti identity Id
	ZitiIdentityID string `json:"zitiIdentityId,omitempty"`
}

// Validate validates this tunnel request
func (m *TunnelRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tunnel request based on context it is used
func (m *TunnelRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TunnelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TunnelRequest) UnmarshalBinary(b []byte) error {
	var res TunnelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
