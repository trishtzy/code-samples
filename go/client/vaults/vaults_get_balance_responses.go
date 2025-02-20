// Code generated by go-swagger; DO NOT EDIT.

package vaults

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

// VaultsGetBalanceReader is a Reader for the VaultsGetBalance structure.
type VaultsGetBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VaultsGetBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVaultsGetBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVaultsGetBalanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /vaults/balance] vaults-get-balance", response, response.Code())
	}
}

// NewVaultsGetBalanceOK creates a VaultsGetBalanceOK with default headers values
func NewVaultsGetBalanceOK() *VaultsGetBalanceOK {
	return &VaultsGetBalanceOK{}
}

/*
VaultsGetBalanceOK describes a response with status code 200, with default header values.

OK
*/
type VaultsGetBalanceOK struct {
	Payload *models.ResponsesGetBalancesResp
}

// IsSuccess returns true when this vaults get balance o k response has a 2xx status code
func (o *VaultsGetBalanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vaults get balance o k response has a 3xx status code
func (o *VaultsGetBalanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vaults get balance o k response has a 4xx status code
func (o *VaultsGetBalanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vaults get balance o k response has a 5xx status code
func (o *VaultsGetBalanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vaults get balance o k response a status code equal to that given
func (o *VaultsGetBalanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vaults get balance o k response
func (o *VaultsGetBalanceOK) Code() int {
	return 200
}

func (o *VaultsGetBalanceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vaults/balance][%d] vaultsGetBalanceOK %s", 200, payload)
}

func (o *VaultsGetBalanceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vaults/balance][%d] vaultsGetBalanceOK %s", 200, payload)
}

func (o *VaultsGetBalanceOK) GetPayload() *models.ResponsesGetBalancesResp {
	return o.Payload
}

func (o *VaultsGetBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesGetBalancesResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultsGetBalanceBadRequest creates a VaultsGetBalanceBadRequest with default headers values
func NewVaultsGetBalanceBadRequest() *VaultsGetBalanceBadRequest {
	return &VaultsGetBalanceBadRequest{}
}

/*
VaultsGetBalanceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VaultsGetBalanceBadRequest struct {
	Payload *models.ResponsesAPIError
}

// IsSuccess returns true when this vaults get balance bad request response has a 2xx status code
func (o *VaultsGetBalanceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this vaults get balance bad request response has a 3xx status code
func (o *VaultsGetBalanceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vaults get balance bad request response has a 4xx status code
func (o *VaultsGetBalanceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this vaults get balance bad request response has a 5xx status code
func (o *VaultsGetBalanceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this vaults get balance bad request response a status code equal to that given
func (o *VaultsGetBalanceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the vaults get balance bad request response
func (o *VaultsGetBalanceBadRequest) Code() int {
	return 400
}

func (o *VaultsGetBalanceBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vaults/balance][%d] vaultsGetBalanceBadRequest %s", 400, payload)
}

func (o *VaultsGetBalanceBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vaults/balance][%d] vaultsGetBalanceBadRequest %s", 400, payload)
}

func (o *VaultsGetBalanceBadRequest) GetPayload() *models.ResponsesAPIError {
	return o.Payload
}

func (o *VaultsGetBalanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
