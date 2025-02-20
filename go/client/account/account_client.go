// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new account API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new account API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFills(params *GetFillsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFillsOK, error)

	GetAccount(params *GetAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountOK, error)

	GetAccountMargin(params *GetAccountMarginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountMarginOK, error)

	GetAccountProfile(params *GetAccountProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountProfileOK, error)

	GetAccountsInfo(params *GetAccountsInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountsInfoOK, error)

	GetBalance(params *GetBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBalanceOK, error)

	GetFunding(params *GetFundingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFundingOK, error)

	GetPositions(params *GetPositionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPositionsOK, error)

	GetPositionsWithFilter(params *GetPositionsWithFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPositionsWithFilterOK, error)

	GetSubAccounts(params *GetSubAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSubAccountsOK, error)

	Tradebusts(params *TradebustsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TradebustsOK, error)

	Transactions(params *TransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TransactionsOK, error)

	UpdateAccountMaxSlippage(params *UpdateAccountMaxSlippageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountMaxSlippageOK, error)

	UpdateAccountProfileReferralCode(params *UpdateAccountProfileReferralCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountProfileReferralCodeOK, error)

	UpdateAccountProfileUsername(params *UpdateAccountProfileUsernameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountProfileUsernameOK, error)

	UpsertAccountMargin(params *UpsertAccountMarginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertAccountMarginOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetFills lists fills

	This API returns a list of matched orders (i.e. fills) that have been sent to

chain for settlement.
*/
func (a *Client) GetFills(params *GetFillsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFillsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFillsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFills",
		Method:             "GET",
		PathPattern:        "/fills",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFillsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFillsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFills: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccount gets account information

Respond with requester's account information
*/
func (a *Client) GetAccount(params *GetAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-account",
		Method:             "GET",
		PathPattern:        "/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-account: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountMargin gets account margin configuration

Get margin configuration for an account and market
*/
func (a *Client) GetAccountMargin(params *GetAccountMarginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountMarginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountMarginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-account-margin",
		Method:             "GET",
		PathPattern:        "/account/margin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountMarginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountMarginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-account-margin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountProfile gets account profile information

Respond with requester's account information
*/
func (a *Client) GetAccountProfile(params *GetAccountProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-account-profile",
		Method:             "GET",
		PathPattern:        "/account/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-account-profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountsInfo returns account info of current account
*/
func (a *Client) GetAccountsInfo(params *GetAccountsInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountsInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-accounts-info",
		Method:             "GET",
		PathPattern:        "/account/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-accounts-info: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBalance lists balances

Respond with requester's own balances
*/
func (a *Client) GetBalance(params *GetBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-balance",
		Method:             "GET",
		PathPattern:        "/balance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBalanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-balance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetFunding fundings payments history

	List funding payments made by/to the requester's account

*Funding payments are periodic payments between traders to make the perpetual
futures contract price is close to the index price.*
*/
func (a *Client) GetFunding(params *GetFundingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFundingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFundingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-funding",
		Method:             "GET",
		PathPattern:        "/funding/payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFundingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFundingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-funding: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPositions lists open positions

Get all positions owned by current user
*/
func (a *Client) GetPositions(params *GetPositionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPositionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPositionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-positions",
		Method:             "GET",
		PathPattern:        "/positions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPositionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPositionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-positions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPositionsWithFilter lists closed positions history

Get all positions owned by current user
*/
func (a *Client) GetPositionsWithFilter(params *GetPositionsWithFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPositionsWithFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPositionsWithFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-positions-with-filter",
		Method:             "GET",
		PathPattern:        "/positions-history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPositionsWithFilterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPositionsWithFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-positions-with-filter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSubAccounts lists sub accounts of current account

Respond with requester's list of sub-accounts
*/
func (a *Client) GetSubAccounts(params *GetSubAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSubAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-sub-accounts",
		Method:             "GET",
		PathPattern:        "/account/subaccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSubAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSubAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-sub-accounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Tradebusts lists tradebusts

Retrieves a list of tradebusts associated to the requester's account
*/
func (a *Client) Tradebusts(params *TradebustsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TradebustsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradebustsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tradebusts",
		Method:             "GET",
		PathPattern:        "/tradebusts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradebustsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TradebustsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tradebusts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Transactions lists transactions

Retrieves a list of transactions initiated by the user
*/
func (a *Client) Transactions(params *TransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "transactions",
		Method:             "GET",
		PathPattern:        "/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAccountMaxSlippage updates account max slippage

Respond with requester's account max_slippage
*/
func (a *Client) UpdateAccountMaxSlippage(params *UpdateAccountMaxSlippageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountMaxSlippageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountMaxSlippageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-account-max-slippage",
		Method:             "POST",
		PathPattern:        "/account/profile/max_slippage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountMaxSlippageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAccountMaxSlippageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-account-max-slippage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAccountProfileReferralCode updates account profile fields

Respond with requester's account information
*/
func (a *Client) UpdateAccountProfileReferralCode(params *UpdateAccountProfileReferralCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountProfileReferralCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountProfileReferralCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-account-profile-referral-code",
		Method:             "POST",
		PathPattern:        "/account/profile/referral_code",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountProfileReferralCodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAccountProfileReferralCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-account-profile-referral-code: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAccountProfileUsername updates account profile username fields

Respond with requester's account information
*/
func (a *Client) UpdateAccountProfileUsername(params *UpdateAccountProfileUsernameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccountProfileUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountProfileUsernameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-account-profile-username",
		Method:             "POST",
		PathPattern:        "/account/profile/username",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountProfileUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAccountProfileUsernameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-account-profile-username: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpsertAccountMargin sets margin configuration

Set margin configuration for an account and market
*/
func (a *Client) UpsertAccountMargin(params *UpsertAccountMarginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertAccountMarginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertAccountMarginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upsert-account-margin",
		Method:             "POST",
		PathPattern:        "/account/margin/{market}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertAccountMarginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpsertAccountMarginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upsert-account-margin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
