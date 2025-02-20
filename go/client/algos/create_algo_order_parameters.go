// Code generated by go-swagger; DO NOT EDIT.

package algos

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

// NewCreateAlgoOrderParams creates a new CreateAlgoOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAlgoOrderParams() *CreateAlgoOrderParams {
	return &CreateAlgoOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAlgoOrderParamsWithTimeout creates a new CreateAlgoOrderParams object
// with the ability to set a timeout on a request.
func NewCreateAlgoOrderParamsWithTimeout(timeout time.Duration) *CreateAlgoOrderParams {
	return &CreateAlgoOrderParams{
		timeout: timeout,
	}
}

// NewCreateAlgoOrderParamsWithContext creates a new CreateAlgoOrderParams object
// with the ability to set a context for a request.
func NewCreateAlgoOrderParamsWithContext(ctx context.Context) *CreateAlgoOrderParams {
	return &CreateAlgoOrderParams{
		Context: ctx,
	}
}

// NewCreateAlgoOrderParamsWithHTTPClient creates a new CreateAlgoOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAlgoOrderParamsWithHTTPClient(client *http.Client) *CreateAlgoOrderParams {
	return &CreateAlgoOrderParams{
		HTTPClient: client,
	}
}

/*
CreateAlgoOrderParams contains all the parameters to send to the API endpoint

	for the create algo order operation.

	Typically these are written to a http.Request.
*/
type CreateAlgoOrderParams struct {

	/* Params.

	   Algo order content
	*/
	Params *models.RequestsAlgoOrderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create algo order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAlgoOrderParams) WithDefaults() *CreateAlgoOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create algo order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAlgoOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create algo order params
func (o *CreateAlgoOrderParams) WithTimeout(timeout time.Duration) *CreateAlgoOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create algo order params
func (o *CreateAlgoOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create algo order params
func (o *CreateAlgoOrderParams) WithContext(ctx context.Context) *CreateAlgoOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create algo order params
func (o *CreateAlgoOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create algo order params
func (o *CreateAlgoOrderParams) WithHTTPClient(client *http.Client) *CreateAlgoOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create algo order params
func (o *CreateAlgoOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the create algo order params
func (o *CreateAlgoOrderParams) WithParams(params *models.RequestsAlgoOrderRequest) *CreateAlgoOrderParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the create algo order params
func (o *CreateAlgoOrderParams) SetParams(params *models.RequestsAlgoOrderRequest) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAlgoOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
