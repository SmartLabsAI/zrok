// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Service service
//
// swagger:model service
type Service struct {

	// backend proxy endpoint
	BackendProxyEndpoint string `json:"backendProxyEndpoint,omitempty"`

	// created at
	CreatedAt int64 `json:"createdAt,omitempty"`

	// frontend endpoint
	FrontendEndpoint string `json:"frontendEndpoint,omitempty"`

	// metrics
	Metrics ServiceMetrics `json:"metrics,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updatedAt,omitempty"`

	// z Id
	ZID string `json:"zId,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if err := m.Metrics.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metrics")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metrics")
		}
		return err
	}

	return nil
}

// ContextValidate validate this service based on the context it is used
func (m *Service) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metrics")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metrics")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
