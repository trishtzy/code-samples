// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponsesAuthResp responses auth resp
//
// swagger:model responses.AuthResp
type ResponsesAuthResp struct {

	// Authentication token
	// Example: eyJhbGciOiJFUzM4NCIsInR5cCI6IkpXVCJ9.eyJ0eXAiOiJhdCtKV1QiLCJleHAiOjE2ODE0NTI5MDcsImlhdCI6MTY4MTQ1MjYwNywiaXNzIjoiUGFyYWRleCBzdGFnaW5nIiwic3ViIjoiMHg0OTVkMmViNTIzNmExMmI4YjRhZDdkMzg0OWNlNmEyMDNjZTIxYzQzZjQ3M2MyNDhkZmQ1Y2U3MGQ5NDU0ZmEifQ.BPihIbGhnnsuPlReqC9x12JFXldpswg5EdA6tTiDQm-_UHaRz_8RfVBqWc2fPN6CzFsXTq7GowZu-2qMxPvZK_fGcxEhTp2k1r8MUxowlUIT4vPu2scCwrsyIujlCAwS
	JwtToken string `json:"jwt_token,omitempty"`
}

// Validate validates this responses auth resp
func (m *ResponsesAuthResp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this responses auth resp based on context it is used
func (m *ResponsesAuthResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesAuthResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesAuthResp) UnmarshalBinary(b []byte) error {
	var res ResponsesAuthResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
