// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EnableHandlerFunc turns a function with the right signature into a enable handler
type EnableHandlerFunc func(EnableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EnableHandlerFunc) Handle(params EnableParams) middleware.Responder {
	return fn(params)
}

// EnableHandler interface for that can handle valid enable params
type EnableHandler interface {
	Handle(EnableParams) middleware.Responder
}

// NewEnable creates a new http.Handler for the enable operation
func NewEnable(ctx *middleware.Context, handler EnableHandler) *Enable {
	return &Enable{Context: ctx, Handler: handler}
}

/* Enable swagger:route POST /enable identity enable

Enable enable API

*/
type Enable struct {
	Context *middleware.Context
	Handler EnableHandler
}

func (o *Enable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEnableParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}