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
	"github.com/go-openapi/swag"
)

// NewGetAlgoOrdersHistoryParams creates a new GetAlgoOrdersHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlgoOrdersHistoryParams() *GetAlgoOrdersHistoryParams {
	return &GetAlgoOrdersHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlgoOrdersHistoryParamsWithTimeout creates a new GetAlgoOrdersHistoryParams object
// with the ability to set a timeout on a request.
func NewGetAlgoOrdersHistoryParamsWithTimeout(timeout time.Duration) *GetAlgoOrdersHistoryParams {
	return &GetAlgoOrdersHistoryParams{
		timeout: timeout,
	}
}

// NewGetAlgoOrdersHistoryParamsWithContext creates a new GetAlgoOrdersHistoryParams object
// with the ability to set a context for a request.
func NewGetAlgoOrdersHistoryParamsWithContext(ctx context.Context) *GetAlgoOrdersHistoryParams {
	return &GetAlgoOrdersHistoryParams{
		Context: ctx,
	}
}

// NewGetAlgoOrdersHistoryParamsWithHTTPClient creates a new GetAlgoOrdersHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlgoOrdersHistoryParamsWithHTTPClient(client *http.Client) *GetAlgoOrdersHistoryParams {
	return &GetAlgoOrdersHistoryParams{
		HTTPClient: client,
	}
}

/*
GetAlgoOrdersHistoryParams contains all the parameters to send to the API endpoint

	for the get algo orders history operation.

	Typically these are written to a http.Request.
*/
type GetAlgoOrdersHistoryParams struct {

	/* ClientID.

	   Unique ID of client generating the order
	*/
	ClientID *string

	/* Cursor.

	   Returns the ‘next’ paginated page.
	*/
	Cursor *string

	/* EndAt.

	   End Time (unix time millisecond)
	*/
	EndAt *int64

	/* Market.

	   Market for the order
	*/
	Market *string

	/* PageSize.

	   Limit the number of responses in the page

	   Default: 100
	*/
	PageSize *int64

	/* Side.

	   Order side
	*/
	Side *string

	/* StartAt.

	   Start Time (unix time millisecond)
	*/
	StartAt *int64

	/* Status.

	   Order status
	*/
	Status *string

	/* Type.

	   Order type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get algo orders history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlgoOrdersHistoryParams) WithDefaults() *GetAlgoOrdersHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get algo orders history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlgoOrdersHistoryParams) SetDefaults() {
	var (
		pageSizeDefault = int64(100)
	)

	val := GetAlgoOrdersHistoryParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithTimeout(timeout time.Duration) *GetAlgoOrdersHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithContext(ctx context.Context) *GetAlgoOrdersHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithHTTPClient(client *http.Client) *GetAlgoOrdersHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithClientID(clientID *string) *GetAlgoOrdersHistoryParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithCursor adds the cursor to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithCursor(cursor *string) *GetAlgoOrdersHistoryParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithEndAt adds the endAt to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithEndAt(endAt *int64) *GetAlgoOrdersHistoryParams {
	o.SetEndAt(endAt)
	return o
}

// SetEndAt adds the endAt to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetEndAt(endAt *int64) {
	o.EndAt = endAt
}

// WithMarket adds the market to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithMarket(market *string) *GetAlgoOrdersHistoryParams {
	o.SetMarket(market)
	return o
}

// SetMarket adds the market to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetMarket(market *string) {
	o.Market = market
}

// WithPageSize adds the pageSize to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithPageSize(pageSize *int64) *GetAlgoOrdersHistoryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSide adds the side to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithSide(side *string) *GetAlgoOrdersHistoryParams {
	o.SetSide(side)
	return o
}

// SetSide adds the side to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetSide(side *string) {
	o.Side = side
}

// WithStartAt adds the startAt to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithStartAt(startAt *int64) *GetAlgoOrdersHistoryParams {
	o.SetStartAt(startAt)
	return o
}

// SetStartAt adds the startAt to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetStartAt(startAt *int64) {
	o.StartAt = startAt
}

// WithStatus adds the status to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithStatus(status *string) *GetAlgoOrdersHistoryParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) WithType(typeVar *string) *GetAlgoOrdersHistoryParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get algo orders history params
func (o *GetAlgoOrdersHistoryParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlgoOrdersHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param client_id
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("client_id", qClientID); err != nil {
				return err
			}
		}
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

	if o.EndAt != nil {

		// query param end_at
		var qrEndAt int64

		if o.EndAt != nil {
			qrEndAt = *o.EndAt
		}
		qEndAt := swag.FormatInt64(qrEndAt)
		if qEndAt != "" {

			if err := r.SetQueryParam("end_at", qEndAt); err != nil {
				return err
			}
		}
	}

	if o.Market != nil {

		// query param market
		var qrMarket string

		if o.Market != nil {
			qrMarket = *o.Market
		}
		qMarket := qrMarket
		if qMarket != "" {

			if err := r.SetQueryParam("market", qMarket); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Side != nil {

		// query param side
		var qrSide string

		if o.Side != nil {
			qrSide = *o.Side
		}
		qSide := qrSide
		if qSide != "" {

			if err := r.SetQueryParam("side", qSide); err != nil {
				return err
			}
		}
	}

	if o.StartAt != nil {

		// query param start_at
		var qrStartAt int64

		if o.StartAt != nil {
			qrStartAt = *o.StartAt
		}
		qStartAt := swag.FormatInt64(qrStartAt)
		if qStartAt != "" {

			if err := r.SetQueryParam("start_at", qStartAt); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
