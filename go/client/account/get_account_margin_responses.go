// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tradeparadex/api/models"
)

// GetAccountMarginReader is a Reader for the GetAccountMargin structure.
type GetAccountMarginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountMarginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountMarginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountMarginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountMarginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /account/margin] get-account-margin", response, response.Code())
	}
}

// NewGetAccountMarginOK creates a GetAccountMarginOK with default headers values
func NewGetAccountMarginOK() *GetAccountMarginOK {
	return &GetAccountMarginOK{}
}

/*
GetAccountMarginOK describes a response with status code 200, with default header values.

OK
*/
type GetAccountMarginOK struct {
	Payload *models.MessageAccountMarginConfig
}

// IsSuccess returns true when this get account margin o k response has a 2xx status code
func (o *GetAccountMarginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get account margin o k response has a 3xx status code
func (o *GetAccountMarginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account margin o k response has a 4xx status code
func (o *GetAccountMarginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account margin o k response has a 5xx status code
func (o *GetAccountMarginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get account margin o k response a status code equal to that given
func (o *GetAccountMarginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get account margin o k response
func (o *GetAccountMarginOK) Code() int {
	return 200
}

func (o *GetAccountMarginOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginOK %s", 200, payload)
}

func (o *GetAccountMarginOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginOK %s", 200, payload)
}

func (o *GetAccountMarginOK) GetPayload() *models.MessageAccountMarginConfig {
	return o.Payload
}

func (o *GetAccountMarginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageAccountMarginConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountMarginBadRequest creates a GetAccountMarginBadRequest with default headers values
func NewGetAccountMarginBadRequest() *GetAccountMarginBadRequest {
	return &GetAccountMarginBadRequest{}
}

/*
GetAccountMarginBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAccountMarginBadRequest struct {
	Payload *models.ResponsesAPIError
}

// IsSuccess returns true when this get account margin bad request response has a 2xx status code
func (o *GetAccountMarginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account margin bad request response has a 3xx status code
func (o *GetAccountMarginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account margin bad request response has a 4xx status code
func (o *GetAccountMarginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account margin bad request response has a 5xx status code
func (o *GetAccountMarginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get account margin bad request response a status code equal to that given
func (o *GetAccountMarginBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get account margin bad request response
func (o *GetAccountMarginBadRequest) Code() int {
	return 400
}

func (o *GetAccountMarginBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginBadRequest %s", 400, payload)
}

func (o *GetAccountMarginBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginBadRequest %s", 400, payload)
}

func (o *GetAccountMarginBadRequest) GetPayload() *models.ResponsesAPIError {
	return o.Payload
}

func (o *GetAccountMarginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountMarginNotFound creates a GetAccountMarginNotFound with default headers values
func NewGetAccountMarginNotFound() *GetAccountMarginNotFound {
	return &GetAccountMarginNotFound{}
}

/*
GetAccountMarginNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAccountMarginNotFound struct {
	Payload *models.ResponsesAPIError
}

// IsSuccess returns true when this get account margin not found response has a 2xx status code
func (o *GetAccountMarginNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account margin not found response has a 3xx status code
func (o *GetAccountMarginNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account margin not found response has a 4xx status code
func (o *GetAccountMarginNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account margin not found response has a 5xx status code
func (o *GetAccountMarginNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get account margin not found response a status code equal to that given
func (o *GetAccountMarginNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get account margin not found response
func (o *GetAccountMarginNotFound) Code() int {
	return 404
}

func (o *GetAccountMarginNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginNotFound %s", 404, payload)
}

func (o *GetAccountMarginNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account/margin][%d] getAccountMarginNotFound %s", 404, payload)
}

func (o *GetAccountMarginNotFound) GetPayload() *models.ResponsesAPIError {
	return o.Payload
}

func (o *GetAccountMarginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
