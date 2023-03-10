// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/alydnhrealgang/moving/cli/api/models"
)

// GetItemByCodeReader is a Reader for the GetItemByCode structure.
type GetItemByCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemByCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetItemByCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetItemByCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetItemByCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetItemByCodeOK creates a GetItemByCodeOK with default headers values
func NewGetItemByCodeOK() *GetItemByCodeOK {
	return &GetItemByCodeOK{}
}

/*
GetItemByCodeOK describes a response with status code 200, with default header values.

OK
*/
type GetItemByCodeOK struct {
	Payload []*models.ItemData
}

// IsSuccess returns true when this get item by code o k response has a 2xx status code
func (o *GetItemByCodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get item by code o k response has a 3xx status code
func (o *GetItemByCodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item by code o k response has a 4xx status code
func (o *GetItemByCodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item by code o k response has a 5xx status code
func (o *GetItemByCodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get item by code o k response a status code equal to that given
func (o *GetItemByCodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get item by code o k response
func (o *GetItemByCodeOK) Code() int {
	return 200
}

func (o *GetItemByCodeOK) Error() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeOK  %+v", 200, o.Payload)
}

func (o *GetItemByCodeOK) String() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeOK  %+v", 200, o.Payload)
}

func (o *GetItemByCodeOK) GetPayload() []*models.ItemData {
	return o.Payload
}

func (o *GetItemByCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemByCodeBadRequest creates a GetItemByCodeBadRequest with default headers values
func NewGetItemByCodeBadRequest() *GetItemByCodeBadRequest {
	return &GetItemByCodeBadRequest{}
}

/*
GetItemByCodeBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetItemByCodeBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get item by code bad request response has a 2xx status code
func (o *GetItemByCodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item by code bad request response has a 3xx status code
func (o *GetItemByCodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item by code bad request response has a 4xx status code
func (o *GetItemByCodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item by code bad request response has a 5xx status code
func (o *GetItemByCodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get item by code bad request response a status code equal to that given
func (o *GetItemByCodeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get item by code bad request response
func (o *GetItemByCodeBadRequest) Code() int {
	return 400
}

func (o *GetItemByCodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeBadRequest  %+v", 400, o.Payload)
}

func (o *GetItemByCodeBadRequest) String() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeBadRequest  %+v", 400, o.Payload)
}

func (o *GetItemByCodeBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetItemByCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemByCodeNotFound creates a GetItemByCodeNotFound with default headers values
func NewGetItemByCodeNotFound() *GetItemByCodeNotFound {
	return &GetItemByCodeNotFound{}
}

/*
GetItemByCodeNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetItemByCodeNotFound struct {
	Payload string
}

// IsSuccess returns true when this get item by code not found response has a 2xx status code
func (o *GetItemByCodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item by code not found response has a 3xx status code
func (o *GetItemByCodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item by code not found response has a 4xx status code
func (o *GetItemByCodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item by code not found response has a 5xx status code
func (o *GetItemByCodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get item by code not found response a status code equal to that given
func (o *GetItemByCodeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get item by code not found response
func (o *GetItemByCodeNotFound) Code() int {
	return 404
}

func (o *GetItemByCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeNotFound  %+v", 404, o.Payload)
}

func (o *GetItemByCodeNotFound) String() string {
	return fmt.Sprintf("[GET /item/{code}][%d] getItemByCodeNotFound  %+v", 404, o.Payload)
}

func (o *GetItemByCodeNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetItemByCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
