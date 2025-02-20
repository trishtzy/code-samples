// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponsesMarketResp responses market resp
//
// swagger:model responses.MarketResp
type ResponsesMarketResp struct {

	// Type of asset
	// Example: PERP
	AssetKind string `json:"asset_kind,omitempty"`

	// Base currency of the market pair
	// Example: ETH
	BaseCurrency string `json:"base_currency,omitempty"`

	// Chain details
	ChainDetails struct {
		ResponsesMarketChainDetails
	} `json:"chain_details,omitempty"`

	// Clamp rate
	// Example: 0.05
	ClampRate string `json:"clamp_rate,omitempty"`

	// Delta1 Cross margin parameters
	Delta1CrossMarginParams struct {
		ResponsesDelta1CrossMarginParams
	} `json:"delta1_cross_margin_params,omitempty"`

	// Market expiry time
	// Example: 0
	ExpiryAt int64 `json:"expiry_at,omitempty"`

	// Funding period in hours
	// Example: 8
	FundingPeriodHours float64 `json:"funding_period_hours,omitempty"`

	// Interest rate
	// Example: 0.01
	InterestRate string `json:"interest_rate,omitempty"`

	// IV Bands Width
	// Example: 0.05
	IvBandsWidth string `json:"iv_bands_width,omitempty"`

	// Type of market - always 'cross'
	// Example: cross
	MarketKind string `json:"market_kind,omitempty"`

	// Max funding rate
	// Example: 0.05
	MaxFundingRate string `json:"max_funding_rate,omitempty"`

	// Max funding rate change
	// Example: 0.0005
	MaxFundingRateChange string `json:"max_funding_rate_change,omitempty"`

	// Max open orders
	// Example: 100
	MaxOpenOrders int64 `json:"max_open_orders,omitempty"`

	// Maximum order size
	// Example: 100
	MaxOrderSize string `json:"max_order_size,omitempty"`

	// The maximum TOB spread allowed to apply funding rate changes
	// Example: 0.2
	MaxTobSpread string `json:"max_tob_spread,omitempty"`

	// Minimum order size in USD
	// Example: 10
	MinNotional string `json:"min_notional,omitempty"`

	// Market open time in milliseconds
	// Example: 0
	OpenAt int64 `json:"open_at,omitempty"`

	// Option Cross margin parameters
	OptionCrossMarginParams struct {
		ResponsesPerpetualOptionCrossMarginParams
	} `json:"option_cross_margin_params,omitempty"`

	// Type of option (PUT or CALL)
	// Example: PUT
	OptionType string `json:"option_type,omitempty"`

	// Oracle EWMA factor
	// Example: 0.2
	OracleEwmaFactor string `json:"oracle_ewma_factor,omitempty"`

	// Minimum size increment for base currency
	// Example: 0.001
	OrderSizeIncrement string `json:"order_size_increment,omitempty"`

	// Position limit
	// Example: 500
	PositionLimit string `json:"position_limit,omitempty"`

	// Price Bands Width, 0.05 means 5% price deviation allowed from mark price
	// Example: 0.05
	PriceBandsWidth string `json:"price_bands_width,omitempty"`

	// Price feed id. Pyth price account used to price underlying asset
	// Example: GVXRSBjFk6e6J3NbVPXohDJetcTjaeeuykUpbQF8UoMU
	PriceFeedID string `json:"price_feed_id,omitempty"`

	// Minimum price increment of the market in USD
	// Example: 0.01
	PriceTickSize string `json:"price_tick_size,omitempty"`

	// Quote currency of the market pair
	// Example: USD
	QuoteCurrency string `json:"quote_currency,omitempty"`

	// Settlement currency of the market pair
	// Example: USDC
	SettlementCurrency string `json:"settlement_currency,omitempty"`

	// Strike price for option market
	// Example: 66500
	StrikePrice string `json:"strike_price,omitempty"`

	// Market symbol
	// Example: ETH-USD-PERP
	Symbol string `json:"symbol,omitempty"`

	// Market tags
	// Example: ["MEME","DEFI"]
	Tags []string `json:"tags"`
}

// Validate validates this responses market resp
func (m *ResponsesMarketResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChainDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelta1CrossMarginParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionCrossMarginParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesMarketResp) validateChainDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ChainDetails) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesMarketResp) validateDelta1CrossMarginParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Delta1CrossMarginParams) { // not required
		return nil
	}

	return nil
}

func (m *ResponsesMarketResp) validateOptionCrossMarginParams(formats strfmt.Registry) error {
	if swag.IsZero(m.OptionCrossMarginParams) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this responses market resp based on the context it is used
func (m *ResponsesMarketResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChainDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelta1CrossMarginParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptionCrossMarginParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesMarketResp) contextValidateChainDetails(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesMarketResp) contextValidateDelta1CrossMarginParams(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ResponsesMarketResp) contextValidateOptionCrossMarginParams(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesMarketResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesMarketResp) UnmarshalBinary(b []byte) error {
	var res ResponsesMarketResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
