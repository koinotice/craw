// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// FindExecutionsHandlerFunc turns a function with the right signature into a find executions handler
type FindExecutionsHandlerFunc func(FindExecutionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindExecutionsHandlerFunc) Handle(params FindExecutionsParams) middleware.Responder {
	return fn(params)
}

// FindExecutionsHandler interface for that can handle valid find executions params
type FindExecutionsHandler interface {
	Handle(FindExecutionsParams) middleware.Responder
}

// NewFindExecutions creates a new http.Handler for the find executions operation
func NewFindExecutions(ctx *middleware.Context, handler FindExecutionsHandler) *FindExecutions {
	return &FindExecutions{Context: ctx, Handler: handler}
}

/*FindExecutions swagger:route GET /projects/{projectId}/exec findExecutions

List Execution

*/
type FindExecutions struct {
	Context *middleware.Context
	Handler FindExecutionsHandler
}

func (o *FindExecutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindExecutionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FindExecutionsOKBody find executions o k body
// swagger:model FindExecutionsOKBody
type FindExecutionsOKBody struct {
	models.SearchResult

	// data
	Data []*models.ExecutionOutput `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *FindExecutionsOKBody) UnmarshalJSON(raw []byte) error {
	// FindExecutionsOKBodyAO0
	var findExecutionsOKBodyAO0 models.SearchResult
	if err := swag.ReadJSON(raw, &findExecutionsOKBodyAO0); err != nil {
		return err
	}
	o.SearchResult = findExecutionsOKBodyAO0

	// FindExecutionsOKBodyAO1
	var dataFindExecutionsOKBodyAO1 struct {
		Data []*models.ExecutionOutput `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataFindExecutionsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataFindExecutionsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o FindExecutionsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	findExecutionsOKBodyAO0, err := swag.WriteJSON(o.SearchResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, findExecutionsOKBodyAO0)

	var dataFindExecutionsOKBodyAO1 struct {
		Data []*models.ExecutionOutput `json:"data"`
	}

	dataFindExecutionsOKBodyAO1.Data = o.Data

	jsonDataFindExecutionsOKBodyAO1, errFindExecutionsOKBodyAO1 := swag.WriteJSON(dataFindExecutionsOKBodyAO1)
	if errFindExecutionsOKBodyAO1 != nil {
		return nil, errFindExecutionsOKBodyAO1
	}
	_parts = append(_parts, jsonDataFindExecutionsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this find executions o k body
func (o *FindExecutionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SearchResult
	if err := o.SearchResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindExecutionsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findExecutionsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *FindExecutionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindExecutionsOKBody) UnmarshalBinary(b []byte) error {
	var res FindExecutionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
