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

// GetAssetReader is a Reader for the GetAsset structure.
type GetAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAssetOK creates a GetAssetOK with default headers values
func NewGetAssetOK() *GetAssetOK {
	return &GetAssetOK{}
}

/*
GetAssetOK describes a response with status code 200, with default header values.

OK
*/
type GetAssetOK struct {
	Payload []*models.AssetData
}

// IsSuccess returns true when this get asset o k response has a 2xx status code
func (o *GetAssetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get asset o k response has a 3xx status code
func (o *GetAssetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset o k response has a 4xx status code
func (o *GetAssetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get asset o k response has a 5xx status code
func (o *GetAssetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset o k response a status code equal to that given
func (o *GetAssetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get asset o k response
func (o *GetAssetOK) Code() int {
	return 200
}

func (o *GetAssetOK) Error() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetOK  %+v", 200, o.Payload)
}

func (o *GetAssetOK) String() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetOK  %+v", 200, o.Payload)
}

func (o *GetAssetOK) GetPayload() []*models.AssetData {
	return o.Payload
}

func (o *GetAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetBadRequest creates a GetAssetBadRequest with default headers values
func NewGetAssetBadRequest() *GetAssetBadRequest {
	return &GetAssetBadRequest{}
}

/*
GetAssetBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetAssetBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get asset bad request response has a 2xx status code
func (o *GetAssetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get asset bad request response has a 3xx status code
func (o *GetAssetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset bad request response has a 4xx status code
func (o *GetAssetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get asset bad request response has a 5xx status code
func (o *GetAssetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset bad request response a status code equal to that given
func (o *GetAssetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get asset bad request response
func (o *GetAssetBadRequest) Code() int {
	return 400
}

func (o *GetAssetBadRequest) Error() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetBadRequest  %+v", 400, o.Payload)
}

func (o *GetAssetBadRequest) String() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetBadRequest  %+v", 400, o.Payload)
}

func (o *GetAssetBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetAssetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetNotFound creates a GetAssetNotFound with default headers values
func NewGetAssetNotFound() *GetAssetNotFound {
	return &GetAssetNotFound{}
}

/*
GetAssetNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetAssetNotFound struct {
}

// IsSuccess returns true when this get asset not found response has a 2xx status code
func (o *GetAssetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get asset not found response has a 3xx status code
func (o *GetAssetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset not found response has a 4xx status code
func (o *GetAssetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get asset not found response has a 5xx status code
func (o *GetAssetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset not found response a status code equal to that given
func (o *GetAssetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get asset not found response
func (o *GetAssetNotFound) Code() int {
	return 404
}

func (o *GetAssetNotFound) Error() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetNotFound ", 404)
}

func (o *GetAssetNotFound) String() string {
	return fmt.Sprintf("[GET /assets/{path}/{name}][%d] getAssetNotFound ", 404)
}

func (o *GetAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
