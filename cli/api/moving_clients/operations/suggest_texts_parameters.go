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

// NewSuggestTextsParams creates a new SuggestTextsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuggestTextsParams() *SuggestTextsParams {
	return &SuggestTextsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuggestTextsParamsWithTimeout creates a new SuggestTextsParams object
// with the ability to set a timeout on a request.
func NewSuggestTextsParamsWithTimeout(timeout time.Duration) *SuggestTextsParams {
	return &SuggestTextsParams{
		timeout: timeout,
	}
}

// NewSuggestTextsParamsWithContext creates a new SuggestTextsParams object
// with the ability to set a context for a request.
func NewSuggestTextsParamsWithContext(ctx context.Context) *SuggestTextsParams {
	return &SuggestTextsParams{
		Context: ctx,
	}
}

// NewSuggestTextsParamsWithHTTPClient creates a new SuggestTextsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuggestTextsParamsWithHTTPClient(client *http.Client) *SuggestTextsParams {
	return &SuggestTextsParams{
		HTTPClient: client,
	}
}

/*
SuggestTextsParams contains all the parameters to send to the API endpoint

	for the suggest texts operation.

	Typically these are written to a http.Request.
*/
type SuggestTextsParams struct {

	// Name.
	Name string

	// Text.
	Text string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suggest texts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuggestTextsParams) WithDefaults() *SuggestTextsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suggest texts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuggestTextsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the suggest texts params
func (o *SuggestTextsParams) WithTimeout(timeout time.Duration) *SuggestTextsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suggest texts params
func (o *SuggestTextsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suggest texts params
func (o *SuggestTextsParams) WithContext(ctx context.Context) *SuggestTextsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suggest texts params
func (o *SuggestTextsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suggest texts params
func (o *SuggestTextsParams) WithHTTPClient(client *http.Client) *SuggestTextsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suggest texts params
func (o *SuggestTextsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the suggest texts params
func (o *SuggestTextsParams) WithName(name string) *SuggestTextsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the suggest texts params
func (o *SuggestTextsParams) SetName(name string) {
	o.Name = name
}

// WithText adds the text to the suggest texts params
func (o *SuggestTextsParams) WithText(text string) *SuggestTextsParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the suggest texts params
func (o *SuggestTextsParams) SetText(text string) {
	o.Text = text
}

// WriteToRequest writes these params to a swagger request
func (o *SuggestTextsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param text
	if err := r.SetPathParam("text", o.Text); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}