// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// NewUpdateScriptParams creates a new UpdateScriptParams object
// no default values defined in spec.
func NewUpdateScriptParams() UpdateScriptParams {

	return UpdateScriptParams{}
}

// UpdateScriptParams contains all the bound params for the update script operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateScript
type UpdateScriptParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body models.ScriptCreate
	/*
	  Required: true
	  Pattern: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	  In: path
	*/
	ProjectID string
	/*
	  Required: true
	  Pattern: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	  In: path
	*/
	ScriptID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateScriptParams() beforehand.
func (o *UpdateScriptParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ScriptCreate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}
	rProjectID, rhkProjectID, _ := route.Params.GetOK("projectId")
	if err := o.bindProjectID(rProjectID, rhkProjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	rScriptID, rhkScriptID, _ := route.Params.GetOK("scriptId")
	if err := o.bindScriptID(rScriptID, rhkScriptID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProjectID binds and validates parameter ProjectID from path.
func (o *UpdateScriptParams) bindProjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectID = raw

	if err := o.validateProjectID(formats); err != nil {
		return err
	}

	return nil
}

// validateProjectID carries on validations for parameter ProjectID
func (o *UpdateScriptParams) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Pattern("projectId", "path", o.ProjectID, `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`); err != nil {
		return err
	}

	return nil
}

// bindScriptID binds and validates parameter ScriptID from path.
func (o *UpdateScriptParams) bindScriptID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ScriptID = raw

	if err := o.validateScriptID(formats); err != nil {
		return err
	}

	return nil
}

// validateScriptID carries on validations for parameter ScriptID
func (o *UpdateScriptParams) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Pattern("scriptId", "path", o.ScriptID, `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`); err != nil {
		return err
	}

	return nil
}
