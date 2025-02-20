// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponsesTransactionResponse responses transaction response
//
// swagger:model responses.TransactionResponse
type ResponsesTransactionResponse struct {

	// Timestamp from when the transaction was completed
	CompletedAt int64 `json:"completed_at,omitempty"`

	// Timestamp from when the transaction was sent to blockchain gateway
	CreatedAt int64 `json:"created_at,omitempty"`

	// Tx Hash of the settled trade                                                // Hash of the transaction
	// Example: 0x445c05d6bfb899e39338440d199971c4d7f4cde7878ed3888df3f716efb8df2
	Hash string `json:"hash,omitempty"`

	// Unique string ID of the event that triggered the transaction. For example, fill ID or liquidation ID
	// Example: 12342423
	ID string `json:"id,omitempty"`

	// Status of the transaction on Starknet
	// Enum: ["ACCEPTED_ON_L1","ACCEPTED_ON_L2","NOT_RECEIVED","RECEIVED","REJECTED","REVERTED"]
	State string `json:"state,omitempty"`

	// Event that triggered the transaction
	// Enum: ["TRANSACTION_FILL","TRANSACTION_LIQUIDATE","TRANSACTION_SETTLE_MARKET"]
	Type string `json:"type,omitempty"`
}

// Validate validates this responses transaction response
func (m *ResponsesTransactionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var responsesTransactionResponseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCEPTED_ON_L1","ACCEPTED_ON_L2","NOT_RECEIVED","RECEIVED","REJECTED","REVERTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesTransactionResponseTypeStatePropEnum = append(responsesTransactionResponseTypeStatePropEnum, v)
	}
}

const (

	// ResponsesTransactionResponseStateACCEPTEDONL1 captures enum value "ACCEPTED_ON_L1"
	ResponsesTransactionResponseStateACCEPTEDONL1 string = "ACCEPTED_ON_L1"

	// ResponsesTransactionResponseStateACCEPTEDONL2 captures enum value "ACCEPTED_ON_L2"
	ResponsesTransactionResponseStateACCEPTEDONL2 string = "ACCEPTED_ON_L2"

	// ResponsesTransactionResponseStateNOTRECEIVED captures enum value "NOT_RECEIVED"
	ResponsesTransactionResponseStateNOTRECEIVED string = "NOT_RECEIVED"

	// ResponsesTransactionResponseStateRECEIVED captures enum value "RECEIVED"
	ResponsesTransactionResponseStateRECEIVED string = "RECEIVED"

	// ResponsesTransactionResponseStateREJECTED captures enum value "REJECTED"
	ResponsesTransactionResponseStateREJECTED string = "REJECTED"

	// ResponsesTransactionResponseStateREVERTED captures enum value "REVERTED"
	ResponsesTransactionResponseStateREVERTED string = "REVERTED"
)

// prop value enum
func (m *ResponsesTransactionResponse) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesTransactionResponseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesTransactionResponse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var responsesTransactionResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRANSACTION_FILL","TRANSACTION_LIQUIDATE","TRANSACTION_SETTLE_MARKET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesTransactionResponseTypeTypePropEnum = append(responsesTransactionResponseTypeTypePropEnum, v)
	}
}

const (

	// ResponsesTransactionResponseTypeTRANSACTIONFILL captures enum value "TRANSACTION_FILL"
	ResponsesTransactionResponseTypeTRANSACTIONFILL string = "TRANSACTION_FILL"

	// ResponsesTransactionResponseTypeTRANSACTIONLIQUIDATE captures enum value "TRANSACTION_LIQUIDATE"
	ResponsesTransactionResponseTypeTRANSACTIONLIQUIDATE string = "TRANSACTION_LIQUIDATE"

	// ResponsesTransactionResponseTypeTRANSACTIONSETTLEMARKET captures enum value "TRANSACTION_SETTLE_MARKET"
	ResponsesTransactionResponseTypeTRANSACTIONSETTLEMARKET string = "TRANSACTION_SETTLE_MARKET"
)

// prop value enum
func (m *ResponsesTransactionResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesTransactionResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesTransactionResponse) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this responses transaction response based on context it is used
func (m *ResponsesTransactionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesTransactionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesTransactionResponse) UnmarshalBinary(b []byte) error {
	var res ResponsesTransactionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
