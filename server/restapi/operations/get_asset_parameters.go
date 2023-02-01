// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetAssetParams creates a new GetAssetParams object
//
// There are no default values defined in the spec.
func NewGetAssetParams() GetAssetParams {

	return GetAssetParams{}
}

// GetAssetParams contains all the bound params for the get asset operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAsset
type GetAssetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*如果为支持通配*，则获取所有path下的内容
	  Required: true
	  In: path
	*/
	Name string
	/*
	  Required: true
	  In: path
	*/
	Path string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAssetParams() beforehand.
func (o *GetAssetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *GetAssetParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	return nil
}

// bindPath binds and validates parameter Path from path.
func (o *GetAssetParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Path = raw

	return nil
}