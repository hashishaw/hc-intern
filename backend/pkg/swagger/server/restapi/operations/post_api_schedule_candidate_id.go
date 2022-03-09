// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAPIScheduleCandidateIDHandlerFunc turns a function with the right signature into a post API schedule candidate ID handler
type PostAPIScheduleCandidateIDHandlerFunc func(PostAPIScheduleCandidateIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAPIScheduleCandidateIDHandlerFunc) Handle(params PostAPIScheduleCandidateIDParams) middleware.Responder {
	return fn(params)
}

// PostAPIScheduleCandidateIDHandler interface for that can handle valid post API schedule candidate ID params
type PostAPIScheduleCandidateIDHandler interface {
	Handle(PostAPIScheduleCandidateIDParams) middleware.Responder
}

// NewPostAPIScheduleCandidateID creates a new http.Handler for the post API schedule candidate ID operation
func NewPostAPIScheduleCandidateID(ctx *middleware.Context, handler PostAPIScheduleCandidateIDHandler) *PostAPIScheduleCandidateID {
	return &PostAPIScheduleCandidateID{Context: ctx, Handler: handler}
}

/*PostAPIScheduleCandidateID swagger:route POST /api/schedule/{candidateId} postApiScheduleCandidateId

Returns an interview id.

*/
type PostAPIScheduleCandidateID struct {
	Context *middleware.Context
	Handler PostAPIScheduleCandidateIDHandler
}

func (o *PostAPIScheduleCandidateID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAPIScheduleCandidateIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}