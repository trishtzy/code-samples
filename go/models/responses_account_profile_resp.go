// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponsesAccountProfileResp responses account profile resp
//
// swagger:model responses.AccountProfileResp
type ResponsesAccountProfileResp struct {

	// discord
	Discord *MessageDiscordProfile `json:"discord,omitempty"`

	// is username private
	// Example: true
	IsUsernamePrivate bool `json:"is_username_private,omitempty"`

	// max slippage
	// Example: 0.05
	MaxSlippage string `json:"max_slippage,omitempty"`

	// nfts
	Nfts []*MessageOwnedNft `json:"nfts"`

	// programs eligibility
	ProgramsEligibility *ResponsesProgramsEligibility `json:"programs_eligibility,omitempty"`

	// referral
	Referral *ResponsesReferralConfigResp `json:"referral,omitempty"`

	// referral code
	// Example: cryptofox8
	ReferralCode string `json:"referral_code,omitempty"`

	// referred by
	// Example: maxDegen
	ReferredBy string `json:"referred_by,omitempty"`

	// twitter
	Twitter *MessageTwitterProfile `json:"twitter,omitempty"`

	// username
	// Example: username
	Username string `json:"username,omitempty"`
}

// Validate validates this responses account profile resp
func (m *ResponsesAccountProfileResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgramsEligibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferral(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesAccountProfileResp) validateDiscord(formats strfmt.Registry) error {
	if swag.IsZero(m.Discord) { // not required
		return nil
	}

	if m.Discord != nil {
		if err := m.Discord.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discord")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discord")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) validateNfts(formats strfmt.Registry) error {
	if swag.IsZero(m.Nfts) { // not required
		return nil
	}

	for i := 0; i < len(m.Nfts); i++ {
		if swag.IsZero(m.Nfts[i]) { // not required
			continue
		}

		if m.Nfts[i] != nil {
			if err := m.Nfts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesAccountProfileResp) validateProgramsEligibility(formats strfmt.Registry) error {
	if swag.IsZero(m.ProgramsEligibility) { // not required
		return nil
	}

	if m.ProgramsEligibility != nil {
		if err := m.ProgramsEligibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("programs_eligibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("programs_eligibility")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) validateReferral(formats strfmt.Registry) error {
	if swag.IsZero(m.Referral) { // not required
		return nil
	}

	if m.Referral != nil {
		if err := m.Referral.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referral")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referral")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) validateTwitter(formats strfmt.Registry) error {
	if swag.IsZero(m.Twitter) { // not required
		return nil
	}

	if m.Twitter != nil {
		if err := m.Twitter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twitter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twitter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this responses account profile resp based on the context it is used
func (m *ResponsesAccountProfileResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscord(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgramsEligibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferral(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTwitter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesAccountProfileResp) contextValidateDiscord(ctx context.Context, formats strfmt.Registry) error {

	if m.Discord != nil {

		if swag.IsZero(m.Discord) { // not required
			return nil
		}

		if err := m.Discord.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discord")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discord")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) contextValidateNfts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nfts); i++ {

		if m.Nfts[i] != nil {

			if swag.IsZero(m.Nfts[i]) { // not required
				return nil
			}

			if err := m.Nfts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesAccountProfileResp) contextValidateProgramsEligibility(ctx context.Context, formats strfmt.Registry) error {

	if m.ProgramsEligibility != nil {

		if swag.IsZero(m.ProgramsEligibility) { // not required
			return nil
		}

		if err := m.ProgramsEligibility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("programs_eligibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("programs_eligibility")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) contextValidateReferral(ctx context.Context, formats strfmt.Registry) error {

	if m.Referral != nil {

		if swag.IsZero(m.Referral) { // not required
			return nil
		}

		if err := m.Referral.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referral")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referral")
			}
			return err
		}
	}

	return nil
}

func (m *ResponsesAccountProfileResp) contextValidateTwitter(ctx context.Context, formats strfmt.Registry) error {

	if m.Twitter != nil {

		if swag.IsZero(m.Twitter) { // not required
			return nil
		}

		if err := m.Twitter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twitter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twitter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesAccountProfileResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesAccountProfileResp) UnmarshalBinary(b []byte) error {
	var res ResponsesAccountProfileResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
