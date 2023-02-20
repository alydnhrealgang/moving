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

// NewUploadAssetParams creates a new UploadAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadAssetParams() *UploadAssetParams {
	return &UploadAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadAssetParamsWithTimeout creates a new UploadAssetParams object
// with the ability to set a timeout on a request.
func NewUploadAssetParamsWithTimeout(timeout time.Duration) *UploadAssetParams {
	return &UploadAssetParams{
		timeout: timeout,
	}
}

// NewUploadAssetParamsWithContext creates a new UploadAssetParams object
// with the ability to set a context for a request.
func NewUploadAssetParamsWithContext(ctx context.Context) *UploadAssetParams {
	return &UploadAssetParams{
		Context: ctx,
	}
}

// NewUploadAssetParamsWithHTTPClient creates a new UploadAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadAssetParamsWithHTTPClient(client *http.Client) *UploadAssetParams {
	return &UploadAssetParams{
		HTTPClient: client,
	}
}

/*
UploadAssetParams contains all the parameters to send to the API endpoint

	for the upload asset operation.

	Typically these are written to a http.Request.
*/
type UploadAssetParams struct {

	// ContentType.
	ContentType string

	// File.
	File runtime.NamedReadCloser

	// Name.
	Name string

	// Path.
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadAssetParams) WithDefaults() *UploadAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload asset params
func (o *UploadAssetParams) WithTimeout(timeout time.Duration) *UploadAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload asset params
func (o *UploadAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload asset params
func (o *UploadAssetParams) WithContext(ctx context.Context) *UploadAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload asset params
func (o *UploadAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload asset params
func (o *UploadAssetParams) WithHTTPClient(client *http.Client) *UploadAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload asset params
func (o *UploadAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the upload asset params
func (o *UploadAssetParams) WithContentType(contentType string) *UploadAssetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the upload asset params
func (o *UploadAssetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithFile adds the file to the upload asset params
func (o *UploadAssetParams) WithFile(file runtime.NamedReadCloser) *UploadAssetParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload asset params
func (o *UploadAssetParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithName adds the name to the upload asset params
func (o *UploadAssetParams) WithName(name string) *UploadAssetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upload asset params
func (o *UploadAssetParams) SetName(name string) {
	o.Name = name
}

// WithPath adds the path to the upload asset params
func (o *UploadAssetParams) WithPath(path string) *UploadAssetParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the upload asset params
func (o *UploadAssetParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *UploadAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param contentType
	frContentType := o.ContentType
	fContentType := frContentType
	if fContentType != "" {
		if err := r.SetFormParam("contentType", fContentType); err != nil {
			return err
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	// form param path
	frPath := o.Path
	fPath := frPath
	if fPath != "" {
		if err := r.SetFormParam("path", fPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
