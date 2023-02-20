// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteItemByCodeHandlerFunc turns a function with the right signature into a delete item by code handler
type DeleteItemByCodeHandlerFunc func(DeleteItemByCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteItemByCodeHandlerFunc) Handle(params DeleteItemByCodeParams) middleware.Responder {
	return fn(params)
}

// DeleteItemByCodeHandler interface for that can handle valid delete item by code params
type DeleteItemByCodeHandler interface {
	Handle(DeleteItemByCodeParams) middleware.Responder
}

// NewDeleteItemByCode creates a new http.Handler for the delete item by code operation
func NewDeleteItemByCode(ctx *middleware.Context, handler DeleteItemByCodeHandler) *DeleteItemByCode {
	return &DeleteItemByCode{Context: ctx, Handler: handler}
}

/*
	DeleteItemByCode swagger:route DELETE /item/{code} deleteItemByCode

Delete an item and its assets
*/
type DeleteItemByCode struct {
	Context *middleware.Context
	Handler DeleteItemByCodeHandler
}

func (o *DeleteItemByCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteItemByCodeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
