// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// QueryItemsURL generates an URL for the query items operation
type QueryItemsURL struct {
	FetchSize  int64
	Keyword    string
	StartIndex int64
	TagName    *string
	Type       string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryItemsURL) WithBasePath(bp string) *QueryItemsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryItemsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *QueryItemsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/items/query"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	fetchSizeQ := swag.FormatInt64(o.FetchSize)
	if fetchSizeQ != "" {
		qs.Set("fetchSize", fetchSizeQ)
	}

	keywordQ := o.Keyword
	if keywordQ != "" {
		qs.Set("keyword", keywordQ)
	}

	startIndexQ := swag.FormatInt64(o.StartIndex)
	if startIndexQ != "" {
		qs.Set("startIndex", startIndexQ)
	}

	var tagNameQ string
	if o.TagName != nil {
		tagNameQ = *o.TagName
	}
	if tagNameQ != "" {
		qs.Set("tagName", tagNameQ)
	}

	typeVarQ := o.Type
	if typeVarQ != "" {
		qs.Set("type", typeVarQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *QueryItemsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *QueryItemsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *QueryItemsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on QueryItemsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on QueryItemsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *QueryItemsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}