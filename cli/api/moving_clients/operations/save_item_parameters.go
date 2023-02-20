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

	"github.com/alydnhrealgang/moving/cli/api/models"
)

// NewSaveItemParams creates a new SaveItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveItemParams() *SaveItemParams {
	return &SaveItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveItemParamsWithTimeout creates a new SaveItemParams object
// with the ability to set a timeout on a request.
func NewSaveItemParamsWithTimeout(timeout time.Duration) *SaveItemParams {
	return &SaveItemParams{
		timeout: timeout,
	}
}

// NewSaveItemParamsWithContext creates a new SaveItemParams object
// with the ability to set a context for a request.
func NewSaveItemParamsWithContext(ctx context.Context) *SaveItemParams {
	return &SaveItemParams{
		Context: ctx,
	}
}

// NewSaveItemParamsWithHTTPClient creates a new SaveItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveItemParamsWithHTTPClient(client *http.Client) *SaveItemParams {
	return &SaveItemParams{
		HTTPClient: client,
	}
}

/*
SaveItemParams contains all the parameters to send to the API endpoint

	for the save item operation.

	Typically these are written to a http.Request.
*/
type SaveItemParams struct {

	// Body.
	Body *models.ItemData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveItemParams) WithDefaults() *SaveItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save item params
func (o *SaveItemParams) WithTimeout(timeout time.Duration) *SaveItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save item params
func (o *SaveItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save item params
func (o *SaveItemParams) WithContext(ctx context.Context) *SaveItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save item params
func (o *SaveItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save item params
func (o *SaveItemParams) WithHTTPClient(client *http.Client) *SaveItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save item params
func (o *SaveItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the save item params
func (o *SaveItemParams) WithBody(body *models.ItemData) *SaveItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the save item params
func (o *SaveItemParams) SetBody(body *models.ItemData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SaveItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
