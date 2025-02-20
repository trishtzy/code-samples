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

// ResponsesOrderResp responses order resp
//
// swagger:model responses.OrderResp
type ResponsesOrderResp struct {

	// Account identifier (user's account address)
	// Example: 0x4638e3041366aa71720be63e32e53e1223316c7f0d56f7aa617542ed1e7512x
	Account string `json:"account,omitempty"`

	// Average fill price of the order
	// Example: 26000
	AvgFillPrice string `json:"avg_fill_price,omitempty"`

	// Reason for order cancellation if it was closed by cancel
	// Example: NOT_ENOUGH_MARGIN
	CancelReason string `json:"cancel_reason,omitempty"`

	// Client id passed on order creation
	// Example: x1234
	ClientID string `json:"client_id,omitempty"`

	// Order creation time
	// Example: 1681493746016
	CreatedAt int64 `json:"created_at,omitempty"`

	// Order flags, allow flag: REDUCE_ONLY
	Flags []ResponsesOrderFlag `json:"flags"`

	// Unique order identifier
	// Example: 123456
	ID string `json:"id,omitempty"`

	// OrderInstruction (GTC, IOC, POST_ONLY)
	// Example: GTC
	Instruction struct {
		ResponsesOrderInstruction
	} `json:"instruction,omitempty"`

	// Order last update time.  No changes once status=CLOSED
	// Example: 1681493746016
	LastUpdatedAt int64 `json:"last_updated_at,omitempty"`

	// Market to which order belongs
	// Example: BTC-USD-PERP
	Market string `json:"market,omitempty"`

	// Order price. 0 for MARKET orders
	// Example: 26000
	Price string `json:"price,omitempty"`

	// Order published to the client time
	// Example: 1681493746016
	PublishedAt int64 `json:"published_at,omitempty"`

	// Order received from the client time
	// Example: 1681493746016
	ReceivedAt int64 `json:"received_at,omitempty"`

	// Remaining size of the order
	// Example: 0
	RemainingSize string `json:"remaining_size,omitempty"`

	// Request info for modify order
	RequestInfo struct {
		ResponsesRequestInfo
	} `json:"request_info,omitempty"`

	// Unique increasing number (non-sequential) that is assigned to this order update and changes on every order update. Can be used to deduplicate multiple feeds. WebSocket and REST responses use independently generated seq_no per event.
	// Example: 1681471234972000000
	SeqNo int64 `json:"seq_no,omitempty"`

	// Order side
	Side struct {
		ResponsesOrderSide
	} `json:"side,omitempty"`

	// Order size
	// Example: 0.05
	Size string `json:"size,omitempty"`

	// Order status
	Status struct {
		ResponsesOrderStatus
	} `json:"status,omitempty"`

	// Self Trade Prevention mode (EXEPIRE_MAKER, EXPIRE_TAKER, EXPIRE_BOTH)
	// Example: EXPIRE_MAKER
	Stp struct {
		ResponsesSTPMode
	} `json:"stp,omitempty"`

	// Order signature timestamp
	// Example: 1681493746016
	Timestamp int64 `json:"timestamp,omitempty"`

	// Trigger price for stop order
	// Example: 26000
	TriggerPrice string `json:"trigger_price,omitempty"`

	// Order type
	Type struct {
		ResponsesOrderType
	} `json:"type,omitempty"`
}

// Validate validates this responses order resp
func (m *ResponsesOrderResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStp(formats); err != nil {
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

func (m *ResponsesOrderResp) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	for i := 0; i < len(m.Flags); i++ {

		if err := m.Flags[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ResponsesOrderResp) validateInstruction(formats strfmt.Registry) error {
	if swag.IsZero(m.Instruction) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesOrderResp) validateRequestInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInfo) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesOrderResp) validateSide(formats strfmt.Registry) error {
	if swag.IsZero(m.Side) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesOrderResp) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesOrderResp) validateStp(formats strfmt.Registry) error {
	if swag.IsZero(m.Stp) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesOrderResp) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this responses order resp based on the context it is used
func (m *ResponsesOrderResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstruction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSide(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesOrderResp) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flags); i++ {

		if swag.IsZero(m.Flags[i]) { // not required
			return nil
		}

		if err := m.Flags[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ResponsesOrderResp) contextValidateInstruction(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesOrderResp) contextValidateRequestInfo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesOrderResp) contextValidateSide(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesOrderResp) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesOrderResp) contextValidateStp(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesOrderResp) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesOrderResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesOrderResp) UnmarshalBinary(b []byte) error {
	var res ResponsesOrderResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
