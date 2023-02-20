// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewMoveItemsParams creates a new MoveItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMoveItemsParams() *MoveItemsParams {
	return &MoveItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMoveItemsParamsWithTimeout creates a new MoveItemsParams object
// with the ability to set a timeout on a request.
func NewMoveItemsParamsWithTimeout(timeout time.Duration) *MoveItemsParams {
	return &MoveItemsParams{
		timeout: timeout,
	}
}

// NewMoveItemsParamsWithContext creates a new MoveItemsParams object
// with the ability to set a context for a request.
func NewMoveItemsParamsWithContext(ctx context.Context) *MoveItemsParams {
	return &MoveItemsParams{
		Context: ctx,
	}
}

// NewMoveItemsParamsWithHTTPClient creates a new MoveItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewMoveItemsParamsWithHTTPClient(client *http.Client) *MoveItemsParams {
	return &MoveItemsParams{
		HTTPClient: client,
	}
}

/*
MoveItemsParams contains all the parameters to send to the API endpoint

	for the move items operation.

	Typically these are written to a http.Request.
*/
type MoveItemsParams struct {

	// Body.
	Body MoveItemsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the move items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveItemsParams) WithDefaults() *MoveItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the move items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the move items params
func (o *MoveItemsParams) WithTimeout(timeout time.Duration) *MoveItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move items params
func (o *MoveItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move items params
func (o *MoveItemsParams) WithContext(ctx context.Context) *MoveItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move items params
func (o *MoveItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move items params
func (o *MoveItemsParams) WithHTTPClient(client *http.Client) *MoveItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move items params
func (o *MoveItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move items params
func (o *MoveItemsParams) WithBody(body MoveItemsBody) *MoveItemsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move items params
func (o *MoveItemsParams) SetBody(body MoveItemsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MoveItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}