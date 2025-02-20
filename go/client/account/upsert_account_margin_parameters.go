// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/tradeparadex/api/models"
)

// NewUpsertAccountMarginParams creates a new UpsertAccountMarginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpsertAccountMarginParams() *UpsertAccountMarginParams {
	return &UpsertAccountMarginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertAccountMarginParamsWithTimeout creates a new UpsertAccountMarginParams object
// with the ability to set a timeout on a request.
func NewUpsertAccountMarginParamsWithTimeout(timeout time.Duration) *UpsertAccountMarginParams {
	return &UpsertAccountMarginParams{
		timeout: timeout,
	}
}

// NewUpsertAccountMarginParamsWithContext creates a new UpsertAccountMarginParams object
// with the ability to set a context for a request.
func NewUpsertAccountMarginParamsWithContext(ctx context.Context) *UpsertAccountMarginParams {
	return &UpsertAccountMarginParams{
		Context: ctx,
	}
}

// NewUpsertAccountMarginParamsWithHTTPClient creates a new UpsertAccountMarginParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpsertAccountMarginParamsWithHTTPClient(client *http.Client) *UpsertAccountMarginParams {
	return &UpsertAccountMarginParams{
		HTTPClient: client,
	}
}

/*
UpsertAccountMarginParams contains all the parameters to send to the API endpoint

	for the upsert account margin operation.

	Typically these are written to a http.Request.
*/
type UpsertAccountMarginParams struct {

	/* Config.

	   Margin Configuration
	*/
	Config *models.RequestsAccountMarginRequest

	/* Market.

	   Market Name - example: BTC-USD-PERP
	*/
	Market string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upsert account margin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertAccountMarginParams) WithDefaults() *UpsertAccountMarginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upsert account margin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertAccountMarginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upsert account margin params
func (o *UpsertAccountMarginParams) WithTimeout(timeout time.Duration) *UpsertAccountMarginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert account margin params
func (o *UpsertAccountMarginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert account margin params
func (o *UpsertAccountMarginParams) WithContext(ctx context.Context) *UpsertAccountMarginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert account margin params
func (o *UpsertAccountMarginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert account margin params
func (o *UpsertAccountMarginParams) WithHTTPClient(client *http.Client) *UpsertAccountMarginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert account margin params
func (o *UpsertAccountMarginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the upsert account margin params
func (o *UpsertAccountMarginParams) WithConfig(config *models.RequestsAccountMarginRequest) *UpsertAccountMarginParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the upsert account margin params
func (o *UpsertAccountMarginParams) SetConfig(config *models.RequestsAccountMarginRequest) {
	o.Config = config
}

// WithMarket adds the market to the upsert account margin params
func (o *UpsertAccountMarginParams) WithMarket(market string) *UpsertAccountMarginParams {
	o.SetMarket(market)
	return o
}

// SetMarket adds the market to the upsert account margin params
func (o *UpsertAccountMarginParams) SetMarket(market string) {
	o.Market = market
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertAccountMarginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param market
	if err := r.SetPathParam("market", o.Market); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
