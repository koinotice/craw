// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// UpdateScriptDataOKCode is the HTTP code returned for type UpdateScriptDataOK
const UpdateScriptDataOKCode int = 200

/*UpdateScriptDataOK update script data o k

swagger:response updateScriptDataOK
*/
type UpdateScriptDataOK struct {

	/*
	  In: Body
	*/
	Payload *models.Entity `json:"body,omitempty"`
}

// NewUpdateScriptDataOK creates UpdateScriptDataOK with default headers values
func NewUpdateScriptDataOK() *UpdateScriptDataOK {

	return &UpdateScriptDataOK{}
}

// WithPayload adds the payload to the update script data o k response
func (o *UpdateScriptDataOK) WithPayload(payload *models.Entity) *UpdateScriptDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update script data o k response
func (o *UpdateScriptDataOK) SetPayload(payload *models.Entity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateScriptDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
