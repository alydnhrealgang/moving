// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SaveItemHandlerFunc turns a function with the right signature into a save item handler
type SaveItemHandlerFunc func(SaveItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveItemHandlerFunc) Handle(params SaveItemParams) middleware.Responder {
	return fn(params)
}

// SaveItemHandler interface for that can handle valid save item params
type SaveItemHandler interface {
	Handle(SaveItemParams) middleware.Responder
}

// NewSaveItem creates a new http.Handler for the save item operation
func NewSaveItem(ctx *middleware.Context, handler SaveItemHandler) *SaveItem {
	return &SaveItem{Context: ctx, Handler: handler}
}

/*
	SaveItem swagger:route POST /items saveItem

Save a item
*/
type SaveItem struct {
	Context *middleware.Context
	Handler SaveItemHandler
}

func (o *SaveItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSaveItemParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
