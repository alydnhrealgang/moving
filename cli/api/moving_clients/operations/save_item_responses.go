// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SaveItemReader is a Reader for the SaveItem structure.
type SaveItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSaveItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSaveItemOK creates a SaveItemOK with default headers values
func NewSaveItemOK() *SaveItemOK {
	return &SaveItemOK{}
}

/*
SaveItemOK describes a response with status code 200, with default header values.

OK
*/
type SaveItemOK struct {
	Payload string
}

// IsSuccess returns true when this save item o k response has a 2xx status code
func (o *SaveItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save item o k response has a 3xx status code
func (o *SaveItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save item o k response has a 4xx status code
func (o *SaveItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this save item o k response has a 5xx status code
func (o *SaveItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this save item o k response a status code equal to that given
func (o *SaveItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the save item o k response
func (o *SaveItemOK) Code() int {
	return 200
}

func (o *SaveItemOK) Error() string {
	return fmt.Sprintf("[POST /items][%d] saveItemOK  %+v", 200, o.Payload)
}

func (o *SaveItemOK) String() string {
	return fmt.Sprintf("[POST /items][%d] saveItemOK  %+v", 200, o.Payload)
}

func (o *SaveItemOK) GetPayload() string {
	return o.Payload
}

func (o *SaveItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveItemBadRequest creates a SaveItemBadRequest with default headers values
func NewSaveItemBadRequest() *SaveItemBadRequest {
	return &SaveItemBadRequest{}
}

/*
SaveItemBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type SaveItemBadRequest struct {
	Payload string
}

// IsSuccess returns true when this save item bad request response has a 2xx status code
func (o *SaveItemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save item bad request response has a 3xx status code
func (o *SaveItemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save item bad request response has a 4xx status code
func (o *SaveItemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this save item bad request response has a 5xx status code
func (o *SaveItemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this save item bad request response a status code equal to that given
func (o *SaveItemBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the save item bad request response
func (o *SaveItemBadRequest) Code() int {
	return 400
}

func (o *SaveItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /items][%d] saveItemBadRequest  %+v", 400, o.Payload)
}

func (o *SaveItemBadRequest) String() string {
	return fmt.Sprintf("[POST /items][%d] saveItemBadRequest  %+v", 400, o.Payload)
}

func (o *SaveItemBadRequest) GetPayload() string {
	return o.Payload
}

func (o *SaveItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
