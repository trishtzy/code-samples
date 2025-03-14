// Code generated by go-swagger; DO NOT EDIT.

package markets

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
)

// NewGetBboParams creates a new GetBboParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBboParams() *GetBboParams {
	return &GetBboParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBboParamsWithTimeout creates a new GetBboParams object
// with the ability to set a timeout on a request.
func NewGetBboParamsWithTimeout(timeout time.Duration) *GetBboParams {
	return &GetBboParams{
		timeout: timeout,
	}
}

// NewGetBboParamsWithContext creates a new GetBboParams object
// with the ability to set a context for a request.
func NewGetBboParamsWithContext(ctx context.Context) *GetBboParams {
	return &GetBboParams{
		Context: ctx,
	}
}

// NewGetBboParamsWithHTTPClient creates a new GetBboParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBboParamsWithHTTPClient(client *http.Client) *GetBboParams {
	return &GetBboParams{
		HTTPClient: client,
	}
}

/*
GetBboParams contains all the parameters to send to the API endpoint

	for the get bbo operation.

	Typically these are written to a http.Request.
*/
type GetBboParams struct {

	/* Market.

	   Market symbol - ex: BTC-USD-PERP
	*/
	Market string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bbo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBboParams) WithDefaults() *GetBboParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bbo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBboParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bbo params
func (o *GetBboParams) WithTimeout(timeout time.Duration) *GetBboParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bbo params
func (o *GetBboParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bbo params
func (o *GetBboParams) WithContext(ctx context.Context) *GetBboParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bbo params
func (o *GetBboParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bbo params
func (o *GetBboParams) WithHTTPClient(client *http.Client) *GetBboParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bbo params
func (o *GetBboParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMarket adds the market to the get bbo params
func (o *GetBboParams) WithMarket(market string) *GetBboParams {
	o.SetMarket(market)
	return o
}

// SetMarket adds the market to the get bbo params
func (o *GetBboParams) SetMarket(market string) {
	o.Market = market
}

// WriteToRequest writes these params to a swagger request
func (o *GetBboParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param market
	if err := r.SetPathParam("market", o.Market); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
